package awsmediapackage

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnChannel`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnChannelProps := &cfnChannelProps{
//   	id: jsii.String("id"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	egressAccessLogs: &logConfigurationProperty{
//   		logGroupName: jsii.String("logGroupName"),
//   	},
//   	ingressAccessLogs: &logConfigurationProperty{
//   		logGroupName: jsii.String("logGroupName"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnChannelProps struct {
	// Unique identifier that you assign to the channel.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Any descriptive information that you want to add to the channel for future identification purposes.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Configures egress access logs.
	EgressAccessLogs interface{} `field:"optional" json:"egressAccessLogs" yaml:"egressAccessLogs"`
	// Configures ingress access logs.
	IngressAccessLogs interface{} `field:"optional" json:"ingressAccessLogs" yaml:"ingressAccessLogs"`
	// The tags to assign to the channel.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

