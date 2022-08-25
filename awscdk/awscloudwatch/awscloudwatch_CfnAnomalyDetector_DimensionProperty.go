package awscloudwatch


// A dimension is a name/value pair that is part of the identity of a metric.
//
// Because dimensions are part of the unique identifier for a metric, whenever you add a unique name/value pair to one of your metrics, you are creating a new variation of that metric. For example, many Amazon EC2 metrics publish `InstanceId` as a dimension name, and the actual instance ID as the value for that dimension.
//
// You can assign up to 10 dimensions to a metric.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dimensionProperty := &dimensionProperty{
//   	name: jsii.String("name"),
//   	value: jsii.String("value"),
//   }
//
type CfnAnomalyDetector_DimensionProperty struct {
	// The name of the dimension.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The value of the dimension.
	//
	// Dimension values must contain only ASCII characters and must include at least one non-whitespace character.
	Value *string `field:"required" json:"value" yaml:"value"`
}

