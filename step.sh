#!/bin/bash

file_path="`pwd`/${file_name}"
echo "$file_content" > $file_path
envman add --key GENERATED_TEXT_FILE_PATH --value $file_path
exit 0
