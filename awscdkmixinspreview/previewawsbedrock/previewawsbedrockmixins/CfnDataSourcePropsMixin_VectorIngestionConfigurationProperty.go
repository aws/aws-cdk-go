package previewawsbedrockmixins


// Contains details about how to ingest the documents in a data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   vectorIngestionConfigurationProperty := &VectorIngestionConfigurationProperty{
//   	ChunkingConfiguration: &ChunkingConfigurationProperty{
//   		ChunkingStrategy: jsii.String("chunkingStrategy"),
//   		FixedSizeChunkingConfiguration: &FixedSizeChunkingConfigurationProperty{
//   			MaxTokens: jsii.Number(123),
//   			OverlapPercentage: jsii.Number(123),
//   		},
//   		HierarchicalChunkingConfiguration: &HierarchicalChunkingConfigurationProperty{
//   			LevelConfigurations: []interface{}{
//   				&HierarchicalChunkingLevelConfigurationProperty{
//   					MaxTokens: jsii.Number(123),
//   				},
//   			},
//   			OverlapTokens: jsii.Number(123),
//   		},
//   		SemanticChunkingConfiguration: &SemanticChunkingConfigurationProperty{
//   			BreakpointPercentileThreshold: jsii.Number(123),
//   			BufferSize: jsii.Number(123),
//   			MaxTokens: jsii.Number(123),
//   		},
//   	},
//   	ContextEnrichmentConfiguration: &ContextEnrichmentConfigurationProperty{
//   		BedrockFoundationModelConfiguration: &BedrockFoundationModelContextEnrichmentConfigurationProperty{
//   			EnrichmentStrategyConfiguration: &EnrichmentStrategyConfigurationProperty{
//   				Method: jsii.String("method"),
//   			},
//   			ModelArn: jsii.String("modelArn"),
//   		},
//   		Type: jsii.String("type"),
//   	},
//   	CustomTransformationConfiguration: &CustomTransformationConfigurationProperty{
//   		IntermediateStorage: &IntermediateStorageProperty{
//   			S3Location: &S3LocationProperty{
//   				Uri: jsii.String("uri"),
//   			},
//   		},
//   		Transformations: []interface{}{
//   			&TransformationProperty{
//   				StepToApply: jsii.String("stepToApply"),
//   				TransformationFunction: &TransformationFunctionProperty{
//   					TransformationLambdaConfiguration: &TransformationLambdaConfigurationProperty{
//   						LambdaArn: jsii.String("lambdaArn"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	ParsingConfiguration: &ParsingConfigurationProperty{
//   		BedrockDataAutomationConfiguration: &BedrockDataAutomationConfigurationProperty{
//   			ParsingModality: jsii.String("parsingModality"),
//   		},
//   		BedrockFoundationModelConfiguration: &BedrockFoundationModelConfigurationProperty{
//   			ModelArn: jsii.String("modelArn"),
//   			ParsingModality: jsii.String("parsingModality"),
//   			ParsingPrompt: &ParsingPromptProperty{
//   				ParsingPromptText: jsii.String("parsingPromptText"),
//   			},
//   		},
//   		ParsingStrategy: jsii.String("parsingStrategy"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-vectoringestionconfiguration.html
//
type CfnDataSourcePropsMixin_VectorIngestionConfigurationProperty struct {
	// Details about how to chunk the documents in the data source.
	//
	// A *chunk* refers to an excerpt from a data source that is returned when the knowledge base that it belongs to is queried.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-vectoringestionconfiguration.html#cfn-bedrock-datasource-vectoringestionconfiguration-chunkingconfiguration
	//
	ChunkingConfiguration interface{} `field:"optional" json:"chunkingConfiguration" yaml:"chunkingConfiguration"`
	// The context enrichment configuration used for ingestion of the data into the vector store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-vectoringestionconfiguration.html#cfn-bedrock-datasource-vectoringestionconfiguration-contextenrichmentconfiguration
	//
	ContextEnrichmentConfiguration interface{} `field:"optional" json:"contextEnrichmentConfiguration" yaml:"contextEnrichmentConfiguration"`
	// A custom document transformer for parsed data source documents.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-vectoringestionconfiguration.html#cfn-bedrock-datasource-vectoringestionconfiguration-customtransformationconfiguration
	//
	CustomTransformationConfiguration interface{} `field:"optional" json:"customTransformationConfiguration" yaml:"customTransformationConfiguration"`
	// Configurations for a parser to use for parsing documents in your data source.
	//
	// If you exclude this field, the default parser will be used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-vectoringestionconfiguration.html#cfn-bedrock-datasource-vectoringestionconfiguration-parsingconfiguration
	//
	ParsingConfiguration interface{} `field:"optional" json:"parsingConfiguration" yaml:"parsingConfiguration"`
}

