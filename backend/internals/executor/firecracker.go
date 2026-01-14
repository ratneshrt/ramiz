package executor

import (
	"bytes"
	"context"
	"encoding/json"
	"net"
	"net/http"
	"os"
	"os/exec"
	"time"
)

type FirecrackerExecutor struct {
	KernelPath string
	RootFSPath string
	CodeFSPath string
}

func (f *FirecrackerExecutor) Run() error {
	socketPath := "/tmp/firecracker.sock"
	os.Remove(socketPath)

	cmd := exec.Command(
		"firecracker",
		"--api-sock", socketPath,
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		return err
	}

	time.Sleep(300 * time.Millisecond)
	client := firecrackerClient(socketPath)

	put(client, "/machine-config", map[string]interface{}{
		"vcpu_count":   1,
		"mem_size_mib": 128,
		"ht_enabled":   false,
	})

	put(client, "/boot-source", map[string]interface{}{
		"kernel_image_path": f.KernelPath,
		"boot_args":         "console=ttyS0 reboot=k panic=1, init=/init",
	})

	put(client, "/drives/rootfs", map[string]interface{}{
		"drive_id":       "rootfs",
		"path_on_host":   f.RootFSPath,
		"is_root_device": true,
		"is_read_only":   false,
	})

	put(client, "/drives/codefs", map[string]interface{}{
		"drive_id":       "codefs",
		"path_on_host":   f.CodeFSPath,
		"is_root_device": false,
		"is_read_only":   false,
	})

	put(client, "/actions", map[string]string{
		"action_type": "InstanceStart",
	})

	cmd.Wait()
	return nil
}

func firecrackerClient(socketPath string) *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				return net.Dial("unix", socketPath)
			},
		},
	}
}

func put(client *http.Client, path string, body interface{}) {
	b, _ := json.Marshal(body)
	req, _ := http.NewRequest("PUT", "http://localhost"+path, bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	client.Do(req)
}
