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
type CfnAnalysis_PercentVisibleRangeProperty struct {
	// `CfnAnalysis.PercentVisibleRangeProperty.From`.
	From *float64 `field:"optional" json:"from" yaml:"from"`
	// `CfnAnalysis.PercentVisibleRangeProperty.To`.
	To *float64 `field:"optional" json:"to" yaml:"to"`
}

