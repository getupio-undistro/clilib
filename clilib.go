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

package clilib

import (
	"context"
	"fmt"
	"io"

	"sigs.k8s.io/cluster-api/test/framework/exec"
)

const baseCommand = "undistro"

const (
	Create  = "create"
	Install = "install"
	Upgrade = "upgrade"
	Get     = "get"
	Move    = "move"
	Logs    = "logs"
	Rollout = "rollout"
	Delete  = "delete"
	Apply  = "apply"
)

type CLI struct {
	Writer io.Writer
}

func NewCLI(writer io.Writer) CLI {
	return CLI{
		Writer: writer,
	}
}

// UndistroExec executes an Undistro command with the parameterized argumments.
func (c CLI) UndistroExec(cmdName string, args ...string) (stdout, stderr string, err error) {
	cmd := exec.NewCommand(
		exec.WithCommand(baseCommand),
		exec.WithArgs(cmdName),
		exec.WithArgs(args...),
	)
	_, err = fmt.Fprintf(c.Writer, "Running command: %s\n", cmd.Cmd)
	if err != nil {
		return "", err.Error(), err
	}

	outByt, errByt, err := cmd.Run(context.Background())
	if err != nil {
		_, err = fmt.Fprintf(c.Writer, "Error: %s\n", err.Error())
		if err != nil {
			return "", err.Error(), err
		}
	}

	stdout = string(outByt)
	stderr = string(errByt)
	return
}

// Create executes "undistro create <args>".
func (c *CLI) Create(args ...string) (string, string, error) {
	return c.UndistroExec(Create, args...)
}

// Install executes "undistro install <args>".
func (c *CLI) Install(args ...string) (string, string, error) {
	return c.UndistroExec(Install, args...)
}

// Delete executes "undistro delete <args>".
func (c *CLI) Delete(args ...string) (string, string, error) {
	return c.UndistroExec(Delete, args...)
}

// Move executes "undistro move <args>".
func (c *CLI) Move(args ...string) (string, string, error) {
	return c.UndistroExec(Move, args...)
}

// Upgrade executes "undistro upgrade <args>".
func (c *CLI) Upgrade(args ...string) (string, string, error) {
	return c.UndistroExec(Upgrade, args...)
}

// Rollout executes "undistro rollout <args>".
func (c *CLI) Rollout(args ...string) (string, string, error) {
	return c.UndistroExec(Rollout, args...)
}

// Logs executes "undistro logs <args>".
func (c *CLI) Logs(args ...string) (string, string, error) {
	return c.UndistroExec(Logs, args...)
}

// Apply executes "undistro apply <args>".
func (c *CLI) Apply(args ...string) (string, string, error) {
	return c.UndistroExec(Apply, args...)
}
