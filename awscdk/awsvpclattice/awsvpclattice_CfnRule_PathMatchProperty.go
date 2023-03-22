package awsvpclattice


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pathMatchProperty := &pathMatchProperty{
//   	match: &pathMatchTypeProperty{
//   		exact: jsii.String("exact"),
//   		prefix: jsii.String("prefix"),
//   	},
//
//   	// the properties below are optional
//   	caseSensitive: jsii.Boolean(false),
//   }
//
type CfnRule_PathMatchProperty struct {
	// `CfnRule.PathMatchProperty.Match`.
	Match interface{} `field:"required" json:"match" yaml:"match"`
	// `CfnRule.PathMatchProperty.CaseSensitive`.
	CaseSensitive interface{} `field:"optional" json:"caseSensitive" yaml:"caseSensitive"`
}

