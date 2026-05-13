package awsdevopsagent

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnAgentSpace`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAgentSpaceProps := &CfnAgentSpaceProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   	Locale: jsii.String("locale"),
//   	OperatorApp: &OperatorAppProperty{
//   		Iam: &IamAuthConfigurationProperty{
//   			OperatorAppRoleArn: jsii.String("operatorAppRoleArn"),
//
//   			// the properties below are optional
//   			CreatedAt: jsii.String("createdAt"),
//   			UpdatedAt: jsii.String("updatedAt"),
//   		},
//   		Idc: &IdcAuthConfigurationProperty{
//   			IdcInstanceArn: jsii.String("idcInstanceArn"),
//   			OperatorAppRoleArn: jsii.String("operatorAppRoleArn"),
//
//   			// the properties below are optional
//   			CreatedAt: jsii.String("createdAt"),
//   			IdcApplicationArn: jsii.String("idcApplicationArn"),
//   			UpdatedAt: jsii.String("updatedAt"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-agentspace.html
//
type CfnAgentSpaceProps struct {
	// The name of the Agent Space.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-agentspace.html#cfn-devopsagent-agentspace-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The description of the Agent Space.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-agentspace.html#cfn-devopsagent-agentspace-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The ARN of the KMS key to use for encryption.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-agentspace.html#cfn-devopsagent-agentspace-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// The locale for the AgentSpace, which determines the language used in agent responses.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-agentspace.html#cfn-devopsagent-agentspace-locale
	//
	Locale *string `field:"optional" json:"locale" yaml:"locale"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-agentspace.html#cfn-devopsagent-agentspace-operatorapp
	//
	OperatorApp interface{} `field:"optional" json:"operatorApp" yaml:"operatorApp"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-agentspace.html#cfn-devopsagent-agentspace-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

