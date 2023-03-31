package awssagemaker


// Specifies the endpoint capacity to activate for production.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   capacitySizeProperty := &capacitySizeProperty{
//   	type: jsii.String("type"),
//   	value: jsii.Number(123),
//   }
//
type CfnEndpoint_CapacitySizeProperty struct {
	// Specifies the endpoint capacity type.
	//
	// - `INSTANCE_COUNT` : The endpoint activates based on the number of instances.
	// - `CAPACITY_PERCENT` : The endpoint activates based on the specified percentage of capacity.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Defines the capacity size, either as a number of instances or a capacity percentage.
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

