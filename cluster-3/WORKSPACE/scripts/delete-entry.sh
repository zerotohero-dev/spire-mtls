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
3fd579b1-2039-4a67-90fe-97de1b2fd5bb
