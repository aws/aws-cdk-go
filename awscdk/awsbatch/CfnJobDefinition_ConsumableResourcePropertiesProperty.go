package awsbatch


// Contains a list of consumable resources required by a job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   consumableResourcePropertiesProperty := &ConsumableResourcePropertiesProperty{
//   	ConsumableResourceList: []interface{}{
//   		&ConsumableResourceRequirementProperty{
//   			ConsumableResource: jsii.String("consumableResource"),
//   			Quantity: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-consumableresourceproperties.html
//
type CfnJobDefinition_ConsumableResourcePropertiesProperty struct {
	// The list of consumable resources required by a job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-consumableresourceproperties.html#cfn-batch-jobdefinition-consumableresourceproperties-consumableresourcelist
	//
	ConsumableResourceList interface{} `field:"required" json:"consumableResourceList" yaml:"consumableResourceList"`
}

