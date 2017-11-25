#!/bin/bash
# Usage watch ./generate-dummy-data.sh ASD

names=(Anna Bob Cilla Dylan)


to=${names[$RANDOM % ${#names[@]} ]}
from=${names[$RANDOM % ${#names[@]} ]}

messages=("Hello!" "Who are you?" "I am $to" "My name is $to" "Goodbye")
msg=${messages[$RANDOM % ${#messages[@]} ]}

echo '{
"to" : "'"$to"'",
"from" : "'"$from"'",
"msg" : "'"$msg"'",
"test-id" : "'"$1"'"
}'
