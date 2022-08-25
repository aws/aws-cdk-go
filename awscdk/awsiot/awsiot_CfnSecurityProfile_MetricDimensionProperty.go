package awsiot


// The dimension of the metric.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricDimensionProperty := &metricDimensionProperty{
//   	dimensionName: jsii.String("dimensionName"),
//
//   	// the properties below are optional
//   	operator: jsii.String("operator"),
//   }
//
type CfnSecurityProfile_MetricDimensionProperty struct {
	// The name of the dimension.
	DimensionName *string `field:"required" json:"dimensionName" yaml:"dimensionName"`
	// Operators are constructs that perform logical operations.
	//
	// Valid values are `IN` and `NOT_IN` .
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
}

