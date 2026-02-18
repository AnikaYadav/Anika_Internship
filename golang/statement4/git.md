# Git Cheatsheet 

## git config
**Description:** Set or view Git configuration like username, email.
**Examples:**
`git config --global user.name "Your Name"` # set global username
`git config --global user.email "your@email.com"` # set global email  
`git config --list` # view the config settings
**Notes:**
`--global` → applies to all repositories on your machine  
Without `--global` applies only to the current repository 


## git init
**Description:** Create a new Git repository  
**Variations:** `git init <folder-name>` creates a new folder and initializes a Git repository inside it.

## git clone
**Description:** Clone a remote repository given the URL
**Variations:** `--depth`, `--branch`  
**Example:**
 `git clone https://github.com/user/repo.git`
`git clone --depth 5 https://github.com/user/repo.git`  # last 5 commits
`git clone --branch <branch-name> <repo-url>` #clone repo and checkout a specific branch instead of main


## git status
**Description:** Checks the status of the repository. Lists all added, changed and newly created files.
**Variations:** `-s` for short output  
**Example:** 
`git status` 

## git add
**Description:** Add files to staging area. These files will be committed 
**Variations:** `.` for all, `-A` for all files, `-u` for files that were tracked
**Example:** 
`git add .` # all files in the current folder 
`git add /path/to/file` # Stage file identified by path
`git add -A` # stages all changes in the repository, regardless of folder
`git add -u` #Stage all tracked files (files which were previously tracked and modified)


## git commit
**Description:** Commit changes after staging them
**Variations:** `-m` , `--amend` 
**Example:**
`git commit -m “Message”` # An empty message aborts the commit command
`git commit --amend -m “Changed message”` # to change the commit message of the previous commit


## git reset
**Description:** Undo changes / unstage files / go back to commit
**Variations:** `--soft`, `--hard`  
**Example:** 
`git reset <filename>` # git reset without any arguments : unstages all added files but preserves all changes.
`git reset --soft <commit_hash>` # undo commit, files are staged , keeps changes ,takes branch back to specified commit
`git reset --soft HEAD~N` # undo most recent N commits , files staged, keeps changes
`git reset --hard <commit hash>` # unstages all added files and deletes all changes you made after specified commit
`git reset --hard HEAD~N` # unstage all added files and deletes all changes in the last N commits



## git restore
**Description:** Restore files to an earlier state 
**Variations:** `--staged` to unstage  
**Example:** 
`git restore file.go` # to undo changes since last commit, only possible when unstaged
`git restore –-staged test.txt` # to unstage added files


## git log
**Description:** Show commit history  
**Variations:** `--oneline`, `--graph`  
**Example:** 
`git log` # full log of the corresponding branch or use filtering arguments such as –-after, –author or -n
`git log –-after=“2022-1-1”` # shows commits made after a specified date 
`git log –-after="yesterday"` # shows after relative time
`git log -n 5` # shows last 10 commits
`git log –-author=“Jose”` # shows commits made by a specific author 
`git log --oneline` # shows short commit history , in one line per commit
`git log --graph` # shows commit history as a branch graph, helps to visualise
`git log --oneline -n 5` # can combine multiple filters



## git diff
**Description:** Show changes between versions of files  
**Variations:** `--staged` for staged diff  
**Example:** 
`git diff ` # list all changes since the last commit (unstaged changes)
`git diff <commitID>` # all changes since specified commit ID
`git diff HEAD~N` # all changes since last N commits
`git diff –-cached` # for staged files 
`git diff <filename>` # for a single file 
`git diff <commitID1> <commitID2>` # compare between commits
`git diff <commitID1> <commitID2> <filename>` # compare particular file between commits 
`git diff <branch1> <branch2>`  # compare between branches


## git branch
**Description:** List,create or delete branches
**Variations:** `-d` delete , `-m` rename
**Example:** 
`git branch` # lists all local branches
`git branch -r` # show remote branches
`git branch -a` # show all (remote+local)
`git branch newbranch` # creates new branch with specified name, does not switch to it
`git branch -d branch-name` # delete specified branch
`git branch -m old-name new-name` # rename branch


## git checkout
**Description:** Switch branch  
**Variations:** `-b` create + switch  
**Example:** 
`git checkout -b branchname` # create a new branch of it does not exist and go to it
`git checkout branchname`  # go to specified branch
`git checkout commitId` # to switch to a previous commit 
`git checkout -m name` # merges the changes of the current branch into name and switches to name
`git checkout filename` # to undo changes since the last commit


## git switch
**Description:** Switch branches (newer command)  
**Variations:** `-c` create branch  
**Example:** 
`git switch name ` # switch to branch name
`git switch -c name`  # to create the branch name if it does not exist and switch to it
`git switch -d commitId`  # to switch to a previous commit
`git switch -m name`  merges the changes of the current branch into name and switches to name


## git push
**Description:** Push new commits to remote repository
**Variations:** `-u` set upstream  
**Example:** 
`git push -u origin branch-name`  # use for the first push in the branch 
`git push` # to push to the branch you are currently on and the remote repository defined in .git/config
`git push origin` # to push to a different remote repository 
`git push origin main` # to push the desired branch (not one you are currently on)
`git push origin main:test` #  push to the targetBranch (test) from the sourceBranch(main)


## git merge
**Description:** Merge branches  
**Example:** 
`git merge branch`  # to merge specified branch into the current branch 
`git merge branch1 branch2` # merges branch1 and branch2 on the current branch (a new commit is created)
`git merge -s strategy branch` # to define the merging strategy

## git fetch
**Description:** Fetch changes from the remote repository (does not update head) 
**Example:** 
`git fetch` # get the new commits from the branch you are currently on and the remote repository defined in .git/config
`git fetch origin` # update from a different remote repository 
`git fetch origin main` # to get the desired branch (not one you are currently on)
`git fetch origin main:test` # to get the sourceBranch (main) into the targetBranch (test)


## git pull
**Description:**  Update local version using remote version (git fetch +  git merge)
**Example:** 
`git pull` # to pull from the branch you are currently on and the remote repository defined in .git/config
`git pull origin` # to pull from a different remote repository 
`git pull origin main` # to pull the desired branch ( not one you are currently on)
`git pull origin main:test` # to pull the sourceBranch (main) into the targetBranch (test)


## git stash
**Description:** Temporarily save changes  
**Variations:** `pop`, `list`, `apply`  
**Example:** 
`git stash` # to add a new stash entry with the current modifications and reset your state to the current HEAD
`git stash list` # to get all stash entries
`git stash show` # to visualize the changes (diff)
`git stash pop` # to pop the first element of git stash list
`git stash pop` # stash@{i} to get the ith element of the stack


## git rebase 
**Description:** Rewrite commit history 
**Example:**
`git rebase -i HEAD~n` # to rebase the last n commits in the interactive mode
`git rebase main` # to rebase main on the current branch


## git remote
**Description:** Manage remotes  
**Variations:** `-v` view, `add`
**Example:** 
`git remote add <name> <url>` # creates convenient shortname for a remote that can be used in other commands
`git remote -v` # to list all the remotes



