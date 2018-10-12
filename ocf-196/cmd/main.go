package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/cbroglie/mustache"
	"github.com/openshift/library-go/pkg/git"
	s2igit "github.com/openshift/source-to-image/pkg/scm/git"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "build",
		Short: "Run openshift builds from knative",
	}
)

func runCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "run the build and show logs",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("run")
			return nil
		},
	}
	return cmd
}

type createOption struct {
	name        string
	image       string
	imageStream string
	toDocker    bool
}

func validateOpts(opt createOption) error {
	if opt.name == "" {
		return errors.New("name is empty")
	}
	if opt.image == "" {
		return errors.New("image is empty")
	}
	if opt.imageStream == "" {
		return errors.New("image-stream is empty")
	}
	return nil
}

func createCommand() *cobra.Command {
	var opt createOption
	cmd := &cobra.Command{
		Use:   "create",
		Short: "create or update the build config",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := validateOpts(opt); err != nil {
				return err
			}
			s := args[0]
			var source string
			var revision string
			url, err := s2igit.Parse(s)
			if err != nil {
				return err
			}
			gitRepo := git.NewRepository()
			if url.IsLocal() {
				remote, ok, err := gitRepo.GetOriginURL(s)
				if err != nil && err != git.ErrGitNotAvailable {
					return err
				}
				if !ok {
					return fmt.Errorf("source is not supported %s (git: %s, %s)", s, url, remote)
				}
				source = remote
				revision = gitRepo.GetRef(s)
			} else {
				source = s
				revision = "master"
			}
			template, err := ioutil.ReadFile("./spec.mustache")
			if err != nil {
				return err
			}
			m := map[string]string{
				"image":       opt.image,
				"imageStream": opt.imageStream,
				"name":        opt.name,
				"revision":    revision,
				"source":      source,
				"toDocker":    fmt.Sprintf("%v", opt.toDocker),
			}
			yaml, err := mustache.Render(string(template), m)
			if err != nil {
				return err
			}
			fmt.Println(yaml)
			return nil
		},
	}
	cmd.Flags().StringVar(&opt.name, "name", "", "build configuration name")
	cmd.Flags().StringVar(&opt.image, "image", "", "image name to push")
	cmd.Flags().StringVar(&opt.imageStream, "image-stream", "", "image stream to use as build input")
	cmd.Flags().BoolVar(&opt.toDocker, "toDocker", true, "push the image to a docker registry")
	return cmd
}

func main() {
	rootCmd.AddCommand(runCommand())
	rootCmd.AddCommand(createCommand())
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
