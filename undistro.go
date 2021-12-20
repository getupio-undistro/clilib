/*
Copyright 2020-2021 The UnDistro authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package undistro

import (
	"context"
	"fmt"
	"io"
	"sigs.k8s.io/cluster-api/test/framework/exec"
)

const baseCommand = "undistro"

type CmdName string

const (
	Create  = CmdName("create")
	Install = CmdName("install")
)

type CLI struct {
	Writer io.Writer
}

func NewCLI(writer io.Writer) CLI {
	return CLI{
		Writer: writer,
	}
}

func (c CLI) CreateCluster(clusterName, namespace, provider, flavor string) (stdout, stderr string) {
	inputs := []string{
		string(Create), "cluster", clusterName,
		"-n", namespace,
		"--infra", provider,
		"--flavor", flavor,
		"--ssh-key-name", "undistro",
		"--generate-file",
	}

	err := validateInputs(inputs)
	if err != nil {
		return
	}

	cmd := exec.NewCommand(
		exec.WithCommand(baseCommand),
		exec.WithArgs(inputs...),
		)
	_, err = fmt.Fprintf(c.Writer, "Running command: %s\n", cmd.Cmd)
	if err != nil {
		return
	}

	outByt, errByt, err := cmd.Run(context.Background())
	if err != nil {
		_, err = fmt.Fprintf(c.Writer, "Error: %s\n", err.Error())
		if err != nil {
			return
		}
	}

	stdout = string(outByt)
	stderr = string(errByt)
	return
}

func validateInputs([]string) error {
	return nil
}
