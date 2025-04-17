package awsbatch


// Information about a consumable resource required to run a job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   consumableResourceRequirementProperty := &ConsumableResourceRequirementProperty{
//   	ConsumableResource: jsii.String("consumableResource"),
//   	Quantity: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-consumableresourcerequirement.html
//
type CfnJobDefinition_ConsumableResourceRequirementProperty struct {
	// The name or ARN of the consumable resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-consumableresourcerequirement.html#cfn-batch-jobdefinition-consumableresourcerequirement-consumableresource
	//
	ConsumableResource *string `field:"required" json:"consumableResource" yaml:"consumableResource"`
	// The quantity of the consumable resource that is needed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-consumableresourcerequirement.html#cfn-batch-jobdefinition-consumableresourcerequirement-quantity
	//
	Quantity *float64 `field:"required" json:"quantity" yaml:"quantity"`
}

