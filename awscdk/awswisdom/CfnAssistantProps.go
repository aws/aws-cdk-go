package awswisdom

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnAssistant`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAssistantProps := &CfnAssistantProps{
//   	Name: jsii.String("name"),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	ServerSideEncryptionConfiguration: &ServerSideEncryptionConfigurationProperty{
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnAssistantProps struct {
	// The name of the assistant.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of assistant.
	Type *string `field:"required" json:"type" yaml:"type"`
	// The description of the assistant.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The KMS key used for encryption.
	ServerSideEncryptionConfiguration interface{} `field:"optional" json:"serverSideEncryptionConfiguration" yaml:"serverSideEncryptionConfiguration"`
	// The tags used to organize, track, or control access for this resource.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

