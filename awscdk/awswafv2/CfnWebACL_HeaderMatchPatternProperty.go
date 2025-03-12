package awswafv2


// The filter to use to identify the subset of headers to inspect in a web request.
//
// You must specify exactly one setting: either `All` , `IncludedHeaders` , or `ExcludedHeaders` .
//
// Example JSON: `"MatchPattern": { "ExcludedHeaders": [ "KeyToExclude1", "KeyToExclude2" ] }`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//
//   headerMatchPatternProperty := &HeaderMatchPatternProperty{
//   	All: all,
//   	ExcludedHeaders: []*string{
//   		jsii.String("excludedHeaders"),
//   	},
//   	IncludedHeaders: []*string{
//   		jsii.String("includedHeaders"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-headermatchpattern.html
//
type CfnWebACL_HeaderMatchPatternProperty struct {
	// Inspect all headers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-headermatchpattern.html#cfn-wafv2-webacl-headermatchpattern-all
	//
	All interface{} `field:"optional" json:"all" yaml:"all"`
	// Inspect only the headers whose keys don't match any of the strings specified here.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-headermatchpattern.html#cfn-wafv2-webacl-headermatchpattern-excludedheaders
	//
	ExcludedHeaders *[]*string `field:"optional" json:"excludedHeaders" yaml:"excludedHeaders"`
	// Inspect only the headers that have a key that matches one of the strings specified here.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-headermatchpattern.html#cfn-wafv2-webacl-headermatchpattern-includedheaders
	//
	IncludedHeaders *[]*string `field:"optional" json:"includedHeaders" yaml:"includedHeaders"`
}

