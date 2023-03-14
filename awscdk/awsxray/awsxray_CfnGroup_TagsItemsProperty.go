package awsxray


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tagsItemsProperty := &tagsItemsProperty{
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//   }
//
type CfnGroup_TagsItemsProperty struct {
	// `CfnGroup.TagsItemsProperty.Key`.
	Key *string `field:"required" json:"key" yaml:"key"`
	// `CfnGroup.TagsItemsProperty.Value`.
	Value *string `field:"required" json:"value" yaml:"value"`
}

