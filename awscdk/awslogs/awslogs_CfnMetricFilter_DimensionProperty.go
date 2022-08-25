package awslogs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dimensionProperty := &dimensionProperty{
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//   }
//
type CfnMetricFilter_DimensionProperty struct {
	// `CfnMetricFilter.DimensionProperty.Key`.
	Key *string `field:"required" json:"key" yaml:"key"`
	// `CfnMetricFilter.DimensionProperty.Value`.
	Value *string `field:"required" json:"value" yaml:"value"`
}

