package cmd

import (
	{{ if .CLI.HasAnyRun }}
	"fmt"
	"os"
	{{end}}

  "github.com/spf13/cobra"
  {{ if or .CLI.Flags .CLI.Pflags }}
  // "github.com/spf13/viper"
  {{ end }}

  {{ if .CLI.Imports }}
	{{ range $i, $I := .CLI.Imports }}
	{{ $I.As }} "{{ $I.Path }}"
	{{ end }}
	{{ end }}

	{{ if .CLI.HasAnyRun }}
	"{{ .CLI.Package }}/lib/cmd"
	{{ end }}
)

{{ if .CLI.Long }}
var {{ .CLI.Name }}Long = `{{ .CLI.Long }}`
{{ end }}

{{ template "flag-vars" .CLI }}
{{ template "flag-init" .CLI }}

var RootCmd = &cobra.Command{

  {{ if .CLI.Usage}}
  Use: "{{ .CLI.Usage }}",
  {{ else }}
  Use: "{{ .CLI.Name }}",
  {{ end }}

  {{ if .CLI.Short}}
  Short: "{{ .CLI.Short }}",
  {{ end }}

  {{ if .CLI.Long }}
  Long: {{ .CLI.Name }}Long,
  {{ end }}

  {{ if .CLI.PersistentPrerun }}
  PersistentPreRun: func(cmd *cobra.Command, args []string) {
		var err error
    {{ template "args-parse" .CLI.Args }}

		err = libcmd.RootPersistentPreRun({{ template "lib-call.go" .CLI.Args }})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
  },
  {{ end }}

  {{ if .CLI.Prerun }}
  PreRun: func(cmd *cobra.Command, args []string) {
		var err error
    {{ template "args-parse" .CLI.Args }}

		err = libcmd.RootPreRun({{ template "lib-call.go" .CLI.Args }})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
  },
  {{ end }}

  {{ if not .CLI.OmitRun}}
  Run: func(cmd *cobra.Command, args []string) {
		var err error
    {{ template "args-parse" .CLI.Args }}

		err = libcmd.RootRun({{ template "lib-call.go" .CLI.Args }})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
  },
  {{ end }}

  {{ if .CLI.PersistentPostrun}}
  PersistentPostRun: func(cmd *cobra.Command, args []string) {
		var err error
    {{ template "args-parse" .CLI.Args }}

		err = libcmd.RootRPersistentPostun({{ template "lib-call.go" .CLI.Args }})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
  },
  {{ end }}

  {{ if .CLI.Postrun}}
  PostRun: func(cmd *cobra.Command, args []string) {
		var err error
    {{ template "args-parse" .CLI.Args }}

		err = libcmd.RootPostRun({{ template "lib-call.go" .CLI.Args }})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
  },
  {{ end }}
}

{{if .CLI.Commands}}
func init() {
	cobra.OnInitialize(initConfig)

	{{- range $i, $C := .CLI.Commands }}
	RootCmd.AddCommand({{ $C.CmdName }}Cmd)
	{{- end }}
}
{{ end }}

func initConfig() {

}