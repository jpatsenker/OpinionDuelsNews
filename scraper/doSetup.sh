
# get the directory of the script
# will use this to set the gopath, then the gopath to set the path
DIR=$(cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )

# set the GOPATH
export GOPATH=$DIR

# set the path for executables
# this should point to (current directory)/bin
export PATH="$PATH:$GOPATH/bin"


