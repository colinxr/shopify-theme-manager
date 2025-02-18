package commands

import "os/exec"

// execCommand is declared at package level for mocking in tests
var execCommand = exec.Command 