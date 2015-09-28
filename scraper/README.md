Setup
=========

To clone repo: git clone https://github.com/jpatsenker/Opinionated.git

Use bash script to setup enviroment variables: . doSetup.sh

If there is a problem with any of the packages, fix them with rebuildPackages.sh

####Getting the repo

1) clone (see above)  
2) create local copy of develop branch (git checkout -b develop origin/develop)   
3) create your branch off of develop (git checkout -b (your-name)-develop)  

####Pushing

1) Push your branch to the repo with git push origin (your-name)-develop  
2) Message group chat about merging  

####Once your updates get merged

1) change to develop branch (git checkout develop)  
2) pull from develop (git pull origin develop)  
3) merge develop back into your branch to get any updates (git checkout (your-name)-develop; git merge develop)  
