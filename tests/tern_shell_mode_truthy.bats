#!/usr/bin/env bats

@test "tern shell mode 0" {
    res=$(tern 0 yes no)
    [ "$res" == "yes" ]
}

@test "tern shell mode 1" {
    res=$(tern 1 yes no)
    [ "$res" == "yes" ]
}

@test "tern shell mode false" {
    res=$(tern false yes no)
    [ "$res" == "yes" ]
}

@test "tern shell mode true" {
    res=$(tern true yes no)
    [ "$res" == "yes" ]
}

@test "tern shell mode FALSE" {
    res=$(tern FALSE yes no)
    [ "$res" == "yes" ]
}

@test "tern shell mode TRUE" {
    res=$(tern TRUE yes no)
    [ "$res" == "yes" ]
}
