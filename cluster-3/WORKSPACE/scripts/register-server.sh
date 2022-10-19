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
	/opt/spire/bin/spire-server entry create \
	-spiffeID spiffe://cluster3.demo/ns/default/sa/default/app/greeter-server \
	-parentID spiffe://cluster3.demo/ns/spire/sa/spire-agent \
	-selector k8s:ns:default \
	-selector k8s:sa:default \
	-selector k8s:pod-label:app:greeter-server \
	-federatesWith "spiffe://cluster4.demo"

