package awsvpclattice


// Describes a header match type.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-rule-headermatchtype.html
//
type CfnRule_HeaderMatchTypeProperty struct {
	// A contains type match.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-rule-headermatchtype.html#cfn-vpclattice-rule-headermatchtype-contains
	//
	Contains *string `field:"optional" json:"contains" yaml:"contains"`
	// An exact type match.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-rule-headermatchtype.html#cfn-vpclattice-rule-headermatchtype-exact
	//
	Exact *string `field:"optional" json:"exact" yaml:"exact"`
	// A prefix type match.
	//
	// Matches the value with the prefix.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-rule-headermatchtype.html#cfn-vpclattice-rule-headermatchtype-prefix
	//
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

