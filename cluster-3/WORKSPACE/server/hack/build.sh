#!/usr/bin/env bash

#  \
#  \\,
#   \\\,^,.,,.                     Zero to Hero
#   ,;7~((\))`;;,,               <zerotohero.dev>
#   ,(@') ;)`))\;;',    stay up to date, be curious: learn
#    )  . ),((  ))\;,
#   /;`,,/7),)) )) )\,,
#  (& )`   (,((,((;( ))\,

go build main.go
if [ $? -ne 0 ]; then
	echo "Poop"
	exit 1
fi

docker build . -t localhost:32000/greeter-server:demo
if [ $? -ne 0 ]; then
	echo "Poop"
	exit 1
fi

docker push localhost:32000/greeter-server:demo
if [ $? -ne 0 ]; then
	echo "Poop"
	exit 1
fi

echo "Everything is awesome!"
