#!/usr/bin/env bash

isLeap() {
    for arg in "$@"; do
        if [ $((arg % 4)) -eq 0 ]; then
            if [ $((arg % 100)) -eq 0 ]; then
                if [ $((arg % 400)) -eq 0 ]; then
                    echo "True"
                else
                    echo "False"
                fi
            else
                echo "True"
            fi
        else
            echo "False"
        fi
    done
}

isLeap "$@"
