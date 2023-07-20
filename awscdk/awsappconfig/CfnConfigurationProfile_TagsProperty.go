package awsappconfig


// Metadata to assign to the configuration profile.
//
// Tags help organize and categorize your AWS AppConfig resources. Each tag consists of a key and an optional value, both of which you define.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tagsProperty := &TagsProperty{
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-configurationprofile-tags.html
//
type CfnConfigurationProfile_TagsProperty struct {
	// The key-value string map.
	//
	// The valid character set is `[a-zA-Z+-=._:/]` . The tag key can be up to 128 characters and must not start with `aws:` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-configurationprofile-tags.html#cfn-appconfig-configurationprofile-tags-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The tag value can be up to 256 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-configurationprofile-tags.html#cfn-appconfig-configurationprofile-tags-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

