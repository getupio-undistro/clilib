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

type ProviderName string

const (
	AWS = ProviderName("aws")
	OpenStack = ProviderName("openstack")
)

type CLI struct {
	Writer io.Writer
}

func NewCLI(writer io.Writer) CLI {
	return CLI{
		Writer: writer,
	}
}

func (c CLI) CreateCluster(clusterName, namespace, provider, flavor string, generateFile bool) (stdout, stderr string) {
	inputs := []string{
		string(Create), "cluster", clusterName,
		"-n", namespace,
		"--infra", provider,
		"--ssh-key-name", "undistro",
	}

	if generateFile {
		inputs = append(inputs, "--generate-file")
	}

	switch provider {
	case string(AWS):
		if flavor == "" {
			return "", "AWS provider requires a flavor"
		}
		inputs = append(inputs, "--flavor", flavor)
	case string(OpenStack):
	default:
		return "", "Invalid provider"
	}

	cmd := exec.NewCommand(
		exec.WithCommand(baseCommand),
		exec.WithArgs(inputs...),
		)
	_, err := fmt.Fprintf(c.Writer, "Running command: %s\n", cmd.Cmd)
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
