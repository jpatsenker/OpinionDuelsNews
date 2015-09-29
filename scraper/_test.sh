#!/bin/bash

DIR=$(cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )

echo $DIR
##### Functions

function smoke()
{

	go test  scraper/config

	printf  "\nSmoke test run"
}	#end of smoke

##### Main
case "$1" in
    "smoke" )	
		smoke
		;;
    * )
		while true; do
		read -p "No test input given.\n Please enter what type of test you would like, or type q to quit: " q
		case $q in 
			smoke ) 
				echo ""
				smoke
				exit;;
			[Qq]* ) break;;
		esac
	done
		;;
esac
