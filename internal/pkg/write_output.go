package pkg

import (
	"fmt"
	"io/ioutil"
	"os"
)

func (p *Package) WriteOutput() error {
	outputDir := "output/" + p.Name
	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		return err
	}

	for _, fn := range p.Fns {
		filePath := fmt.Sprintf("%s/%s.go", outputDir, fn.Name)
		// s := ""package " + p.Name"
		s := fileOutput(p.Name, fn.Body)
		err := ioutil.WriteFile(filePath, []byte(s), 0644)
		if err != nil {
			return err
		}
	}

	return nil
}

func fileOutput(packageName, fnBody string) string {
	return fmt.Sprintf(`
package %s

%s
`, packageName, fnBody)
}
