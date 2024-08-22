package awsbedrock


// Contains details about how to ingest the documents in a data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vectorIngestionConfigurationProperty := &VectorIngestionConfigurationProperty{
//   	ChunkingConfiguration: &ChunkingConfigurationProperty{
//   		ChunkingStrategy: jsii.String("chunkingStrategy"),
//
//   		// the properties below are optional
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
//   		ParsingStrategy: jsii.String("parsingStrategy"),
//
//   		// the properties below are optional
//   		BedrockFoundationModelConfiguration: &BedrockFoundationModelConfigurationProperty{
//   			ModelArn: jsii.String("modelArn"),
//
//   			// the properties below are optional
//   			ParsingPrompt: &ParsingPromptProperty{
//   				ParsingPromptText: jsii.String("parsingPromptText"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-vectoringestionconfiguration.html
//
type CfnDataSource_VectorIngestionConfigurationProperty struct {
	// Details about how to chunk the documents in the data source.
	//
	// A *chunk* refers to an excerpt from a data source that is returned when the knowledge base that it belongs to is queried.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-vectoringestionconfiguration.html#cfn-bedrock-datasource-vectoringestionconfiguration-chunkingconfiguration
	//
	ChunkingConfiguration interface{} `field:"optional" json:"chunkingConfiguration" yaml:"chunkingConfiguration"`
	// A custom document transformer for parsed data source documents.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-vectoringestionconfiguration.html#cfn-bedrock-datasource-vectoringestionconfiguration-customtransformationconfiguration
	//
	CustomTransformationConfiguration interface{} `field:"optional" json:"customTransformationConfiguration" yaml:"customTransformationConfiguration"`
	// A custom parser for data source documents.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-vectoringestionconfiguration.html#cfn-bedrock-datasource-vectoringestionconfiguration-parsingconfiguration
	//
	ParsingConfiguration interface{} `field:"optional" json:"parsingConfiguration" yaml:"parsingConfiguration"`
}

