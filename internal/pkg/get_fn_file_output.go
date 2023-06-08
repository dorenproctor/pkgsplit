package pkg

import (
	"fmt"
)

func (fn *Fn) GetFnFileOutput() string {
	return fmt.Sprintf(`package %s

%s`, fn.pkg.Name, fn.Body)
}

// func (fn *Fn) GetImportStr() string {
// 	imports := ""
// 	if len(fn.Imports) > 0 {
// 		importStrs := []string{}
// 		for _, i := range fn.Imports {
// 			s := ""
// 			importStrs = append(importStrs, s)
// 		}
// 		imports = fmt.Sprintf(
// 			"import (\n%s\n)",
// 			strings.Join(importStrs, "\n"),
// 		)
// 	}
// 	return imports
// }
