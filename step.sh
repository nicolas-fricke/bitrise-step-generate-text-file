#!/bin/bash

file_path="`pwd`/${file_name}"
eval "echo \"$file_content\"" > $file_path
printf "> cat $file_path\n\n"
cat $file_path
printf "\n\nexporting GENERATED_TEXT_FILE_PATH=$file_path\n"
envman add --key GENERATED_TEXT_FILE_PATH --value $file_path
exit 0
