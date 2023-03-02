package awsinspector


// Properties for defining a `CfnResourceGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResourceGroupProps := &cfnResourceGroupProps{
//   	resourceGroupTags: []interface{}{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnResourceGroupProps struct {
	// The tags (key and value pairs) that will be associated with the resource group.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	ResourceGroupTags interface{} `field:"required" json:"resourceGroupTags" yaml:"resourceGroupTags"`
}

