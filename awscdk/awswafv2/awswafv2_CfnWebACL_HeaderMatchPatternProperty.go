package awswafv2


// The filter to use to identify the subset of headers to inspect in a web request.
//
// You must specify exactly one setting: either `All` , `IncludedHeaders` , or `ExcludedHeaders` .
//
// Example JSON: `"MatchPattern": { "ExcludedHeaders": {"KeyToExclude1", "KeyToExclude2"} }`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//
//   headerMatchPatternProperty := &headerMatchPatternProperty{
//   	all: all,
//   	excludedHeaders: []*string{
//   		jsii.String("excludedHeaders"),
//   	},
//   	includedHeaders: []*string{
//   		jsii.String("includedHeaders"),
//   	},
//   }
//
type CfnWebACL_HeaderMatchPatternProperty struct {
	// Inspect all headers.
	All interface{} `field:"optional" json:"all" yaml:"all"`
	// Inspect only the headers whose keys don't match any of the strings specified here.
	ExcludedHeaders *[]*string `field:"optional" json:"excludedHeaders" yaml:"excludedHeaders"`
	// Inspect only the headers that have a key that matches one of the strings specified here.
	IncludedHeaders *[]*string `field:"optional" json:"includedHeaders" yaml:"includedHeaders"`
}

