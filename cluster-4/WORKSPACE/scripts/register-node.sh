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
	-spiffeID spiffe://cluster4.demo/ns/spire/sa/spire-agent \
	-selector k8s_sat:cluster:cluster4 \
	-selector k8s_sat:agent_ns:spire \
	-selector k8s_sat:agent_sa:spire-agent \
	-node
