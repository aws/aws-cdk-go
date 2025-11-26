package previewawsvpclatticemixins


// Describes the conditions that can be applied when matching a path for incoming requests.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   pathMatchProperty := &PathMatchProperty{
//   	CaseSensitive: jsii.Boolean(false),
//   	Match: &PathMatchTypeProperty{
//   		Exact: jsii.String("exact"),
//   		Prefix: jsii.String("prefix"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-rule-pathmatch.html
//
type CfnRulePropsMixin_PathMatchProperty struct {
	// Indicates whether the match is case sensitive.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-rule-pathmatch.html#cfn-vpclattice-rule-pathmatch-casesensitive
	//
	// Default: - false.
	//
	CaseSensitive interface{} `field:"optional" json:"caseSensitive" yaml:"caseSensitive"`
	// The type of path match.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-rule-pathmatch.html#cfn-vpclattice-rule-pathmatch-match
	//
	Match interface{} `field:"optional" json:"match" yaml:"match"`
}

