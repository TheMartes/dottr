# Dottr — suckless dotfiles manager

## About this project
Dottr is aiming to solve users most important files. Config files a.k.a. "dotfiles". In curent age when we're using thousands of
different application using their own dotfiles, it's hard to track what happens to them, which one are synced and, which not.

There are ways people are now trying to solve this issue by either using *symlinks* and managing their own repo or using tools
like `yadm` or `stow`, which in my opinion are not that intuitive.

Dottr should provide ways to sync, manage and try different dotfiles just by using simple commands and config file. So why don't you take a look?

## Usage
`dottr pull` — will pull the latest version from repo
`dottr push` — will push the latest version to repo from your computer
`dottr init` — will create dottr repo from pre-defined config

## Config (yaml)
```yaml
---
repo_dir: <dottr dir_name in $HOME>
repo_url: <git url for repo>
include_config: <bool> # This will tell dottr if .dottr.yaml should be present in repo itself

folders_to_sync: # Array of folders with path relative to your $HOME
    -   # Entry
        path: <path to folder>
        exclude: 
            - <files or subfolders to exclude from this entry>
files_to_sync: # Array of paths to single files relative to your $HOME
    - <path to file>
```

### Disclaimer
- This project is unrelated to software that suck less -> https://suckless.org/
