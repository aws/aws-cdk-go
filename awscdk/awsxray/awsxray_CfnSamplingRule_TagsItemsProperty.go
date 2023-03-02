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
type CfnSamplingRule_TagsItemsProperty struct {
	// `CfnSamplingRule.TagsItemsProperty.Key`.
	Key *string `field:"required" json:"key" yaml:"key"`
	// `CfnSamplingRule.TagsItemsProperty.Value`.
	Value *string `field:"required" json:"value" yaml:"value"`
}

