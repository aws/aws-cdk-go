package awscloudwatch

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnSpace`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSpaceProps := &CfnSpaceProps{
//   	DataAccessRoleArn: jsii.String("dataAccessRoleArn"),
//   	DomainId: jsii.String("domainId"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	AgentCoreEvaluationRoleArn: jsii.String("agentCoreEvaluationRoleArn"),
//   	EncryptionConfiguration: &EncryptionConfigurationProperty{
//   		EncryptionStrategy: jsii.String("encryptionStrategy"),
//
//   		// the properties below are optional
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
//   	},
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
type CfnSpaceProps struct {
	// The ARN of the IAM role used for data access.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-space.html#cfn-cloudwatch-space-dataaccessrolearn
	//
	DataAccessRoleArn *string `field:"required" json:"dataAccessRoleArn" yaml:"dataAccessRoleArn"`
	// The ID of the domain to create the space in.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-space.html#cfn-cloudwatch-space-domainid
	//
	DomainId *string `field:"required" json:"domainId" yaml:"domainId"`
	// A name that identifies the space.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-space.html#cfn-cloudwatch-space-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ARN of the IAM role used by AgentCore online evaluation.
	//
	// Absent when the space was created without one.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-space.html#cfn-cloudwatch-space-agentcoreevaluationrolearn
	//
	AgentCoreEvaluationRoleArn *string `field:"optional" json:"agentCoreEvaluationRoleArn" yaml:"agentCoreEvaluationRoleArn"`
	// The encryption configuration for the space.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-space.html#cfn-cloudwatch-space-encryptionconfiguration
	//
	EncryptionConfiguration interface{} `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-space.html#cfn-cloudwatch-space-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

