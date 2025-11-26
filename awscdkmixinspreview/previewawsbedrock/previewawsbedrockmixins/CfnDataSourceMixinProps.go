package previewawsbedrockmixins


// Properties for CfnDataSourcePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDataSourceMixinProps := &CfnDataSourceMixinProps{
//   	DataDeletionPolicy: jsii.String("dataDeletionPolicy"),
//   	DataSourceConfiguration: &DataSourceConfigurationProperty{
//   		ConfluenceConfiguration: &ConfluenceDataSourceConfigurationProperty{
//   			CrawlerConfiguration: &ConfluenceCrawlerConfigurationProperty{
//   				FilterConfiguration: &CrawlFilterConfigurationProperty{
//   					PatternObjectFilter: &PatternObjectFilterConfigurationProperty{
//   						Filters: []interface{}{
//   							&PatternObjectFilterProperty{
//   								ExclusionFilters: []*string{
//   									jsii.String("exclusionFilters"),
//   								},
//   								InclusionFilters: []*string{
//   									jsii.String("inclusionFilters"),
//   								},
//   								ObjectType: jsii.String("objectType"),
//   							},
//   						},
//   					},
//   					Type: jsii.String("type"),
//   				},
//   			},
//   			SourceConfiguration: &ConfluenceSourceConfigurationProperty{
//   				AuthType: jsii.String("authType"),
//   				CredentialsSecretArn: jsii.String("credentialsSecretArn"),
//   				HostType: jsii.String("hostType"),
//   				HostUrl: jsii.String("hostUrl"),
//   			},
//   		},
//   		S3Configuration: &S3DataSourceConfigurationProperty{
//   			BucketArn: jsii.String("bucketArn"),
//   			BucketOwnerAccountId: jsii.String("bucketOwnerAccountId"),
//   			InclusionPrefixes: []*string{
//   				jsii.String("inclusionPrefixes"),
//   			},
//   		},
//   		SalesforceConfiguration: &SalesforceDataSourceConfigurationProperty{
//   			CrawlerConfiguration: &SalesforceCrawlerConfigurationProperty{
//   				FilterConfiguration: &CrawlFilterConfigurationProperty{
//   					PatternObjectFilter: &PatternObjectFilterConfigurationProperty{
//   						Filters: []interface{}{
//   							&PatternObjectFilterProperty{
//   								ExclusionFilters: []*string{
//   									jsii.String("exclusionFilters"),
//   								},
//   								InclusionFilters: []*string{
//   									jsii.String("inclusionFilters"),
//   								},
//   								ObjectType: jsii.String("objectType"),
//   							},
//   						},
//   					},
//   					Type: jsii.String("type"),
//   				},
//   			},
//   			SourceConfiguration: &SalesforceSourceConfigurationProperty{
//   				AuthType: jsii.String("authType"),
//   				CredentialsSecretArn: jsii.String("credentialsSecretArn"),
//   				HostUrl: jsii.String("hostUrl"),
//   			},
//   		},
//   		SharePointConfiguration: &SharePointDataSourceConfigurationProperty{
//   			CrawlerConfiguration: &SharePointCrawlerConfigurationProperty{
//   				FilterConfiguration: &CrawlFilterConfigurationProperty{
//   					PatternObjectFilter: &PatternObjectFilterConfigurationProperty{
//   						Filters: []interface{}{
//   							&PatternObjectFilterProperty{
//   								ExclusionFilters: []*string{
//   									jsii.String("exclusionFilters"),
//   								},
//   								InclusionFilters: []*string{
//   									jsii.String("inclusionFilters"),
//   								},
//   								ObjectType: jsii.String("objectType"),
//   							},
//   						},
//   					},
//   					Type: jsii.String("type"),
//   				},
//   			},
//   			SourceConfiguration: &SharePointSourceConfigurationProperty{
//   				AuthType: jsii.String("authType"),
//   				CredentialsSecretArn: jsii.String("credentialsSecretArn"),
//   				Domain: jsii.String("domain"),
//   				HostType: jsii.String("hostType"),
//   				SiteUrls: []*string{
//   					jsii.String("siteUrls"),
//   				},
//   				TenantId: jsii.String("tenantId"),
//   			},
//   		},
//   		Type: jsii.String("type"),
//   		WebConfiguration: &WebDataSourceConfigurationProperty{
//   			CrawlerConfiguration: &WebCrawlerConfigurationProperty{
//   				CrawlerLimits: &WebCrawlerLimitsProperty{
//   					MaxPages: jsii.Number(123),
//   					RateLimit: jsii.Number(123),
//   				},
//   				ExclusionFilters: []*string{
//   					jsii.String("exclusionFilters"),
//   				},
//   				InclusionFilters: []*string{
//   					jsii.String("inclusionFilters"),
//   				},
//   				Scope: jsii.String("scope"),
//   				UserAgent: jsii.String("userAgent"),
//   				UserAgentHeader: jsii.String("userAgentHeader"),
//   			},
//   			SourceConfiguration: &WebSourceConfigurationProperty{
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
//   	Description: jsii.String("description"),
//   	KnowledgeBaseId: jsii.String("knowledgeBaseId"),
//   	Name: jsii.String("name"),
//   	ServerSideEncryptionConfiguration: &ServerSideEncryptionConfigurationProperty{
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
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
//   		ContextEnrichmentConfiguration: &ContextEnrichmentConfigurationProperty{
//   			BedrockFoundationModelConfiguration: &BedrockFoundationModelContextEnrichmentConfigurationProperty{
//   				EnrichmentStrategyConfiguration: &EnrichmentStrategyConfigurationProperty{
//   					Method: jsii.String("method"),
//   				},
//   				ModelArn: jsii.String("modelArn"),
//   			},
//   			Type: jsii.String("type"),
//   		},
//   		CustomTransformationConfiguration: &CustomTransformationConfigurationProperty{
//   			IntermediateStorage: &IntermediateStorageProperty{
//   				S3Location: &S3LocationProperty{
//   					Uri: jsii.String("uri"),
//   				},
//   			},
//   			Transformations: []interface{}{
//   				&TransformationProperty{
//   					StepToApply: jsii.String("stepToApply"),
//   					TransformationFunction: &TransformationFunctionProperty{
//   						TransformationLambdaConfiguration: &TransformationLambdaConfigurationProperty{
//   							LambdaArn: jsii.String("lambdaArn"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		ParsingConfiguration: &ParsingConfigurationProperty{
//   			BedrockDataAutomationConfiguration: &BedrockDataAutomationConfigurationProperty{
//   				ParsingModality: jsii.String("parsingModality"),
//   			},
//   			BedrockFoundationModelConfiguration: &BedrockFoundationModelConfigurationProperty{
//   				ModelArn: jsii.String("modelArn"),
//   				ParsingModality: jsii.String("parsingModality"),
//   				ParsingPrompt: &ParsingPromptProperty{
//   					ParsingPromptText: jsii.String("parsingPromptText"),
//   				},
//   			},
//   			ParsingStrategy: jsii.String("parsingStrategy"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-datasource.html
//
type CfnDataSourceMixinProps struct {
	// The data deletion policy for the data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-datasource.html#cfn-bedrock-datasource-datadeletionpolicy
	//
	DataDeletionPolicy *string `field:"optional" json:"dataDeletionPolicy" yaml:"dataDeletionPolicy"`
	// The connection configuration for the data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-datasource.html#cfn-bedrock-datasource-datasourceconfiguration
	//
	DataSourceConfiguration interface{} `field:"optional" json:"dataSourceConfiguration" yaml:"dataSourceConfiguration"`
	// The description of the data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-datasource.html#cfn-bedrock-datasource-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The unique identifier of the knowledge base to which the data source belongs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-datasource.html#cfn-bedrock-datasource-knowledgebaseid
	//
	KnowledgeBaseId *string `field:"optional" json:"knowledgeBaseId" yaml:"knowledgeBaseId"`
	// The name of the data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-datasource.html#cfn-bedrock-datasource-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Contains details about the configuration of the server-side encryption.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-datasource.html#cfn-bedrock-datasource-serversideencryptionconfiguration
	//
	ServerSideEncryptionConfiguration interface{} `field:"optional" json:"serverSideEncryptionConfiguration" yaml:"serverSideEncryptionConfiguration"`
	// Contains details about how to ingest the documents in the data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-datasource.html#cfn-bedrock-datasource-vectoringestionconfiguration
	//
	VectorIngestionConfiguration interface{} `field:"optional" json:"vectorIngestionConfiguration" yaml:"vectorIngestionConfiguration"`
}

