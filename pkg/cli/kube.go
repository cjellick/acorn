package cli

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	cli "github.com/acorn-io/runtime/pkg/cli/builder"
	"github.com/acorn-io/runtime/pkg/client"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func NewKubectl(c CommandContext) *cobra.Command {
	cmd := cli.Command(&Kube{client: c.ClientFactory}, cobra.Command{
		Use:          "kube [flags] COMMAND",
		Hidden:       true,
		SilenceUsage: true,
		Short:        "Run command with KUBECONFIG env set to a generated kubeconfig of the current project",
		Example: `
  # Write the kubeconfig for the current project that would be used to the provided file:
  acorn kube -w <filepath>

  # Run 'k9s' in the Account API Server for the 'acorn' project:
  acorn -j acorn kube k9s

  # Access the cluster of the 'aws-us-east-2' region for the current project:
  acorn -j acorn kube --region aws-us-east-2 $SHELL
`})
	cmd.Flags().SetInterspersed(false)
	return cmd
}

type Kube struct {
	client    ClientFactory
	Region    string `usage:"Get access to the cluster supporting the defined region"`
	WriteFile string `usage:"Write kubeconfig to file" short:"w"`
}

func (s *Kube) Run(cmd *cobra.Command, args []string) error {
	c, err := s.client.CreateDefault()
	if err != nil {
		return err
	}

	if s.WriteFile != "" {
		data, err := c.KubeConfig(cmd.Context(), &client.KubeProxyAddressOptions{
			Region: s.Region,
		})
		if err != nil {
			return err
		}
		return os.WriteFile(s.WriteFile, data, 0644)
	}

	if err = cobra.MinimumNArgs(1)(cmd, args); err != nil {
		return errors.Wrap(err, "command required when not using '-w' flag")
	}

	ctx, cancel := context.WithCancel(cmd.Context())
	defer cancel()

	server, err := c.KubeProxyAddress(ctx, &client.KubeProxyAddressOptions{
		Region: s.Region,
	})
	if err != nil {
		return err
	}

	f, err := os.CreateTemp("", "acorn-kube")
	if err != nil {
		return err
	}
	defer func() {
		_ = os.Remove(f.Name())
	}()

	_, err = f.Write([]byte(fmt.Sprintf(`
apiVersion: v1
clusters:
- cluster:
    server: "%s"
  name: default
contexts:
- context:
    cluster: default
    user: default
  name: default
current-context: default
kind: Config
preferences: {}
users:
- name: default
`, server)))
	if err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}

	k := exec.Command(args[0], args[1:]...)
	k.Env = append(os.Environ(), fmt.Sprintf("KUBECONFIG=%s", f.Name()), fmt.Sprintf("ACORN_KUBECONFIG=%s", f.Name()))
	k.Stdin = os.Stdin
	k.Stdout = os.Stdout
	k.Stderr = os.Stderr
	return k.Run()
}
