package operations

import (
	"path"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/stackrox/rox/tools/generate-helpers/blevebindings/packagenames"
)

func renderSearchFunctionSignature(statement *jen.Statement) *jen.Statement {
	functionName := "Search"
	return statement.Id(functionName).Params(jen.Id("q").Op("*").Qual(packagenames.V1, "Query")).Parens(jen.List(jen.Index().Qual(packagenames.RoxSearch, "Result"), jen.Error()))
}

func generateSearch(props GeneratorProperties) (jen.Code, jen.Code) {
	interfaceMethod := renderSearchFunctionSignature(&jen.Statement{})

	objectName := props.Object
	if props.ObjectPathName != "" {
		objectName = props.ObjectPathName
	}
	mappingPath := path.Join(packagenames.RoxCentral, strings.ToLower(objectName), props.OptionsPath)
	implementation := renderSearchFunctionSignature(renderFuncBStarIndexer()).Block(
		metricLine("Search", props.Object),
		jen.Return(jen.Qual(packagenames.RoxBleve, "RunSearchRequest").Call(jen.Qual(packagenames.V1, props.SearchCategory), jen.Id("q"), jen.Id("b").Dot("index"), jen.Qual(mappingPath, "OptionsMap"))),
	)

	return interfaceMethod, implementation
}

func init() {
	supportedMethods["search"] = generateSearch
}
