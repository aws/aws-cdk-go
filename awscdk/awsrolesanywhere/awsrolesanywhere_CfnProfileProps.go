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
	// The customer specified name of the resource.
	Name *string `field:"required" json:"name" yaml:"name"`
	// A list of IAM role ARNs that can be assumed when this profile is specified in a CreateSession request.
	RoleArns *[]*string `field:"required" json:"roleArns" yaml:"roleArns"`
	// The number of seconds vended session credentials will be valid for.
	DurationSeconds *float64 `field:"optional" json:"durationSeconds" yaml:"durationSeconds"`
	// The enabled status of the resource.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// A list of managed policy ARNs.
	//
	// Managed policies identified by this list will be applied to the vended session credentials.
	ManagedPolicyArns *[]*string `field:"optional" json:"managedPolicyArns" yaml:"managedPolicyArns"`
	// Specifies whether instance properties are required in CreateSession requests with this profile.
	RequireInstanceProperties interface{} `field:"optional" json:"requireInstanceProperties" yaml:"requireInstanceProperties"`
	// A session policy that will applied to the trust boundary of the vended session credentials.
	SessionPolicy *string `field:"optional" json:"sessionPolicy" yaml:"sessionPolicy"`
	// A list of Tags.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

