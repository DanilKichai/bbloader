#!/usr/bin/env bash

command_line () {
    ! exec bash && \
        return 1
}
