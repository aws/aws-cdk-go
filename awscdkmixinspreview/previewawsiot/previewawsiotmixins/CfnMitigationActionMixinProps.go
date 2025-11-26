package previewawsiotmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnMitigationActionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnMitigationActionMixinProps := &CfnMitigationActionMixinProps{
//   	ActionName: jsii.String("actionName"),
//   	ActionParams: &ActionParamsProperty{
//   		AddThingsToThingGroupParams: &AddThingsToThingGroupParamsProperty{
//   			OverrideDynamicGroups: jsii.Boolean(false),
//   			ThingGroupNames: []*string{
//   				jsii.String("thingGroupNames"),
//   			},
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
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-mitigationaction.html
//
type CfnMitigationActionMixinProps struct {
	// The friendly name of the mitigation action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-mitigationaction.html#cfn-iot-mitigationaction-actionname
	//
	ActionName *string `field:"optional" json:"actionName" yaml:"actionName"`
	// The set of parameters for this mitigation action.
	//
	// The parameters vary, depending on the kind of action you apply.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-mitigationaction.html#cfn-iot-mitigationaction-actionparams
	//
	ActionParams interface{} `field:"optional" json:"actionParams" yaml:"actionParams"`
	// The IAM role ARN used to apply this mitigation action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-mitigationaction.html#cfn-iot-mitigationaction-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Metadata that can be used to manage the mitigation action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-mitigationaction.html#cfn-iot-mitigationaction-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

