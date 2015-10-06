#!/bin/bash

DIR=$(cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )


##### Functions

function smoke()
{
	go test  scraper/config -run TestGetNested -v
	go test  scraper/config -run TestMultiRead -v
	go test scraper/fetcher -run TestTricky -v
	go test scraper/fetcher -run TestSchedulableRSS -v
	go test scraper/fetcher -run TestSchedulableRSSMock -v
	go test scraper/fetcher -run MakeTestSchedulable -v
	go test scraper/scheduler -run TestScheduler -v
	printf  "\nSmoke test run"
}	#end of smoke
function full()
{
	go test scraper/config -v
	go test scraper/fetcher -v
	go test scraper/scheduler -v

}
function specific()
{	
	for (( i = 2; i <= $#; i++ )); do
	go test scraper/scheduler -run $i -v
	go test  scraper/config -run $i -v
	go test scraper/fetcher -run $i -v	
	done
}
##### Main
case "$1" in
    "smoke" )	
		smoke
		;;
	"full" )
		full
		;;
	"specific" )
		specific $@
		;;
    * ) echo "No test input given."
		while true; do
		read -p "Please enter what type of test you would like, or type q to quit: " q
		case $q in 
			smoke ) 
				echo ""
				smoke
				exit;;
				full )
				echo ""
				full
				exit;;
				specific* )
				specific $q
				exit;;
			[Hh] ) echo "Possible commands are:
smoke: runs smoke test
full: runs full test
specific  arg: runs specified tests 
Hh: shows commands
Qq: quits script
";;
			[Qq] ) break;;
		esac
	done
		;;
esac
