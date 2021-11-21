package ping

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/ycd/dstp/pkg/common"
)

func RunTest(ctx context.Context, addr common.Address, count int, timeout int) (common.Output, error) {
	return runPing(ctx, addr, count, timeout)
}

func runPing(ctx context.Context, addr common.Address, count int, timeout int) (common.Output, error) {
	var output string

	pinger, err := createPinger(addr.String())
	if err != nil {
		return "", err
	}

	pinger.Count = count
	if timeout == -1 {
		pinger.Timeout = time.Duration(2*count) * time.Second
	} else {
		pinger.Timeout = time.Duration(timeout) * time.Second
	}
	err = pinger.Run()
	if err != nil {
		if out, err := runPingFallback(ctx, addr, count, timeout); err == nil {
			output += out.String()
		} else {
			return "", fmt.Errorf("failed to run ping: %v", err.Error())
		}
	} else {
		stats := pinger.Statistics()
		if stats.PacketsRecv == 0 {
			output += "no response"
		} else {
			output += joinS(joinC(stats.AvgRtt.String()))
		}
	}

	return common.Output(output), nil
}

// runPingFallback executes the ping command from cli
// Currently fallback is not implemented for windows.
func runPingFallback(ctx context.Context, addr common.Address, count int, timeout int) (common.Output, error) {
	args := fmt.Sprintf("-c %v -t %v", count, timeout)
	command := fmt.Sprintf("ping %s %s", args, addr.String())

	if runtime.GOOS == "windows" {
		return common.Output(""), fmt.Errorf("fallback is not implemented for windows")
	}

	out, err := executeCommand("bash", command)
	if err != nil {
		return common.Output(""), err
	}

	po, err := parsePingOutput(out)
	if err != nil {
		return common.Output(""), err
	}

	return common.Output(po.AvgRTT + "ms"), nil
}

func executeCommand(shell, command string) (string, error) {
	var errb bytes.Buffer
	var out string

	cmd := exec.Command(shell, "-c", command)
	cmd.Stderr = &errb
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Printf("got error while tracing pipe: %v", err)
	}
	err = cmd.Start()
	if err != nil {
		return "", err
	}

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		out += scanner.Text() + "\n"
	}

	if err := cmd.Wait(); err != nil {
		return "", fmt.Errorf("got error: %v, stderr: %v", err, errb.String())
	}

	return out, nil
}

type pingOutput struct {
	PacketLoss        string
	PacketReceived    string
	PacketTransmitted string
	MinRTT            string
	AvgRTT            string
	MaxRTT            string
}

var (
	RequestTimeoutError = fmt.Errorf("requests timed out")
	PacketLossError     = fmt.Errorf("timeout error: 100.0%% packet loss")
)

// parsePingOutput parses the output of ping by parsing the stdout
// example output:
//
// ping -c 3 jvns.ca
// PING jvns.ca (104.21.91.206): 56 data bytes
// 64 bytes from 104.21.91.206: icmp_seq=0 ttl=58 time=14.468 ms
// 64 bytes from 104.21.91.206: icmp_seq=1 ttl=58 time=14.450 ms
// 64 bytes from 104.21.91.206: icmp_seq=2 ttl=58 time=14.683 ms
//
// --- jvns.ca ping statistics ---
// 3 packets transmitted, 3 packets received, 0.0% packet loss
// round-trip min/avg/max/stddev = 14.450/14.534/14.683/0.106 ms
func parsePingOutput(out string) (pingOutput, error) {
	var po pingOutput

	lines := strings.Split(out, "\n")

	for _, line := range lines {
		switch {
		case strings.Contains(line, "packets transmitted"):
			arr := strings.Split(line, ",")
			fmt.Println(arr)
			if len(arr) != 3 {
				continue
			}

			po.PacketTransmitted, po.PacketReceived, po.PacketLoss = arr[0], arr[1], arr[2]

		case strings.Contains(line, "round-trip min/avg/max"):
			l := strings.ReplaceAll(line, " = ", " ")
			arr := strings.Split(l, " ")

			if len(arr) != 4 {
				continue
			}

			rttArr := strings.Split(arr[2], "/")
			if len(rttArr) != 4 {
				continue
			}

			po.MinRTT, po.AvgRTT, po.MaxRTT = rttArr[0], rttArr[1], rttArr[2]
		}
	}

	if po.MinRTT == "" && po.AvgRTT == "" && po.MaxRTT == "" {
		return po, RequestTimeoutError
	}

	if po.PacketLoss == "100.0% packet loss" {
		return po, PacketLossError
	}

	return po, nil
}

func joinC(args ...string) string {
	return strings.Join(args, ",")
}

func joinS(args ...string) string {
	return strings.Join(args, " ")
}
