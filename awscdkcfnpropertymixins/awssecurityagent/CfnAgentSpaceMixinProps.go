package awssecurityagent

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnAgentSpacePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnAgentSpaceMixinProps := &CfnAgentSpaceMixinProps{
//   	AwsResources: &AWSResourcesProperty{
//   		IamRoles: []*string{
//   			jsii.String("iamRoles"),
//   		},
//   		LambdaFunctionArns: []*string{
//   			jsii.String("lambdaFunctionArns"),
//   		},
//   		LogGroups: []*string{
//   			jsii.String("logGroups"),
//   		},
//   		S3Buckets: []*string{
//   			jsii.String("s3Buckets"),
//   		},
//   		SecretArns: []*string{
//   			jsii.String("secretArns"),
//   		},
//   		Vpcs: []interface{}{
//   			&VpcConfigProperty{
//   				SecurityGroupArns: []*string{
//   					jsii.String("securityGroupArns"),
//   				},
//   				SubnetArns: []*string{
//   					jsii.String("subnetArns"),
//   				},
//   				VpcArn: jsii.String("vpcArn"),
//   			},
//   		},
//   	},
//   	CodeReviewSettings: &CodeReviewSettingsProperty{
//   		ControlsScanning: jsii.Boolean(false),
//   		GeneralPurposeScanning: jsii.Boolean(false),
//   	},
//   	Description: jsii.String("description"),
//   	IntegratedResources: []interface{}{
//   		&IntegratedResourceProperty{
//   			Integration: jsii.String("integration"),
//   		},
//   	},
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TargetDomainIds: []*string{
//   		jsii.String("targetDomainIds"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityagent-agentspace.html
//
type CfnAgentSpaceMixinProps struct {
	// AWS resource configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityagent-agentspace.html#cfn-securityagent-agentspace-awsresources
	//
	AwsResources interface{} `field:"optional" json:"awsResources" yaml:"awsResources"`
	// Details of code review settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityagent-agentspace.html#cfn-securityagent-agentspace-codereviewsettings
	//
	CodeReviewSettings interface{} `field:"optional" json:"codeReviewSettings" yaml:"codeReviewSettings"`
	// Description of the agent space.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityagent-agentspace.html#cfn-securityagent-agentspace-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Integrated Resources configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityagent-agentspace.html#cfn-securityagent-agentspace-integratedresources
	//
	IntegratedResources interface{} `field:"optional" json:"integratedResources" yaml:"integratedResources"`
	// Identifier of the KMS key used to encrypt data.
	//
	// Can be a key ID, key ARN, alias name, or alias ARN. If not specified, an AWS managed key is used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityagent-agentspace.html#cfn-securityagent-agentspace-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Name of the agent space.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityagent-agentspace.html#cfn-securityagent-agentspace-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Tags for the agent space.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityagent-agentspace.html#cfn-securityagent-agentspace-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// List of target domain identifiers registered with the agent space.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityagent-agentspace.html#cfn-securityagent-agentspace-targetdomainids
	//
	TargetDomainIds *[]*string `field:"optional" json:"targetDomainIds" yaml:"targetDomainIds"`
}

