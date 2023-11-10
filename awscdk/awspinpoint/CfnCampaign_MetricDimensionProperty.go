package awspinpoint


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricDimensionProperty := &MetricDimensionProperty{
//   	ComparisonOperator: jsii.String("comparisonOperator"),
//   	Value: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-metricdimension.html
//
type CfnCampaign_MetricDimensionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-metricdimension.html#cfn-pinpoint-campaign-metricdimension-comparisonoperator
	//
	ComparisonOperator *string `field:"optional" json:"comparisonOperator" yaml:"comparisonOperator"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-metricdimension.html#cfn-pinpoint-campaign-metricdimension-value
	//
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

