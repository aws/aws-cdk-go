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
	// The number of seconds vended session credentials will be valid for.
	DurationSeconds *float64 `field:"optional" json:"durationSeconds" yaml:"durationSeconds"`
	// The enabled status of the resource.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// A list of managed policy ARNs.
	//
	// Managed policies identified by this list will be applied to the vended session credentials.
	ManagedPolicyArns *[]*string `field:"optional" json:"managedPolicyArns" yaml:"managedPolicyArns"`
	// The customer specified name of the resource.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specifies whether instance properties are required in CreateSession requests with this profile.
	RequireInstanceProperties interface{} `field:"optional" json:"requireInstanceProperties" yaml:"requireInstanceProperties"`
	// A list of IAM role ARNs that can be assumed when this profile is specified in a CreateSession request.
	RoleArns *[]*string `field:"optional" json:"roleArns" yaml:"roleArns"`
	// A session policy that will applied to the trust boundary of the vended session credentials.
	SessionPolicy *string `field:"optional" json:"sessionPolicy" yaml:"sessionPolicy"`
	// A list of Tags.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

