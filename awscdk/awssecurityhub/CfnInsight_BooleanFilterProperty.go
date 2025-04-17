package awssecurityhub


// Boolean filter for querying findings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   booleanFilterProperty := &BooleanFilterProperty{
//   	Value: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-booleanfilter.html
//
type CfnInsight_BooleanFilterProperty struct {
	// The value of the boolean.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-booleanfilter.html#cfn-securityhub-insight-booleanfilter-value
	//
	Value interface{} `field:"required" json:"value" yaml:"value"`
}

