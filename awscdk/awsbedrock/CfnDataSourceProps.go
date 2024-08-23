package awsbedrock


// Properties for defining a `CfnDataSource`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDataSourceProps := &CfnDataSourceProps{
//   	DataSourceConfiguration: &DataSourceConfigurationProperty{
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
//   		ConfluenceConfiguration: &ConfluenceDataSourceConfigurationProperty{
//   			SourceConfiguration: &ConfluenceSourceConfigurationProperty{
//   				AuthType: jsii.String("authType"),
//   				CredentialsSecretArn: jsii.String("credentialsSecretArn"),
//   				HostType: jsii.String("hostType"),
//   				HostUrl: jsii.String("hostUrl"),
//   			},
//
//   			// the properties below are optional
//   			CrawlerConfiguration: &ConfluenceCrawlerConfigurationProperty{
//   				FilterConfiguration: &CrawlFilterConfigurationProperty{
//   					Type: jsii.String("type"),
//
//   					// the properties below are optional
//   					PatternObjectFilter: &PatternObjectFilterConfigurationProperty{
//   						Filters: []interface{}{
//   							&PatternObjectFilterProperty{
//   								ObjectType: jsii.String("objectType"),
//
//   								// the properties below are optional
//   								ExclusionFilters: []*string{
//   									jsii.String("exclusionFilters"),
//   								},
//   								InclusionFilters: []*string{
//   									jsii.String("inclusionFilters"),
//   								},
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   		S3Configuration: &S3DataSourceConfigurationProperty{
//   			BucketArn: jsii.String("bucketArn"),
//
//   			// the properties below are optional
//   			BucketOwnerAccountId: jsii.String("bucketOwnerAccountId"),
//   			InclusionPrefixes: []*string{
//   				jsii.String("inclusionPrefixes"),
//   			},
//   		},
//   		SalesforceConfiguration: &SalesforceDataSourceConfigurationProperty{
//   			SourceConfiguration: &SalesforceSourceConfigurationProperty{
//   				AuthType: jsii.String("authType"),
//   				CredentialsSecretArn: jsii.String("credentialsSecretArn"),
//   				HostUrl: jsii.String("hostUrl"),
//   			},
//
//   			// the properties below are optional
//   			CrawlerConfiguration: &SalesforceCrawlerConfigurationProperty{
//   				FilterConfiguration: &CrawlFilterConfigurationProperty{
//   					Type: jsii.String("type"),
//
//   					// the properties below are optional
//   					PatternObjectFilter: &PatternObjectFilterConfigurationProperty{
//   						Filters: []interface{}{
//   							&PatternObjectFilterProperty{
//   								ObjectType: jsii.String("objectType"),
//
//   								// the properties below are optional
//   								ExclusionFilters: []*string{
//   									jsii.String("exclusionFilters"),
//   								},
//   								InclusionFilters: []*string{
//   									jsii.String("inclusionFilters"),
//   								},
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   		SharePointConfiguration: &SharePointDataSourceConfigurationProperty{
//   			SourceConfiguration: &SharePointSourceConfigurationProperty{
//   				AuthType: jsii.String("authType"),
//   				CredentialsSecretArn: jsii.String("credentialsSecretArn"),
//   				Domain: jsii.String("domain"),
//   				HostType: jsii.String("hostType"),
//   				SiteUrls: []*string{
//   					jsii.String("siteUrls"),
//   				},
//
//   				// the properties below are optional
//   				TenantId: jsii.String("tenantId"),
//   			},
//
//   			// the properties below are optional
//   			CrawlerConfiguration: &SharePointCrawlerConfigurationProperty{
//   				FilterConfiguration: &CrawlFilterConfigurationProperty{
//   					Type: jsii.String("type"),
//
//   					// the properties below are optional
//   					PatternObjectFilter: &PatternObjectFilterConfigurationProperty{
//   						Filters: []interface{}{
//   							&PatternObjectFilterProperty{
//   								ObjectType: jsii.String("objectType"),
//
//   								// the properties below are optional
//   								ExclusionFilters: []*string{
//   									jsii.String("exclusionFilters"),
//   								},
//   								InclusionFilters: []*string{
//   									jsii.String("inclusionFilters"),
//   								},
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   		WebConfiguration: &WebDataSourceConfigurationProperty{
//   			SourceConfiguration: &WebSourceConfigurationProperty{
//   				UrlConfiguration: &UrlConfigurationProperty{
//   					SeedUrls: []interface{}{
//   						&SeedUrlProperty{
//   							Url: jsii.String("url"),
//   						},
//   					},
//   				},
//   			},
//
//   			// the properties below are optional
//   			CrawlerConfiguration: &WebCrawlerConfigurationProperty{
//   				CrawlerLimits: &WebCrawlerLimitsProperty{
//   					RateLimit: jsii.Number(123),
//   				},
//   				ExclusionFilters: []*string{
//   					jsii.String("exclusionFilters"),
//   				},
//   				InclusionFilters: []*string{
//   					jsii.String("inclusionFilters"),
//   				},
//   				Scope: jsii.String("scope"),
//   			},
//   		},
//   	},
//   	KnowledgeBaseId: jsii.String("knowledgeBaseId"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	DataDeletionPolicy: jsii.String("dataDeletionPolicy"),
//   	Description: jsii.String("description"),
//   	ServerSideEncryptionConfiguration: &ServerSideEncryptionConfigurationProperty{
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
//   	},
//   	VectorIngestionConfiguration: &VectorIngestionConfigurationProperty{
//   		ChunkingConfiguration: &ChunkingConfigurationProperty{
//   			ChunkingStrategy: jsii.String("chunkingStrategy"),
//
//   			// the properties below are optional
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
//   			ParsingStrategy: jsii.String("parsingStrategy"),
//
//   			// the properties below are optional
//   			BedrockFoundationModelConfiguration: &BedrockFoundationModelConfigurationProperty{
//   				ModelArn: jsii.String("modelArn"),
//
//   				// the properties below are optional
//   				ParsingPrompt: &ParsingPromptProperty{
//   					ParsingPromptText: jsii.String("parsingPromptText"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-datasource.html
//
type CfnDataSourceProps struct {
	// The connection configuration for the data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-datasource.html#cfn-bedrock-datasource-datasourceconfiguration
	//
	DataSourceConfiguration interface{} `field:"required" json:"dataSourceConfiguration" yaml:"dataSourceConfiguration"`
	// The unique identifier of the knowledge base to which the data source belongs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-datasource.html#cfn-bedrock-datasource-knowledgebaseid
	//
	KnowledgeBaseId *string `field:"required" json:"knowledgeBaseId" yaml:"knowledgeBaseId"`
	// The name of the data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-datasource.html#cfn-bedrock-datasource-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The data deletion policy for the data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-datasource.html#cfn-bedrock-datasource-datadeletionpolicy
	//
	DataDeletionPolicy *string `field:"optional" json:"dataDeletionPolicy" yaml:"dataDeletionPolicy"`
	// The description of the data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-datasource.html#cfn-bedrock-datasource-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Contains details about the configuration of the server-side encryption.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-datasource.html#cfn-bedrock-datasource-serversideencryptionconfiguration
	//
	ServerSideEncryptionConfiguration interface{} `field:"optional" json:"serverSideEncryptionConfiguration" yaml:"serverSideEncryptionConfiguration"`
	// Contains details about how to ingest the documents in the data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-datasource.html#cfn-bedrock-datasource-vectoringestionconfiguration
	//
	VectorIngestionConfiguration interface{} `field:"optional" json:"vectorIngestionConfiguration" yaml:"vectorIngestionConfiguration"`
}

