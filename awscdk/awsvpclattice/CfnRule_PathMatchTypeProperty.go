package awsvpclattice


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pathMatchTypeProperty := &PathMatchTypeProperty{
//   	Exact: jsii.String("exact"),
//   	Prefix: jsii.String("prefix"),
//   }
//
type CfnRule_PathMatchTypeProperty struct {
	// `CfnRule.PathMatchTypeProperty.Exact`.
	Exact *string `field:"optional" json:"exact" yaml:"exact"`
	// `CfnRule.PathMatchTypeProperty.Prefix`.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

