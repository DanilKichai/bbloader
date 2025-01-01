#!/usr/bin/env bash

chainload_file () {
    local FILE="$1"

    if ! exec "$FILE"; then
        echo "Failed to exec file: \"$FILE\"" 1>&2

        return 1
    fi
}

chainload_uri () {
    local URI="$1"
    local FILE="/archshell/downloads/$RANDOM"

    if ! curl \
        --output "$FILE" \
        --silent \
        --fail \
        "$URI"
    then
        echo "Failed to download the file via URI: \"$URI\"!" 1>&2

        return 1
    fi

    if ! chmod +x "$FILE"; then
        echo "Failed to chmod the file: \"$FILE\"!" 1>&2

        return 1
    fi

    ! chainload_file "$FILE" && \
        return 1
}
