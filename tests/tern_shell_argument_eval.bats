#!/usr/bin/env bats

@test "shell argument evaluation" {
    # Be aware of argument evaluation.
    # Both expressions will be evaluated.
    res=$(tern -b FALSE "$(echo yes)" "$(echo no)")
    [ "$res" == "no" ]
}
