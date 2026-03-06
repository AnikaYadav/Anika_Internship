# Rook + Ceph on Minikube — Setup Cheat Sheet

A step-by-step guide to installing Rook/Ceph on a local Minikube cluster (Apple Silicon / arm64), provisioning storage, and working with snapshots.

---

## Prerequisites

- Minikube installed
- kubectl installed
- QEMU2 driver available (`brew install qemu` on Mac)

---

## Step 1: Create a Dedicated Minikube Profile

Create a fresh isolated Minikube cluster specifically for Rook/Ceph. Using a separate profile keeps this completely independent from other projects.

```bash
minikube start -p rook-cluster \
  --driver=qemu2 \
  --container-runtime=containerd \
  --memory=8192 \
  --cpus=4 \
  --extra-disks=1
```

| Flag | Why it's needed |
|---|---|
| `-p rook-cluster` | Creates a new isolated profile, won't affect other Minikube clusters |
| `--driver=qemu2` | Uses QEMU as the VM driver (required on Apple Silicon) |
| `--container-runtime=containerd` | Rook works best with containerd, avoids image pull issues |
| `--memory=8192` | Ceph is memory-hungry, needs at least 8GB |
| `--cpus=4` | Minimum CPUs for Rook operator + Ceph daemons |
| `--extra-disks=1` | Gives Ceph a raw unformatted disk to use as OSD storage |

Set this as the active profile:
```bash
minikube profile rook-cluster
```

Switch between profiles anytime:
```bash
minikube profile rook-cluster   # switch to rook
minikube profile minikube       # switch back to another project
```

---

## Step 2: Install Rook Operator

Apply the Rook CRDs, RBAC rules, and the operator itself. These files come from the official [Rook GitHub repo](https://github.com/rook/rook) at tag `v1.19.2`.

First, clone the repo so you can use local file paths for all commands:

```bash
git clone --branch v1.19.2 https://github.com/rook/rook.git
cd rook/deploy/examples
```

Then install the operator:

```bash
kubectl apply -f crds.yaml
kubectl apply -f common.yaml
kubectl apply -f operator.yaml
```

| File | What it does |
|---|---|
| `crds.yaml` | Registers Ceph-specific Kubernetes resource types (CephCluster, CephBlockPool, etc.) |
| `common.yaml` | Creates the `rook-ceph` namespace, service accounts, and RBAC rules |
| `operator.yaml` | Deploys the Rook operator pod which manages all Ceph resources |

Also apply the CSI operator config — **this is required and often missed**:
```bash
kubectl apply -f csi-operator.yaml
```

This installs the `CephConnection` and `OperatorConfig` CRDs that the Rook operator needs to configure CSI drivers.

Verify the operator is running:
```bash
kubectl get pods -n rook-ceph
# rook-ceph-operator-xxx   1/1   Running
```

---

## Step 3: Deploy the Ceph Cluster

Deploy a single-node Ceph cluster suitable for Minikube:

```bash
kubectl apply -f cluster-test.yaml
```

`cluster-test.yaml` is configured for single-node setups. The full `cluster.yaml` requires multiple nodes.

Watch the Ceph daemons come up (takes 3–5 minutes):
```bash
kubectl get pods -n rook-ceph -w
```

Expected pods:
| Pod | Role |
|---|---|
| `rook-ceph-mon-a` | Monitor — maintains cluster state and quorum |
| `rook-ceph-mgr-a` | Manager — handles metrics, dashboard, orchestration |
| `rook-ceph-osd-0` | Object Storage Daemon — stores the actual data |
| `rook-ceph-exporter` | Exports Ceph metrics |
| `ceph-csi-*` | CSI driver pods for block (RBD) and file (CephFS) storage |

Check overall cluster health:
```bash
kubectl get cephcluster -n rook-ceph
# PHASE: Ready   HEALTH: HEALTH_OK
```

---

## Step 4: Fix CephBlockPool for Single Node

By default the block pool requires 3 replicas (needs 3 OSDs). Since Minikube has only 1 OSD, edit the pool:

```bash
EDITOR=nano kubectl edit cephblockpool replicapool -n rook-ceph
```

Change the `replicated` section to:
```yaml
spec:
  replicated:
    size: 1
    requireSafeReplicaSize: false
```

Verify it becomes Ready:
```bash
kubectl get cephblockpool -n rook-ceph
# replicapool   Ready
```

---

## Step 5: Create a StorageClass

A StorageClass tells Kubernetes how to dynamically provision storage volumes using Ceph RBD (RADOS Block Device).

```bash
kubectl apply -f csi/rbd/storageclass-test.yaml
```

Verify:
```bash
kubectl get storageclass
# rook-ceph-block   rook-ceph.rbd.csi.ceph.com   Delete   Immediate
```

---

## Step 6: Provision a PVC

A PersistentVolumeClaim (PVC) requests storage from the StorageClass. Kubernetes + Rook automatically provision a volume in Ceph.

```bash
kubectl apply -f csi/rbd/pvc.yaml
```

Verify it is bound:
```bash
kubectl get pvc
# rbd-pvc   Bound   pvc-xxx   1Gi   RWO   rook-ceph-block
```

`Bound` means Ceph successfully allocated 1GB of block storage for this PVC.

---

## Step 7: Install VolumeSnapshot CRDs

Kubernetes doesn't ship with snapshot support by default. Install the CRDs and snapshot controller from the [external-snapshotter](https://github.com/kubernetes-csi/external-snapshotter) project:

```bash
# Clone the external-snapshotter repo (in a separate directory from rook)
git clone --branch v8.2.0 https://github.com/kubernetes-csi/external-snapshotter.git
cd external-snapshotter
```

```bash
kubectl apply -f client/config/crd/snapshot.storage.k8s.io_volumesnapshotclasses.yaml
kubectl apply -f client/config/crd/snapshot.storage.k8s.io_volumesnapshotcontents.yaml
kubectl apply -f client/config/crd/snapshot.storage.k8s.io_volumesnapshots.yaml

kubectl apply -f deploy/kubernetes/snapshot-controller/rbac-snapshot-controller.yaml
kubectl apply -f deploy/kubernetes/snapshot-controller/setup-snapshot-controller.yaml
```

Verify the snapshot controller is running:
```bash
kubectl get pods -n kube-system | grep snapshot
```

---

## Step 8: Create a VolumeSnapshotClass

A VolumeSnapshotClass tells Kubernetes which CSI driver to use when taking snapshots (similar to how StorageClass works for PVCs).

```bash
# Back in the rook/deploy/examples directory
cd ../rook/deploy/examples

kubectl apply -f csi/rbd/snapshotclass.yaml
kubectl get volumesnapshotclass
# csi-rbdplugin-snapclass   rook-ceph.rbd.csi.ceph.com   Delete
```

---

## Step 9: Take a Snapshot

Create a point-in-time snapshot of the PVC:

```bash
kubectl apply -f csi/rbd/snapshot.yaml
kubectl get volumesnapshot
# rbd-pvc-snapshot   true   rbd-pvc   1Gi   csi-rbdplugin-snapclass
```

`READYTOUSE: true` means the snapshot was successfully taken and stored in Ceph.

---

## Bonus: Restore from Snapshot

Create a new PVC from an existing snapshot:
```bash
kubectl apply -f csi/rbd/pvc-restore.yaml
```

Clone a PVC directly:
```bash
kubectl apply -f csi/rbd/pvc-clone.yaml
```

---

## Quick Reference

```bash
# Cluster health
kubectl get cephcluster -n rook-ceph

# All rook pods
kubectl get pods -n rook-ceph

# Storage
kubectl get storageclass
kubectl get pvc
kubectl get cephblockpool -n rook-ceph

# Snapshots
kubectl get volumesnapshotclass
kubectl get volumesnapshot
kubectl get volumesnapshotcontent

# Switch minikube profiles
minikube profile list
minikube profile rook-cluster
```

---

## How the Pieces Fit Together

```
StorageClass          → defines HOW storage is provisioned (which Ceph pool, which CSI driver)
    ↓
PVC                   → requests storage, Ceph allocates a block device
    ↓
VolumeSnapshotClass   → defines HOW snapshots are taken (which CSI driver)
    ↓
VolumeSnapshot        → point-in-time copy of the PVC data stored in Ceph
    ↓
VolumeSnapshotContent → auto-created, represents the actual snapshot data in Ceph
```

---

*All example files sourced from [github.com/rook/rook](https://github.com/rook/rook) at tag `v1.19.2` and [github.com/kubernetes-csi/external-snapshotter](https://github.com/kubernetes-csi/external-snapshotter) at tag `v8.2.0`.*