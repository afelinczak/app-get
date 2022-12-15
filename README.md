# app-get

## What problem does app-get solves
Not all software installed on Linux systems comes from package manager, snap or flathub. App-get helps to update software downloaded from Github in form of deb package (other pages or appimage support may come later).

## How to use it
There is no binaries provided as yet, so clone and compile.  
First, application needs to be installed via app-get.  
### For example to install `micro` from github type. 
> sudo app-get install zyedidia/micro

It will install amd64 version of deb and create `apps.json` to track installed apps.
### Now to update:
> sudo app-get update


## Disclaimer
This my first project in Go, so codebase may be not perfect.
At the moment it is more POC than something useful.