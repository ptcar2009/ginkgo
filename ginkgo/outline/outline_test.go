package outline

import (
	"encoding/json"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"path/filepath"

	. "github.com/ptcar2009/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("Validate outline from file with",
	func(srcFilename, jsonOutlineFilename, csvOutlineFilename string) {
		fset := token.NewFileSet()
		astFile, err := parser.ParseFile(fset, filepath.Join("_testdata", srcFilename), nil, 0)
		Expect(err).To(BeNil(), "error parsing source: %s", err)

		if err != nil {
			log.Fatalf("error parsing source: %s", err)
		}

		o, err := FromASTFile(astFile)
		Expect(err).To(BeNil(), "error creating outline: %s", err)

		gotJSON, err := json.MarshalIndent(o, "", "  ")
		Expect(err).To(BeNil(), "error marshalling outline to json: %s", err)

		wantJSON, err := ioutil.ReadFile(filepath.Join("_testdata", jsonOutlineFilename))
		Expect(err).To(BeNil(), "error reading JSON outline fixture: %s", err)

		Expect(gotJSON).To(MatchJSON(wantJSON))

		gotCSV := o.String()

		wantCSV, err := ioutil.ReadFile(filepath.Join("_testdata", csvOutlineFilename))
		Expect(err).To(BeNil(), "error reading CSV outline fixture: %s", err)

		Expect(gotCSV).To(Equal(string(wantCSV)))
	},
	// To add a test:
	// 1. Create the input, e.g., `myspecialcase_test.go`
	// 2. Create the sample CSV and JSON results: Run `bash ./_testdata/create_result.sh ./_testdata/myspecialcase_test.go`
	// 3. Add an Entry below, by copying an existing one, and substituting `myspecialcase` where needed.
	// To re-create the sample results for a test:
	// 1. Run `bash ./_testdata/create_result.sh ./testdata/myspecialcase_test.go`
	// To re-create the sample results for all tests:
	// 1. Run `for name in ./_testdata/*_test.go; do bash ./_testdata/create_result.sh $name; done`
	Entry("normal import of ginkgo package (no dot, no alias), normal container and specs", "nodot_test.go", "nodot_test.go.json", "nodot_test.go.csv"),
	Entry("aliased import of ginkgo package, normal container and specs", "alias_test.go", "alias_test.go.json", "alias_test.go.csv"),
	Entry("normal containers and specs", "normal_test.go", "normal_test.go.json", "normal_test.go.csv"),
	Entry("focused containers and specs", "focused_test.go", "focused_test.go.json", "focused_test.go.csv"),
	Entry("pending containers and specs", "pending_test.go", "pending_test.go.json", "pending_test.go.csv"),
	Entry("nested focused containers and specs", "nestedfocused_test.go", "nestedfocused_test.go.json", "nestedfocused_test.go.csv"),
	Entry("mixed focused containers and specs", "mixed_test.go", "mixed_test.go.json", "mixed_test.go.csv"),
	Entry("suite setup", "suite_test.go", "suite_test.go.json", "suite_test.go.csv"),
)
