package awsvpclattice


// Describes a path match type.
//
// Each rule can include only one of the following types of paths.
//
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
	// An exact match of the path.
	Exact *string `field:"optional" json:"exact" yaml:"exact"`
	// A prefix match of the path.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

