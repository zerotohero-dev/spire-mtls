#!/usr/bin/env bash

#  \
#  \\,
#   \\\,^,.,,.                     Zero to Hero
#   ,;7~((\))`;;,,               <zerotohero.dev>
#   ,(@') ;)`))\;;',    stay up to date, be curious: learn
#    )  . ),((  ))\;,
#   /;`,,/7),)) )) )\,,
#  (& )`   (,((,((;( ))\,


microk8s kubectl exec -n spire spire-server-0 -- \
	/opt/spire/bin/spire-server bundle set -format spiffe \
	-id spiffe://cluster3.demo \
	-path /run/spire/data/peer.json
	
