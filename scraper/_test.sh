#!/bin/bash

DIR=$(cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )


##### Functions

function smoke()
{
	go test  scraper/config -run TestGetNEsted -v
	go test  scraper/config -run TestMultiRead -v
	go test scraper/fetcher -run TestTricky -v
	go test scraper/fetcher -run TestSchedulableRSS -v
	go test scraper/fetcher -run TestSchedulableRSSMock -v
	go test scraper/fetcher -run MakeTestSchedulable -v
	go test scraper/scheduler -run TestScheduler -v
	printf  "\nSmoke test run"
}	#end of smoke

##### Main
case "$1" in
    "smoke" )	
		smoke
		;;
    * ) echo "No test input given."
		while true; do
		read -p "Please enter what type of test you would like, or type q to quit: " q
		case $q in 
			smoke ) 
				echo ""
				smoke
				exit;;
			[Hh] ) echo "Possible commands are:
smoke: runs smoke test
Hh: shows commands
Qq: quits script
";;
			[Qq] ) break;;
		esac
	done
		;;
esac
