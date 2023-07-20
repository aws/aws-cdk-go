package awscloudwatch


// Dimension is an embedded property of the `AWS::CloudWatch::Alarm` type.
//
// Dimensions are name/value pairs that can be associated with a CloudWatch metric. You can specify a maximum of 10 dimensions for a given metric.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dimensionProperty := &DimensionProperty{
//   	Name: jsii.String("name"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alarm-dimension.html
//
type CfnAlarm_DimensionProperty struct {
	// The name of the dimension, from 1–255 characters in length.
	//
	// This dimension name must have been included when the metric was published.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alarm-dimension.html#cfn-cloudwatch-alarm-dimension-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The value for the dimension, from 1–255 characters in length.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alarm-dimension.html#cfn-cloudwatch-alarm-dimension-value
	//
	Value *string `field:"required" json:"value" yaml:"value"`
}

