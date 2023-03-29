package awsvpclattice


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   headerMatchProperty := &HeaderMatchProperty{
//   	Match: &HeaderMatchTypeProperty{
//   		Contains: jsii.String("contains"),
//   		Exact: jsii.String("exact"),
//   		Prefix: jsii.String("prefix"),
//   	},
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	CaseSensitive: jsii.Boolean(false),
//   }
//
type CfnRule_HeaderMatchProperty struct {
	// `CfnRule.HeaderMatchProperty.Match`.
	Match interface{} `field:"required" json:"match" yaml:"match"`
	// `CfnRule.HeaderMatchProperty.Name`.
	Name *string `field:"required" json:"name" yaml:"name"`
	// `CfnRule.HeaderMatchProperty.CaseSensitive`.
	CaseSensitive interface{} `field:"optional" json:"caseSensitive" yaml:"caseSensitive"`
}

