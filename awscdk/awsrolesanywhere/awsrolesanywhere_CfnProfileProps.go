package awsrolesanywhere

import (
	"github.com/aws/aws-cdk-go/awscdk"
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
type CfnProfileProps struct {
	// The name of the profile.
	Name *string `field:"required" json:"name" yaml:"name"`
	// A list of IAM role ARNs.
	//
	// During `CreateSession` , if a matching role ARN is provided, the properties in this profile will be applied to the intersection session policy.
	RoleArns *[]*string `field:"required" json:"roleArns" yaml:"roleArns"`
	// Sets the maximum number of seconds that vended temporary credentials through [CreateSession](https://docs.aws.amazon.com/rolesanywhere/latest/userguide/authentication-create-session.html) will be valid for, between 900 and 3600.
	DurationSeconds *float64 `field:"optional" json:"durationSeconds" yaml:"durationSeconds"`
	// Indicates whether the profile is enabled.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// A list of managed policy ARNs that apply to the vended session credentials.
	ManagedPolicyArns *[]*string `field:"optional" json:"managedPolicyArns" yaml:"managedPolicyArns"`
	// Specifies whether instance properties are required in temporary credential requests with this profile.
	RequireInstanceProperties interface{} `field:"optional" json:"requireInstanceProperties" yaml:"requireInstanceProperties"`
	// A session policy that applies to the trust boundary of the vended session credentials.
	SessionPolicy *string `field:"optional" json:"sessionPolicy" yaml:"sessionPolicy"`
	// The tags to attach to the profile.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

