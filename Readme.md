Following environment variables have to be set 

a) GOPATH -> this is the path of the work space which is till the "go poc" project is placed <PROJECT_WKSPACE>/"go poc"
b) GOBIN -> $GOPATH/bin

Following are the commands :

a) go run src/*.go -> all the go files which belong to the main package are executed.

b) go run src/<project_name>/*.go -> all the go files which belong to the main package residing the project_name directory is executed

c) go install src/calculator/*.go will create calculator.exe for Windows and will be stored in GOBIN folder.

 