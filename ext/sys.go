package ext

import (
	"os"
	"os/exec"
	"runtime"
	"github.com/Igazine/hank-go"
)

type SysExtension struct{}

func (e *SysExtension) Name() string {
	return "SysExtension"
}

func (e *SysExtension) GetTasks() map[string]hank.NativeFunc {
	tasks := make(map[string]hank.NativeFunc)

	// host
	tasks["host_cwd"] = func(args []hank.Value, ctx hank.ExecutionContext) hank.Value {
		cwd, _ := os.Getwd()
		return hank.Value{Type: hank.TypeString, String: cwd}
	}
	tasks["host_isRoot"] = func(args []hank.Value, ctx hank.ExecutionContext) hank.Value {
		if os.Getuid() == 0 {
			return hank.Value{Type: hank.TypeNumber, Number: 1}
		}
		return hank.Value{Type: hank.TypeVoid}
	}
	tasks["host_pid"] = func(args []hank.Value, ctx hank.ExecutionContext) hank.Value {
		return hank.Value{Type: hank.TypeNumber, Number: float64(os.Getpid())}
	}

	// os
	tasks["os_type"] = func(args []hank.Value, ctx hank.ExecutionContext) hank.Value {
		return hank.Value{Type: hank.TypeString, String: runtime.GOOS}
	}
	tasks["os_name"] = func(args []hank.Value, ctx hank.ExecutionContext) hank.Value {
		return hank.Value{Type: hank.TypeString, String: runtime.GOOS}
	}
	tasks["os_arch"] = func(args []hank.Value, ctx hank.ExecutionContext) hank.Value {
		return hank.Value{Type: hank.TypeString, String: runtime.GOARCH}
	}
	tasks["os_memory"] = func(args []hank.Value, ctx hank.ExecutionContext) hank.Value {
		m := make(map[string]hank.Value)
		m["total"] = hank.Value{Type: hank.TypeNumber, Number: 1024} // Mock
		m["free"] = hank.Value{Type: hank.TypeNumber, Number: 512}   // Mock
		m["used"] = hank.Value{Type: hank.TypeNumber, Number: 512}   // Mock
		return hank.Value{Type: hank.TypeMap, Map: m}
	}
	tasks["os_cpu"] = func(args []hank.Value, ctx hank.ExecutionContext) hank.Value {
		return hank.Value{Type: hank.TypeNumber, Number: float64(runtime.NumCPU())}
	}

	// fs
	tasks["fs_exists"] = func(args []hank.Value, ctx hank.ExecutionContext) hank.Value {
		if len(args) == 0 || args[0].Type != hank.TypeString {
			return hank.Value{Type: hank.TypeVoid}
		}
		if _, err := os.Stat(args[0].String); err == nil {
			return hank.Value{Type: hank.TypeNumber, Number: 1}
		}
		return hank.Value{Type: hank.TypeVoid}
	}
	tasks["fs_read"] = func(args []hank.Value, ctx hank.ExecutionContext) hank.Value {
		if len(args) == 0 || args[0].Type != hank.TypeString {
			return hank.Value{Type: hank.TypeVoid}
		}
		content, err := os.ReadFile(args[0].String)
		if err != nil {
			return hank.Value{Type: hank.TypeVoid}
		}
		return hank.Value{Type: hank.TypeString, String: string(content)}
	}
	tasks["fs_write"] = func(args []hank.Value, ctx hank.ExecutionContext) hank.Value {
		if len(args) < 2 || args[0].Type != hank.TypeString || args[1].Type != hank.TypeString {
			return hank.Value{Type: hank.TypeVoid}
		}
		err := os.WriteFile(args[0].String, []byte(args[1].String), 0644)
		if err != nil {
			return hank.Value{Type: hank.TypeVoid}
		}
		return hank.Value{Type: hank.TypeNumber, Number: 1}
	}
	tasks["fs_deleteFile"] = func(args []hank.Value, ctx hank.ExecutionContext) hank.Value {
		if len(args) == 0 || args[0].Type != hank.TypeString {
			return hank.Value{Type: hank.TypeVoid}
		}
		err := os.Remove(args[0].String)
		if err != nil {
			return hank.Value{Type: hank.TypeVoid}
		}
		return hank.Value{Type: hank.TypeNumber, Number: 1}
	}
	tasks["fs_stat"] = func(args []hank.Value, ctx hank.ExecutionContext) hank.Value {
		if len(args) == 0 || args[0].Type != hank.TypeString {
			return hank.Value{Type: hank.TypeVoid}
		}
		info, err := os.Stat(args[0].String)
		if err != nil {
			return hank.Value{Type: hank.TypeVoid}
		}
		m := make(map[string]hank.Value)
		m["size"] = hank.Value{Type: hank.TypeNumber, Number: float64(info.Size())}
		m["isDir"] = hank.Value{Type: hank.TypeVoid}
		if info.IsDir() {
			m["isDir"] = hank.Value{Type: hank.TypeNumber, Number: 1}
		}
		m["mtime"] = hank.Value{Type: hank.TypeNumber, Number: float64(info.ModTime().UnixNano() / 1e6)}
		return hank.Value{Type: hank.TypeMap, Map: m}
	}

	// proc
	tasks["proc_run"] = func(args []hank.Value, ctx hank.ExecutionContext) hank.Value {
		if len(args) == 0 || args[0].Type != hank.TypeString {
			return hank.Value{Type: hank.TypeVoid}
		}
		cmdName := args[0].String
		var cmdArgs []string
		if len(args) > 1 && args[1].Type == hank.TypeArray {
			for _, a := range *args[1].Array {
				cmdArgs = append(cmdArgs, hank.ValueToString(a))
			}
		}
		cmd := exec.Command(cmdName, cmdArgs...)
		out, err := cmd.CombinedOutput()
		exitCode := 0
		if err != nil {
			if exitErr, ok := err.(*exec.ExitError); ok {
				exitCode = exitErr.ExitCode()
			} else {
				exitCode = -1
			}
		}
		m := make(map[string]hank.Value)
		m["code"] = hank.Value{Type: hank.TypeNumber, Number: float64(exitCode)}
		m["stdout"] = hank.Value{Type: hank.TypeString, String: string(out)}
		m["stderr"] = hank.Value{Type: hank.TypeString, String: ""} // CombinedOutput merges them
		return hank.Value{Type: hank.TypeMap, Map: m}
	}

	return tasks
}
