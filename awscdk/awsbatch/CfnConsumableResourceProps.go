package awsbatch


// Properties for defining a `CfnConsumableResource`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnConsumableResourceProps := &CfnConsumableResourceProps{
//   	ResourceType: jsii.String("resourceType"),
//   	TotalQuantity: jsii.Number(123),
//
//   	// the properties below are optional
//   	ConsumableResourceName: jsii.String("consumableResourceName"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-consumableresource.html
//
type CfnConsumableResourceProps struct {
	// Indicates whether the resource is available to be re-used after a job completes. Can be one of:.
	//
	// - `REPLENISHABLE`
	// - `NON_REPLENISHABLE`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-consumableresource.html#cfn-batch-consumableresource-resourcetype
	//
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
	// The total amount of the consumable resource that is available.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-consumableresource.html#cfn-batch-consumableresource-totalquantity
	//
	TotalQuantity *float64 `field:"required" json:"totalQuantity" yaml:"totalQuantity"`
	// The name of the consumable resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-consumableresource.html#cfn-batch-consumableresource-consumableresourcename
	//
	ConsumableResourceName *string `field:"optional" json:"consumableResourceName" yaml:"consumableResourceName"`
	// The tags that you apply to the consumable resource to help you categorize and organize your resources.
	//
	// Each tag consists of a key and an optional value. For more information, see [Tagging your AWS Batch resources](https://docs.aws.amazon.com/batch/latest/userguide/using-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-consumableresource.html#cfn-batch-consumableresource-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

