make VERSION=`git describe --exact-match --tags $(git log -n1 --pretty='%h')` update-homebrew