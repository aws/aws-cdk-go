package awssagemaker


// Specifies the type and size of the endpoint capacity to activate for a blue/green deployment, a rolling deployment, or a rollback strategy.
//
// You can specify your batches as either instance count or the overall percentage or your fleet.
//
// For a rollback strategy, if you don't specify the fields in this object, or if you set the `Value` to 100%, then SageMaker uses a blue/green rollback strategy and rolls all traffic back to the blue fleet.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   capacitySizeProperty := &CapacitySizeProperty{
//   	Type: jsii.String("type"),
//   	Value: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpoint-capacitysize.html
//
type CfnEndpoint_CapacitySizeProperty struct {
	// Specifies the endpoint capacity type.
	//
	// - `INSTANCE_COUNT` : The endpoint activates based on the number of instances.
	// - `CAPACITY_PERCENT` : The endpoint activates based on the specified percentage of capacity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpoint-capacitysize.html#cfn-sagemaker-endpoint-capacitysize-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// Defines the capacity size, either as a number of instances or a capacity percentage.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpoint-capacitysize.html#cfn-sagemaker-endpoint-capacitysize-value
	//
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

