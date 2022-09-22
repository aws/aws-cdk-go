package awsaps

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnRuleGroupsNamespace`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRuleGroupsNamespaceProps := &cfnRuleGroupsNamespaceProps{
//   	data: jsii.String("data"),
//   	name: jsii.String("name"),
//   	workspace: jsii.String("workspace"),
//
//   	// the properties below are optional
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnRuleGroupsNamespaceProps struct {
	// The rules definition file for this namespace.
	Data *string `field:"required" json:"data" yaml:"data"`
	// The name of the rule groups namespace.
	//
	// This property is required.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ARN of the workspace that contains this rule groups namespace.
	Workspace *string `field:"required" json:"workspace" yaml:"workspace"`
	// A list of key and value pairs for the workspace resources.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

