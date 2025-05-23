package awsiot


// The dimension of the metric.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricDimensionProperty := &MetricDimensionProperty{
//   	DimensionName: jsii.String("dimensionName"),
//
//   	// the properties below are optional
//   	Operator: jsii.String("operator"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-securityprofile-metricdimension.html
//
type CfnSecurityProfile_MetricDimensionProperty struct {
	// The name of the dimension.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-securityprofile-metricdimension.html#cfn-iot-securityprofile-metricdimension-dimensionname
	//
	DimensionName *string `field:"required" json:"dimensionName" yaml:"dimensionName"`
	// Operators are constructs that perform logical operations.
	//
	// Valid values are `IN` and `NOT_IN` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-securityprofile-metricdimension.html#cfn-iot-securityprofile-metricdimension-operator
	//
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
}

