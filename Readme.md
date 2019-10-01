Following environment variables have to be set 

a) GOPATH -> this is the path of the work space which is till the "go poc" project is placed <PROJECT_WKSPACE>/"go poc"
b) GOBIN -> $GOPATH/bin

Following are the commands :

go run src/*.go -> all the go files which belong to the main package are executed.
go run src/<xyz>/*.go -> all the go files which belong to the main package residing the xyz directory is executed
go install src/calculator/*.go will create calculator.exe for Windows and will be stored in GOBIN folder.

 