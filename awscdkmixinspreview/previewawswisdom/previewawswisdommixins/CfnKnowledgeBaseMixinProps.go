package previewawswisdommixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnKnowledgeBasePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnKnowledgeBaseMixinProps := &CfnKnowledgeBaseMixinProps{
//   	Description: jsii.String("description"),
//   	KnowledgeBaseType: jsii.String("knowledgeBaseType"),
//   	Name: jsii.String("name"),
//   	RenderingConfiguration: &RenderingConfigurationProperty{
//   		TemplateUri: jsii.String("templateUri"),
//   	},
//   	ServerSideEncryptionConfiguration: &ServerSideEncryptionConfigurationProperty{
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	SourceConfiguration: &SourceConfigurationProperty{
//   		AppIntegrations: &AppIntegrationsConfigurationProperty{
//   			AppIntegrationArn: jsii.String("appIntegrationArn"),
//   			ObjectFields: []*string{
//   				jsii.String("objectFields"),
//   			},
//   		},
//   		ManagedSourceConfiguration: &ManagedSourceConfigurationProperty{
//   			WebCrawlerConfiguration: &WebCrawlerConfigurationProperty{
//   				CrawlerLimits: &CrawlerLimitsProperty{
//   					RateLimit: jsii.Number(123),
//   				},
//   				ExclusionFilters: []*string{
//   					jsii.String("exclusionFilters"),
//   				},
//   				InclusionFilters: []*string{
//   					jsii.String("inclusionFilters"),
//   				},
//   				Scope: jsii.String("scope"),
//   				UrlConfiguration: &UrlConfigurationProperty{
//   					SeedUrls: []interface{}{
//   						&SeedUrlProperty{
//   							Url: jsii.String("url"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VectorIngestionConfiguration: &VectorIngestionConfigurationProperty{
//   		ChunkingConfiguration: &ChunkingConfigurationProperty{
//   			ChunkingStrategy: jsii.String("chunkingStrategy"),
//   			FixedSizeChunkingConfiguration: &FixedSizeChunkingConfigurationProperty{
//   				MaxTokens: jsii.Number(123),
//   				OverlapPercentage: jsii.Number(123),
//   			},
//   			HierarchicalChunkingConfiguration: &HierarchicalChunkingConfigurationProperty{
//   				LevelConfigurations: []interface{}{
//   					&HierarchicalChunkingLevelConfigurationProperty{
//   						MaxTokens: jsii.Number(123),
//   					},
//   				},
//   				OverlapTokens: jsii.Number(123),
//   			},
//   			SemanticChunkingConfiguration: &SemanticChunkingConfigurationProperty{
//   				BreakpointPercentileThreshold: jsii.Number(123),
//   				BufferSize: jsii.Number(123),
//   				MaxTokens: jsii.Number(123),
//   			},
//   		},
//   		ParsingConfiguration: &ParsingConfigurationProperty{
//   			BedrockFoundationModelConfiguration: &BedrockFoundationModelConfigurationProperty{
//   				ModelArn: jsii.String("modelArn"),
//   				ParsingPrompt: &ParsingPromptProperty{
//   					ParsingPromptText: jsii.String("parsingPromptText"),
//   				},
//   			},
//   			ParsingStrategy: jsii.String("parsingStrategy"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-knowledgebase.html
//
type CfnKnowledgeBaseMixinProps struct {
	// The description.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-knowledgebase.html#cfn-wisdom-knowledgebase-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The type of knowledge base.
	//
	// Only CUSTOM knowledge bases allow you to upload your own content. EXTERNAL knowledge bases support integrations with third-party systems whose content is synchronized automatically.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-knowledgebase.html#cfn-wisdom-knowledgebase-knowledgebasetype
	//
	KnowledgeBaseType *string `field:"optional" json:"knowledgeBaseType" yaml:"knowledgeBaseType"`
	// The name of the knowledge base.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-knowledgebase.html#cfn-wisdom-knowledgebase-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Information about how to render the content.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-knowledgebase.html#cfn-wisdom-knowledgebase-renderingconfiguration
	//
	RenderingConfiguration interface{} `field:"optional" json:"renderingConfiguration" yaml:"renderingConfiguration"`
	// This customer managed key must have a policy that allows `kms:CreateGrant` and `kms:DescribeKey` permissions to the IAM identity using the key to invoke Wisdom.
	//
	// For more information about setting up a customer managed key for Wisdom, see [Enable Amazon Connect Wisdom for your instance](https://docs.aws.amazon.com/connect/latest/adminguide/enable-wisdom.html) . For information about valid ID values, see [Key identifiers (KeyId)](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-id) in the *AWS Key Management Service Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-knowledgebase.html#cfn-wisdom-knowledgebase-serversideencryptionconfiguration
	//
	ServerSideEncryptionConfiguration interface{} `field:"optional" json:"serverSideEncryptionConfiguration" yaml:"serverSideEncryptionConfiguration"`
	// The source of the knowledge base content.
	//
	// Only set this argument for EXTERNAL or Managed knowledge bases.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-knowledgebase.html#cfn-wisdom-knowledgebase-sourceconfiguration
	//
	SourceConfiguration interface{} `field:"optional" json:"sourceConfiguration" yaml:"sourceConfiguration"`
	// The tags used to organize, track, or control access for this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-knowledgebase.html#cfn-wisdom-knowledgebase-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Contains details about how to ingest the documents in a data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-knowledgebase.html#cfn-wisdom-knowledgebase-vectoringestionconfiguration
	//
	VectorIngestionConfiguration interface{} `field:"optional" json:"vectorIngestionConfiguration" yaml:"vectorIngestionConfiguration"`
}

