#!/bin/sh

cwd=$(pwd)
path_to_watch="~/test_dir/path_to_watch/"

inotifywait -m \
  --timefmt '%d/%m/%y %H:%M' --format '%T %w %f' \
  -e close $path_to_watch |
while read -r date time dir file; do
    changed_abs=${dir}${file}
    changed_rel=${changed_abs#"$cwd"/}

    # run the script here
    file-org ~/test_dir/org_this ~/test_dir/here
    echo "At ${time} on ${date}, file $changed_abs was organized!" >&2
done
