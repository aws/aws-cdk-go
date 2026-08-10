package awssagemaker


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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-experimenttrialcomponent-tagsitems.html
//
type CfnExperimentTrialComponent_TagsItemsProperty struct {
	// The tag key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-experimenttrialcomponent-tagsitems.html#cfn-sagemaker-experimenttrialcomponent-tagsitems-key
	//
	Key *string `field:"required" json:"key" yaml:"key"`
	// The tag value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-experimenttrialcomponent-tagsitems.html#cfn-sagemaker-experimenttrialcomponent-tagsitems-value
	//
	Value *string `field:"required" json:"value" yaml:"value"`
}

