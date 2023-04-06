package awsvpclattice


// Describes a header match type.
//
// Only one can be provided.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   headerMatchTypeProperty := &HeaderMatchTypeProperty{
//   	Contains: jsii.String("contains"),
//   	Exact: jsii.String("exact"),
//   	Prefix: jsii.String("prefix"),
//   }
//
type CfnRule_HeaderMatchTypeProperty struct {
	// Specifies a contains type match.
	Contains *string `field:"optional" json:"contains" yaml:"contains"`
	// Specifies an exact type match.
	Exact *string `field:"optional" json:"exact" yaml:"exact"`
	// Specifies a prefix type match.
	//
	// Matches the value with the prefix.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

