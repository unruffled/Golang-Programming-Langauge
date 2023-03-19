This is an overview of Todd McLeod's Learn How to Code: Google's Go (golang) Programming Language.
Go was created at Google under the guidance of Rob Pike (Unix), Robert Griesmer (who studied under the creator of Pascal), and Ken Thompson (co-author of the C programming langauge).
Go creates compiled programs.
In 2005/2006, the first dual core processors were released. Go was created to take advantage of dual core processors and utilize concurrent computing and parallelsim.
YouTube has been rewritten in Go.


The Terminal
GUI == graphical user interface
CLI == Command line interface
Terminal == text input/output envrionment; console = physical terminal
unix / linux / mac
shell / bash / terminal
Windows
command prompt / windows command prompt / cmd / DOS prompt


Basic shell commands
pwd
ls
ls -la == shows the permissions owner, group, and world
mkdir
touch
cat
clear
chmod
emv
rm
grep
chmod
rm


sha256 allows us to confirm the integrity of data

package management
go modules

Go Workspace
Go is opinionated with how programmers need to do things
Backed by research numbers and data

Project Aristotle talks about what makes a team work well
Go is opinionated about how code is structured for effeciency

The workspace is one of the convetions where the machine needs to be setup in a particular way

Go workspace needs to have three folders within it bin, pkg, and src
bin == binaries
pkg == archived files
src == amespacing and package management (e.g., github.com and folders with projects and repos)

Go mantra
effecient execution
effecient compliation
ease of programming

namespacing
go get can be used to install packages
GOPATH points to workspace
GOROOT points to the binary installation of Go

Setup workspace
go en reveals the go path and go environment settings

Environment variables are varaibles for your environment
Varaibles that can be set into the computer that the computer can then use

system variables are for everybody

To run a go program within the terminal, one would type "go run [filename].go"

What is the difference between a Go package and a Go module?
A package is a directory of .go files.
A module is a collection of Go packages with dependencies and versioning built-in.


go get -d [location]
go help
go env
go help command
go fmt
go fmt updates the format of code

go fmt ./... formats everything in the directory and down the directory
go run needs the filename, eg, go run main.go
go build builds teh fiel, reports the errors if any, puts the executable in the current oflder
go install compiles the program, names the executable the folder name holding hte code

GitHub
git is version control software
git is the de facto industry standards for version control software
GitHub is the most popular website for version control software
Each repo has code in them
GitHub keeps track of the changes over time

go get grabs a repo

git clone ssh

package management
node package manager (npm)

Modules help manage packages