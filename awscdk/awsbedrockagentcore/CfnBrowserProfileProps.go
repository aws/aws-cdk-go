package awsbedrockagentcore


// Properties for defining a `CfnBrowserProfile`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnBrowserProfileProps := &CfnBrowserProfileProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-browserprofile.html
//
type CfnBrowserProfileProps struct {
	// The name of the browser profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-browserprofile.html#cfn-bedrockagentcore-browserprofile-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The description of the browser profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-browserprofile.html#cfn-bedrockagentcore-browserprofile-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A map of tag keys and values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-browserprofile.html#cfn-bedrockagentcore-browserprofile-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

