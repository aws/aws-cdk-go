package awsimagebuilder


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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-allworkflowbuildversions-tagsitems.html
//
type CfnAllWorkflowBuildVersionsPropsMixin_TagsItemsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-allworkflowbuildversions-tagsitems.html#cfn-imagebuilder-allworkflowbuildversions-tagsitems-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-allworkflowbuildversions-tagsitems.html#cfn-imagebuilder-allworkflowbuildversions-tagsitems-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

