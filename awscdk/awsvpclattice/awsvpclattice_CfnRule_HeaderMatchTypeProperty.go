package awsvpclattice


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   headerMatchTypeProperty := &headerMatchTypeProperty{
//   	contains: jsii.String("contains"),
//   	exact: jsii.String("exact"),
//   	prefix: jsii.String("prefix"),
//   }
//
type CfnRule_HeaderMatchTypeProperty struct {
	// `CfnRule.HeaderMatchTypeProperty.Contains`.
	Contains *string `field:"optional" json:"contains" yaml:"contains"`
	// `CfnRule.HeaderMatchTypeProperty.Exact`.
	Exact *string `field:"optional" json:"exact" yaml:"exact"`
	// `CfnRule.HeaderMatchTypeProperty.Prefix`.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

