package awsquicksight


// The percent range in the visible range.
//
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
type CfnDashboard_PercentVisibleRangeProperty struct {
	// The lower bound of the range.
	From *float64 `field:"optional" json:"from" yaml:"from"`
	// The top bound of the range.
	To *float64 `field:"optional" json:"to" yaml:"to"`
}

