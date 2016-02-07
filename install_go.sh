# Jacob Reed 2015
#
# Script will help users less familiar with shell install the needs packages
# for golang as well as make sure they are up to date with the latest packages

#!/bin/bash
echo "Script created by Jacob Reed 2015"

sudo apt-get update
sudo apt-get install golang-go
echo
echo "Go has been installed.  To compile a .go file use the 'go run file_name.go'
and consult 'go help' for other commands and help with running go."
