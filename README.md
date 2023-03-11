# app-get

## What problem does app-get solves
Not all software installed on Linux systems comes from package manager, snap or flathub. App-get helps to update software downloaded from Github in form of deb package (other pages or appimage support may come later).

## How to use it
You can download executable or deb package.
It is a good idea to set add app-get to updates automatically.
> sudo app-get install afelinczak/app-get

Other applications can be installed in same maner via app-get.
### For example to install `micro` from github type. 
> sudo app-get install zyedidia/micro

It will install amd64 version of deb and create `apps.json` to track installed apps.
List of apps is stored in /opt/app-get/apps.json

### Now to update all the apps:
> sudo app-get update

### To remove application from update list:
> sudo app-get remove zyedidia/micro


## Disclaimer
This my first project in Go, so codebase may be not perfect.
At the moment it is more POC than something useful.
