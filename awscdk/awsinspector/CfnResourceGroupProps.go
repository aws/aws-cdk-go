package awsinspector


// Properties for defining a `CfnResourceGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResourceGroupProps := &CfnResourceGroupProps{
//   	ResourceGroupTags: []interface{}{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspector-resourcegroup.html
//
type CfnResourceGroupProps struct {
	// The tags (key and value pairs) that will be associated with the resource group.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspector-resourcegroup.html#cfn-inspector-resourcegroup-resourcegrouptags
	//
	ResourceGroupTags interface{} `field:"required" json:"resourceGroupTags" yaml:"resourceGroupTags"`
}

