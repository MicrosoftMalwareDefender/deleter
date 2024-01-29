package Deleter

import (
    "fmt"
    "os"
    "os/exec"
    "runtime"
)

func Deleter() {
    // Get the PID of the current process
    pid := os.Getpid()

    // Kill the current process
    if err := killProcess(pid); err != nil {
        fmt.Println("Error killing process:", err)
        return
    }

    // Delete the executable
    if err := deleteExecutable(); err != nil {
        fmt.Println("Error deleting executable:", err)
        return
    }

    fmt.Println("Process killed and executable deleted successfully.")
}

// killProcess kills the process with the given PID.
func killProcess(pid int) error {
    // Use platform-specific command to kill the process
    var cmd *exec.Cmd
    switch runtime.GOOS {
    case "windows":
        cmd = exec.Command("taskkill", "/F", "/T", "/PID", fmt.Sprintf("%d", pid))
    default:
        cmd = exec.Command("kill", "-9", fmt.Sprintf("%d", pid))
    }
    return cmd.Run()
}

// deleteExecutable deletes the executable file.
func deleteExecutable() error {
    // Get the path of the executable
    executable, err := os.Executable()
    if err != nil {
        return err
    }
    // Delete the executable file
    return os.Remove(executable)
}
