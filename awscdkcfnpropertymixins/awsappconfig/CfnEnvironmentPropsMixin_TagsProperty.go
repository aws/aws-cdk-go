package awsappconfig


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   tagsProperty := &TagsProperty{
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-environment-tags.html
//
type CfnEnvironmentPropsMixin_TagsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-environment-tags.html#cfn-appconfig-environment-tags-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-environment-tags.html#cfn-appconfig-environment-tags-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

