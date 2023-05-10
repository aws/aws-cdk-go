package awsdetective

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnGraph`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGraphProps := &CfnGraphProps{
//   	AutoEnableMembers: jsii.Boolean(false),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnGraphProps struct {
	// `AWS::Detective::Graph.AutoEnableMembers`.
	AutoEnableMembers interface{} `field:"optional" json:"autoEnableMembers" yaml:"autoEnableMembers"`
	// The tag values to assign to the new behavior graph.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

