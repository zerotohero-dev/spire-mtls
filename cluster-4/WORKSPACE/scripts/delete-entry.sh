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
    /opt/spire/bin/spire-server entry delete -entryID \
dc4f8c96-5afb-4edf-811f-bfbd7c3cd8e1
