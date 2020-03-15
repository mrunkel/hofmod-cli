package schema

import (
  "strings"
)

Command :: common & {
  Name:     string
  cmdName: strings.ToCamel(Name)
  CmdName: strings.ToTitle(Name)

  Aliases?: [...string]
  PackageName?: string
}
