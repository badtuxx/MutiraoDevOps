package main

import(
  "os/exec"
  "syscall"
)


func main() {
    var qoses = [3]string{"0", "1", "2"}
    echo, _ := exec.LookPath("echo")

    for  _,qos := range qoses {
        syscall.ForkExec(echo, []string{"Tua mãe", qos}, &syscall.ProcAttr{})
    }
}
