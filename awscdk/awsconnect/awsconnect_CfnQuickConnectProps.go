package awsconnect

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnQuickConnect`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnQuickConnectProps := &cfnQuickConnectProps{
//   	instanceArn: jsii.String("instanceArn"),
//   	name: jsii.String("name"),
//   	quickConnectConfig: &quickConnectConfigProperty{
//   		quickConnectType: jsii.String("quickConnectType"),
//
//   		// the properties below are optional
//   		phoneConfig: &phoneNumberQuickConnectConfigProperty{
//   			phoneNumber: jsii.String("phoneNumber"),
//   		},
//   		queueConfig: &queueQuickConnectConfigProperty{
//   			contactFlowArn: jsii.String("contactFlowArn"),
//   			queueArn: jsii.String("queueArn"),
//   		},
//   		userConfig: &userQuickConnectConfigProperty{
//   			contactFlowArn: jsii.String("contactFlowArn"),
//   			userArn: jsii.String("userArn"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnQuickConnectProps struct {
	// The Amazon Resource Name (ARN) of the instance.
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// The name of the quick connect.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Contains information about the quick connect.
	QuickConnectConfig interface{} `field:"required" json:"quickConnectConfig" yaml:"quickConnectConfig"`
	// The description of the quick connect.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The tags used to organize, track, or control access for this resource.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

