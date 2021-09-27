package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

type buildFlags struct {
	Version string
	Name    string
	Pkg     string
	OutDir  string
	OS      string
	Arch    string
	Lambda  bool
}

var (
	defaultOutDir = "bin"
)

func main() {
	flags := buildFlags{}
	flag.StringVar(&flags.Version, "version", "", "Build version tag")
	flag.BoolVar(&flags.Lambda, "lambda", false, "Build Lambda zip file")
	flag.StringVar(&flags.Name, "target", "", "Build target (binary name)")
	flag.StringVar(&flags.Pkg, "pkg", "", "Build package")
	flag.StringVar(&flags.OutDir, "output", defaultOutDir, "Build output dir")
	flag.StringVar(&flags.OS, "os", "linux", "Build target OS (eg. darwin, linux, windows)")
	flag.StringVar(&flags.Arch, "arch", "amd64", "Build target Arch (eg. 386, amd64, arm, arm64)")
	flag.Parse()

	if flags.Lambda {
		flags.OS = "linux"
		flags.Arch = "amd64"
	}

	if err := build(&flags); err != nil {
		fmt.Printf("failed to run building process: %v\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}

func validate(flags *buildFlags) error {
	if flags.Name == "" {
		return fmt.Errorf("target is required")
	}

	if flags.OutDir == "" {
		return fmt.Errorf("output dir is required")
	}
	return nil
}

type Command struct {
	Env  []string
	Args []string
}

func newCommand() *Command {
	return &Command{}
}

func (c *Command) WithEnv(env []string) *Command {
	c.Env = env
	return c
}

func (c *Command) WithArgs(args []string) *Command {
	c.Args = args
	return c
}

func (c *Command) Run(name string) error {
	fmt.Println(name, c.Args)
	cmd := exec.Command(name, c.Args...)
	cmd.Env = append(os.Environ(), c.Env...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	if len(out) != 0 {
		fmt.Printf("combined out:\n%s\n", string(out))
	}
	return nil
}

func build(flags *buildFlags) error {
	// validate flags
	if err := validate(flags); err != nil {
		return fmt.Errorf("failed to pass validation: %v", err)
	}

	// prepare env
	env := []string{}
	env = append(env, "GOOS="+flags.OS)
	env = append(env, "GOARCH="+flags.Arch)

	// paepare output file
	outputFile := filepath.Join(flags.OutDir, flags.Name)
	if flags.OS == "windows" {
		outputFile += ".exe"
	}
	removeFiles(outputFile)

	// build go app
	args := []string{"build"}
	args = append(args, "-o", outputFile, flags.Pkg)
	if err := newCommand().WithArgs(args).WithEnv(env).Run("go"); err != nil {
		return fmt.Errorf("[GO_APP] cmd.Run() failed with %s", err)
	}

	// build lambda function
	if flags.Lambda {
		outputZip := outputFile + ".zip"
		args := []string{}
		args = append(args, "--output", outputZip, outputFile)
		if err := newCommand().WithArgs(args).WithEnv(env).Run("build-lambda-zip"); err != nil {
			return fmt.Errorf("[LAMBDA] cmd.Run() failed with %s", err)
		}
		removeFiles(outputFile)
	}

	return nil
}

func removeFiles(paths ...string) {
	for _, path := range paths {
		os.RemoveAll(path)
	}
}
