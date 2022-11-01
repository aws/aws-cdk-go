package awswisdom

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnKnowledgeBase`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnKnowledgeBaseProps := &cfnKnowledgeBaseProps{
//   	knowledgeBaseType: jsii.String("knowledgeBaseType"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	renderingConfiguration: &renderingConfigurationProperty{
//   		templateUri: jsii.String("templateUri"),
//   	},
//   	serverSideEncryptionConfiguration: &serverSideEncryptionConfigurationProperty{
//   		kmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	sourceConfiguration: &sourceConfigurationProperty{
//   		appIntegrations: &appIntegrationsConfigurationProperty{
//   			appIntegrationArn: jsii.String("appIntegrationArn"),
//   			objectFields: []*string{
//   				jsii.String("objectFields"),
//   			},
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnKnowledgeBaseProps struct {
	// The type of knowledge base.
	//
	// Only CUSTOM knowledge bases allow you to upload your own content. EXTERNAL knowledge bases support integrations with third-party systems whose content is synchronized automatically.
	KnowledgeBaseType *string `field:"required" json:"knowledgeBaseType" yaml:"knowledgeBaseType"`
	// The name of the knowledge base.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The description.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Information about how to render the content.
	RenderingConfiguration interface{} `field:"optional" json:"renderingConfiguration" yaml:"renderingConfiguration"`
	// The KMS key used for encryption.
	ServerSideEncryptionConfiguration interface{} `field:"optional" json:"serverSideEncryptionConfiguration" yaml:"serverSideEncryptionConfiguration"`
	// The source of the knowledge base content.
	//
	// Only set this argument for EXTERNAL knowledge bases.
	SourceConfiguration interface{} `field:"optional" json:"sourceConfiguration" yaml:"sourceConfiguration"`
	// The tags used to organize, track, or control access for this resource.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

