package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   percentVisibleRangeProperty := &PercentVisibleRangeProperty{
//   	From: jsii.Number(123),
//   	To: jsii.Number(123),
//   }
//
type CfnTemplate_PercentVisibleRangeProperty struct {
	// `CfnTemplate.PercentVisibleRangeProperty.From`.
	From *float64 `field:"optional" json:"from" yaml:"from"`
	// `CfnTemplate.PercentVisibleRangeProperty.To`.
	To *float64 `field:"optional" json:"to" yaml:"to"`
}

