package awsorganizations

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnPolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPolicyProps := &cfnPolicyProps{
//   	content: jsii.String("content"),
//   	name: jsii.String("name"),
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	targetIds: []*string{
//   		jsii.String("targetIds"),
//   	},
//   }
//
type CfnPolicyProps struct {
	// `AWS::Organizations::Policy.Content`.
	Content *string `field:"required" json:"content" yaml:"content"`
	// `AWS::Organizations::Policy.Name`.
	Name *string `field:"required" json:"name" yaml:"name"`
	// `AWS::Organizations::Policy.Type`.
	Type *string `field:"required" json:"type" yaml:"type"`
	// `AWS::Organizations::Policy.Description`.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// `AWS::Organizations::Policy.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// `AWS::Organizations::Policy.TargetIds`.
	TargetIds *[]*string `field:"optional" json:"targetIds" yaml:"targetIds"`
}

