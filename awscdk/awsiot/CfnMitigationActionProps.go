package awsiot

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnMitigationAction`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMitigationActionProps := &CfnMitigationActionProps{
//   	ActionParams: &ActionParamsProperty{
//   		AddThingsToThingGroupParams: &AddThingsToThingGroupParamsProperty{
//   			ThingGroupNames: []*string{
//   				jsii.String("thingGroupNames"),
//   			},
//
//   			// the properties below are optional
//   			OverrideDynamicGroups: jsii.Boolean(false),
//   		},
//   		EnableIoTLoggingParams: &EnableIoTLoggingParamsProperty{
//   			LogLevel: jsii.String("logLevel"),
//   			RoleArnForLogging: jsii.String("roleArnForLogging"),
//   		},
//   		PublishFindingToSnsParams: &PublishFindingToSnsParamsProperty{
//   			TopicArn: jsii.String("topicArn"),
//   		},
//   		ReplaceDefaultPolicyVersionParams: &ReplaceDefaultPolicyVersionParamsProperty{
//   			TemplateName: jsii.String("templateName"),
//   		},
//   		UpdateCaCertificateParams: &UpdateCACertificateParamsProperty{
//   			Action: jsii.String("action"),
//   		},
//   		UpdateDeviceCertificateParams: &UpdateDeviceCertificateParamsProperty{
//   			Action: jsii.String("action"),
//   		},
//   	},
//   	RoleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	ActionName: jsii.String("actionName"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnMitigationActionProps struct {
	// The set of parameters for this mitigation action.
	//
	// The parameters vary, depending on the kind of action you apply.
	ActionParams interface{} `field:"required" json:"actionParams" yaml:"actionParams"`
	// The IAM role ARN used to apply this mitigation action.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The friendly name of the mitigation action.
	ActionName *string `field:"optional" json:"actionName" yaml:"actionName"`
	// Metadata that can be used to manage the mitigation action.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

