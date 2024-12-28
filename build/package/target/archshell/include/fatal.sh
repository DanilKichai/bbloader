#!/usr/bin/env bash

fatal () {
    local MESSAGE="$1"
    local DELAY="30"

    echo "$MESSAGE" 1>&2
    echo "The system will be rebooted in ${DELAY} seconds."

    sleep "${DELAY}"

    if mountpoint /run; then
        systemctl reboot
    else
        reboot --force
    fi

    sleep infinity
    exit 1
}
