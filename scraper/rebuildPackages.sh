# scrip to rebuild packages if there are any config problems
# if you add a package to the code, also add it in here

# Net library, gives us html parser
# first remove the folder
rm -r src/golang.org/x
go get golang.org/x/net/html

# mock library, gives us mock-test codes
rm -r src/github.com/golang/mock
go get github.com/golang/mock/gomock
go get github.com/golang/mock/mockgen
