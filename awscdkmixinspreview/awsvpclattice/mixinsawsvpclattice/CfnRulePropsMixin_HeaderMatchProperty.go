package mixinsawsvpclattice


// Describes the constraints for a header match.
//
// Matches incoming requests with rule based on request header value before applying rule action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   headerMatchProperty := &HeaderMatchProperty{
//   	CaseSensitive: jsii.Boolean(false),
//   	Match: &HeaderMatchTypeProperty{
//   		Contains: jsii.String("contains"),
//   		Exact: jsii.String("exact"),
//   		Prefix: jsii.String("prefix"),
//   	},
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-rule-headermatch.html
//
type CfnRulePropsMixin_HeaderMatchProperty struct {
	// Indicates whether the match is case sensitive.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-rule-headermatch.html#cfn-vpclattice-rule-headermatch-casesensitive
	//
	// Default: - false.
	//
	CaseSensitive interface{} `field:"optional" json:"caseSensitive" yaml:"caseSensitive"`
	// The header match type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-rule-headermatch.html#cfn-vpclattice-rule-headermatch-match
	//
	Match interface{} `field:"optional" json:"match" yaml:"match"`
	// The name of the header.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-rule-headermatch.html#cfn-vpclattice-rule-headermatch-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

