package awssecurityhub


// A number filter for querying findings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   numberFilterProperty := &NumberFilterProperty{
//   	Eq: jsii.Number(123),
//   	Gte: jsii.Number(123),
//   	Lte: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-numberfilter.html
//
type CfnInsight_NumberFilterProperty struct {
	// The equal-to condition to be applied to a single field when querying for findings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-numberfilter.html#cfn-securityhub-insight-numberfilter-eq
	//
	Eq *float64 `field:"optional" json:"eq" yaml:"eq"`
	// The greater-than-equal condition to be applied to a single field when querying for findings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-numberfilter.html#cfn-securityhub-insight-numberfilter-gte
	//
	Gte *float64 `field:"optional" json:"gte" yaml:"gte"`
	// The less-than-equal condition to be applied to a single field when querying for findings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-numberfilter.html#cfn-securityhub-insight-numberfilter-lte
	//
	Lte *float64 `field:"optional" json:"lte" yaml:"lte"`
}

