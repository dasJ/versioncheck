#!/bin/bash

upstream="$1"
module="$2"
old="$3"
new="$4"
shift 4

mail -s "[versioncheck] Package update for $upstream" "root@localhost" <<EOF
Hello,

the $module module found that your package $upsteram got updated:

$old -> $new

Your custom tags are:

$@

Thank you for maintaining your package!
EOF
