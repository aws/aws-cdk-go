package awscloudwatch

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnSpacePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnSpaceMixinProps := &CfnSpaceMixinProps{
//   	AgentCoreEvaluationRoleArn: jsii.String("agentCoreEvaluationRoleArn"),
//   	DataAccessRoleArn: jsii.String("dataAccessRoleArn"),
//   	DomainId: jsii.String("domainId"),
//   	EncryptionConfiguration: &EncryptionConfigurationProperty{
//   		EncryptionStrategy: jsii.String("encryptionStrategy"),
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
//   	},
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-space.html
//
type CfnSpaceMixinProps struct {
	// The ARN of the IAM role used by AgentCore online evaluation.
	//
	// Absent when the space was created without one.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-space.html#cfn-cloudwatch-space-agentcoreevaluationrolearn
	//
	AgentCoreEvaluationRoleArn *string `field:"optional" json:"agentCoreEvaluationRoleArn" yaml:"agentCoreEvaluationRoleArn"`
	// The ARN of the IAM role used for data access.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-space.html#cfn-cloudwatch-space-dataaccessrolearn
	//
	DataAccessRoleArn *string `field:"optional" json:"dataAccessRoleArn" yaml:"dataAccessRoleArn"`
	// The ID of the domain to create the space in.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-space.html#cfn-cloudwatch-space-domainid
	//
	DomainId *string `field:"optional" json:"domainId" yaml:"domainId"`
	// The encryption configuration for the space.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-space.html#cfn-cloudwatch-space-encryptionconfiguration
	//
	EncryptionConfiguration interface{} `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// A name that identifies the space.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-space.html#cfn-cloudwatch-space-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-space.html#cfn-cloudwatch-space-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

