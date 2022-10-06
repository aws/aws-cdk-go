package awsiot

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnMitigationAction`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMitigationActionProps := &cfnMitigationActionProps{
//   	actionParams: &actionParamsProperty{
//   		addThingsToThingGroupParams: &addThingsToThingGroupParamsProperty{
//   			thingGroupNames: []*string{
//   				jsii.String("thingGroupNames"),
//   			},
//
//   			// the properties below are optional
//   			overrideDynamicGroups: jsii.Boolean(false),
//   		},
//   		enableIoTLoggingParams: &enableIoTLoggingParamsProperty{
//   			logLevel: jsii.String("logLevel"),
//   			roleArnForLogging: jsii.String("roleArnForLogging"),
//   		},
//   		publishFindingToSnsParams: &publishFindingToSnsParamsProperty{
//   			topicArn: jsii.String("topicArn"),
//   		},
//   		replaceDefaultPolicyVersionParams: &replaceDefaultPolicyVersionParamsProperty{
//   			templateName: jsii.String("templateName"),
//   		},
//   		updateCaCertificateParams: &updateCACertificateParamsProperty{
//   			action: jsii.String("action"),
//   		},
//   		updateDeviceCertificateParams: &updateDeviceCertificateParamsProperty{
//   			action: jsii.String("action"),
//   		},
//   	},
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	actionName: jsii.String("actionName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
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

