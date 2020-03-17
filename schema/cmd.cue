package schema

import (
  "strings"
)

Command : Common & {
  Name:     string
  cmdName:  strings.ToCamel(Name)
  CmdName:  strings.ToTitle(Name)

  Hidden:        bool | *false
  Aliases?:      [...string]
  PackageName?:  string
}

