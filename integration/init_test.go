package integration

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
)

const packagePath string = "github.com/digitalocean/doctl/cmd/doctl"

var (
	suite           spec.Suite
	builtBinaryPath string
)

func TestAll(t *testing.T) {
	suite.Run(t)
}

func TestMain(m *testing.M) {
	specOptions := []spec.Option{
		spec.Report(report.Terminal{}),
		spec.Random(),
		spec.Parallel(),
	}

	suite = spec.New("acceptance", specOptions...)
	suite("account/get", testAccountGet)
	suite("account/ratelimit", testAccountRateLimit)
	suite("auth/init", testAuthInit)

	tmpDir, err := ioutil.TempDir("", "acceptance-doctl")
	if err != nil {
		panic("failed to create temp dir")
	}
	defer os.RemoveAll(tmpDir) // yes, this is best effort only

	builtBinaryPath = filepath.Join(tmpDir, path.Base(packagePath))
	if runtime.GOOS == "windows" {
		builtBinaryPath += ".exe"
	}

	// tried to use -mod=vendor but it blew up
	cmd := exec.Command("go", "build", "-o", builtBinaryPath, packagePath)
	cmd.Env = append(os.Environ(), "CGO_ENABLED=0")
	output, err := cmd.CombinedOutput()
	if err != nil {
		panic(fmt.Sprintf("failed to build doctl: %s", output))
	}

	code := m.Run()

	os.Exit(code)
}
