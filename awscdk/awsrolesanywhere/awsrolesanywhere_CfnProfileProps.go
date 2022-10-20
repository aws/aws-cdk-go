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
	// `AWS::RolesAnywhere::Profile.DurationSeconds`.
	DurationSeconds *float64 `field:"optional" json:"durationSeconds" yaml:"durationSeconds"`
	// `AWS::RolesAnywhere::Profile.Enabled`.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// `AWS::RolesAnywhere::Profile.ManagedPolicyArns`.
	ManagedPolicyArns *[]*string `field:"optional" json:"managedPolicyArns" yaml:"managedPolicyArns"`
	// `AWS::RolesAnywhere::Profile.Name`.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// `AWS::RolesAnywhere::Profile.RequireInstanceProperties`.
	RequireInstanceProperties interface{} `field:"optional" json:"requireInstanceProperties" yaml:"requireInstanceProperties"`
	// `AWS::RolesAnywhere::Profile.RoleArns`.
	RoleArns *[]*string `field:"optional" json:"roleArns" yaml:"roleArns"`
	// `AWS::RolesAnywhere::Profile.SessionPolicy`.
	SessionPolicy *string `field:"optional" json:"sessionPolicy" yaml:"sessionPolicy"`
	// `AWS::RolesAnywhere::Profile.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

