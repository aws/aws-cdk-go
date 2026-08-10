package awswellarchitected


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tagsItemsProperty := &TagsItemsProperty{
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wellarchitected-reviewtemplate-tagsitems.html
//
type CfnReviewTemplate_TagsItemsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wellarchitected-reviewtemplate-tagsitems.html#cfn-wellarchitected-reviewtemplate-tagsitems-key
	//
	Key *string `field:"required" json:"key" yaml:"key"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wellarchitected-reviewtemplate-tagsitems.html#cfn-wellarchitected-reviewtemplate-tagsitems-value
	//
	Value *string `field:"required" json:"value" yaml:"value"`
}

