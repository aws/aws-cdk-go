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
//   cfnProfileProps := &cfnProfileProps{
//   	durationSeconds: jsii.Number(123),
//   	enabled: jsii.Boolean(false),
//   	managedPolicyArns: []*string{
//   		jsii.String("managedPolicyArns"),
//   	},
//   	name: jsii.String("name"),
//   	requireInstanceProperties: jsii.Boolean(false),
//   	roleArns: []*string{
//   		jsii.String("roleArns"),
//   	},
//   	sessionPolicy: jsii.String("sessionPolicy"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnProfileProps struct {
	// Sets the maximum number of seconds that vended temporary credentials through [CreateSession](https://docs.aws.amazon.com/rolesanywhere/latest/userguide/authentication-create-session.html) will be valid for, between 900 and 3600.
	DurationSeconds *float64 `field:"optional" json:"durationSeconds" yaml:"durationSeconds"`
	// Indicates whether the profile is enabled.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// A list of managed policy ARNs that apply to the vended session credentials.
	ManagedPolicyArns *[]*string `field:"optional" json:"managedPolicyArns" yaml:"managedPolicyArns"`
	// The name of the profile.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specifies whether instance properties are required in temporary credential requests with this profile.
	RequireInstanceProperties interface{} `field:"optional" json:"requireInstanceProperties" yaml:"requireInstanceProperties"`
	// A list of IAM role ARNs.
	//
	// During `CreateSession` , if a matching role ARN is provided, the properties in this profile will be applied to the intersection session policy.
	RoleArns *[]*string `field:"optional" json:"roleArns" yaml:"roleArns"`
	// A session policy that applies to the trust boundary of the vended session credentials.
	SessionPolicy *string `field:"optional" json:"sessionPolicy" yaml:"sessionPolicy"`
	// A list of tags to attach to the profile.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

