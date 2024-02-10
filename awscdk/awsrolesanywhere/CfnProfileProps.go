package awsrolesanywhere

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnProfile`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnProfileProps := &CfnProfileProps{
//   	Name: jsii.String("name"),
//   	RoleArns: []*string{
//   		jsii.String("roleArns"),
//   	},
//
//   	// the properties below are optional
//   	DurationSeconds: jsii.Number(123),
//   	Enabled: jsii.Boolean(false),
//   	ManagedPolicyArns: []*string{
//   		jsii.String("managedPolicyArns"),
//   	},
//   	RequireInstanceProperties: jsii.Boolean(false),
//   	SessionPolicy: jsii.String("sessionPolicy"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rolesanywhere-profile.html
//
type CfnProfileProps struct {
	// The name of the profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rolesanywhere-profile.html#cfn-rolesanywhere-profile-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// A list of IAM role ARNs.
	//
	// During `CreateSession` , if a matching role ARN is provided, the properties in this profile will be applied to the intersection session policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rolesanywhere-profile.html#cfn-rolesanywhere-profile-rolearns
	//
	RoleArns *[]*string `field:"required" json:"roleArns" yaml:"roleArns"`
	// Sets the maximum number of seconds that vended temporary credentials through [CreateSession](https://docs.aws.amazon.com/rolesanywhere/latest/userguide/authentication-create-session.html) will be valid for, between 900 and 3600.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rolesanywhere-profile.html#cfn-rolesanywhere-profile-durationseconds
	//
	DurationSeconds *float64 `field:"optional" json:"durationSeconds" yaml:"durationSeconds"`
	// Indicates whether the profile is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rolesanywhere-profile.html#cfn-rolesanywhere-profile-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// A list of managed policy ARNs that apply to the vended session credentials.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rolesanywhere-profile.html#cfn-rolesanywhere-profile-managedpolicyarns
	//
	ManagedPolicyArns *[]*string `field:"optional" json:"managedPolicyArns" yaml:"managedPolicyArns"`
	// Specifies whether instance properties are required in temporary credential requests with this profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rolesanywhere-profile.html#cfn-rolesanywhere-profile-requireinstanceproperties
	//
	RequireInstanceProperties interface{} `field:"optional" json:"requireInstanceProperties" yaml:"requireInstanceProperties"`
	// A session policy that applies to the trust boundary of the vended session credentials.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rolesanywhere-profile.html#cfn-rolesanywhere-profile-sessionpolicy
	//
	SessionPolicy *string `field:"optional" json:"sessionPolicy" yaml:"sessionPolicy"`
	// The tags to attach to the profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rolesanywhere-profile.html#cfn-rolesanywhere-profile-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

