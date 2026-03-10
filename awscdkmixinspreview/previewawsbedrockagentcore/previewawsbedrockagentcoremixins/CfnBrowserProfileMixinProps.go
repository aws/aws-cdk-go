package previewawsbedrockagentcoremixins


// Properties for CfnBrowserProfilePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnBrowserProfileMixinProps := &CfnBrowserProfileMixinProps{
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-browserprofile.html
//
type CfnBrowserProfileMixinProps struct {
	// The description of the browser profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-browserprofile.html#cfn-bedrockagentcore-browserprofile-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the browser profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-browserprofile.html#cfn-bedrockagentcore-browserprofile-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A map of tag keys and values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-browserprofile.html#cfn-bedrockagentcore-browserprofile-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

