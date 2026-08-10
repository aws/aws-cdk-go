package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   tagsItemsProperty := &TagsItemsProperty{
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-context-tagsitems.html
//
type CfnContextPropsMixin_TagsItemsProperty struct {
	// The tag key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-context-tagsitems.html#cfn-sagemaker-context-tagsitems-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The tag value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-context-tagsitems.html#cfn-sagemaker-context-tagsitems-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

