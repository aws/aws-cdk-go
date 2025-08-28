package awsbedrock


// Contains configurations for a knowledge base node in a flow.
//
// This node takes a query as the input and returns, as the output, the retrieved responses directly (as an array) or a response generated based on the retrieved responses. For more information, see [Node types in a flow](https://docs.aws.amazon.com/bedrock/latest/userguide/flows-nodes.html) in the Amazon Bedrock User Guide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var additionalModelRequestFields interface{}
//
//   knowledgeBaseFlowNodeConfigurationProperty := &KnowledgeBaseFlowNodeConfigurationProperty{
//   	KnowledgeBaseId: jsii.String("knowledgeBaseId"),
//
//   	// the properties below are optional
//   	GuardrailConfiguration: &GuardrailConfigurationProperty{
//   		GuardrailIdentifier: jsii.String("guardrailIdentifier"),
//   		GuardrailVersion: jsii.String("guardrailVersion"),
//   	},
//   	InferenceConfiguration: &PromptInferenceConfigurationProperty{
//   		Text: &PromptModelInferenceConfigurationProperty{
//   			MaxTokens: jsii.Number(123),
//   			StopSequences: []*string{
//   				jsii.String("stopSequences"),
//   			},
//   			Temperature: jsii.Number(123),
//   			TopP: jsii.Number(123),
//   		},
//   	},
//   	ModelId: jsii.String("modelId"),
//   	NumberOfResults: jsii.Number(123),
//   	OrchestrationConfiguration: &KnowledgeBaseOrchestrationConfigurationProperty{
//   		AdditionalModelRequestFields: additionalModelRequestFields,
//   		InferenceConfig: &PromptInferenceConfigurationProperty{
//   			Text: &PromptModelInferenceConfigurationProperty{
//   				MaxTokens: jsii.Number(123),
//   				StopSequences: []*string{
//   					jsii.String("stopSequences"),
//   				},
//   				Temperature: jsii.Number(123),
//   				TopP: jsii.Number(123),
//   			},
//   		},
//   		PerformanceConfig: &PerformanceConfigurationProperty{
//   			Latency: jsii.String("latency"),
//   		},
//   		PromptTemplate: &KnowledgeBasePromptTemplateProperty{
//   			TextPromptTemplate: jsii.String("textPromptTemplate"),
//   		},
//   	},
//   	PromptTemplate: &KnowledgeBasePromptTemplateProperty{
//   		TextPromptTemplate: jsii.String("textPromptTemplate"),
//   	},
//   	RerankingConfiguration: &VectorSearchRerankingConfigurationProperty{
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
//   		BedrockRerankingConfiguration: &VectorSearchBedrockRerankingConfigurationProperty{
//   			ModelConfiguration: &VectorSearchBedrockRerankingModelConfigurationProperty{
//   				ModelArn: jsii.String("modelArn"),
//
//   				// the properties below are optional
//   				AdditionalModelRequestFields: additionalModelRequestFields,
//   			},
//
//   			// the properties below are optional
//   			MetadataConfiguration: &MetadataConfigurationForRerankingProperty{
//   				SelectionMode: jsii.String("selectionMode"),
//
//   				// the properties below are optional
//   				SelectiveModeConfiguration: &RerankingMetadataSelectiveModeConfigurationProperty{
//   					FieldsToExclude: []interface{}{
//   						&FieldForRerankingProperty{
//   							FieldName: jsii.String("fieldName"),
//   						},
//   					},
//   					FieldsToInclude: []interface{}{
//   						&FieldForRerankingProperty{
//   							FieldName: jsii.String("fieldName"),
//   						},
//   					},
//   				},
//   			},
//   			NumberOfRerankedResults: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-knowledgebaseflownodeconfiguration.html
//
type CfnFlowVersion_KnowledgeBaseFlowNodeConfigurationProperty struct {
	// The unique identifier of the knowledge base to query.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-knowledgebaseflownodeconfiguration.html#cfn-bedrock-flowversion-knowledgebaseflownodeconfiguration-knowledgebaseid
	//
	KnowledgeBaseId *string `field:"required" json:"knowledgeBaseId" yaml:"knowledgeBaseId"`
	// Contains configurations for a guardrail to apply during query and response generation for the knowledge base in this configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-knowledgebaseflownodeconfiguration.html#cfn-bedrock-flowversion-knowledgebaseflownodeconfiguration-guardrailconfiguration
	//
	GuardrailConfiguration interface{} `field:"optional" json:"guardrailConfiguration" yaml:"guardrailConfiguration"`
	// Contains inference configurations for the prompt.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-knowledgebaseflownodeconfiguration.html#cfn-bedrock-flowversion-knowledgebaseflownodeconfiguration-inferenceconfiguration
	//
	InferenceConfiguration interface{} `field:"optional" json:"inferenceConfiguration" yaml:"inferenceConfiguration"`
	// The unique identifier of the model or [inference profile](https://docs.aws.amazon.com/bedrock/latest/userguide/cross-region-inference.html) to use to generate a response from the query results. Omit this field if you want to return the retrieved results as an array.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-knowledgebaseflownodeconfiguration.html#cfn-bedrock-flowversion-knowledgebaseflownodeconfiguration-modelid
	//
	ModelId *string `field:"optional" json:"modelId" yaml:"modelId"`
	// The number of results to retrieve from the knowledge base.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-knowledgebaseflownodeconfiguration.html#cfn-bedrock-flowversion-knowledgebaseflownodeconfiguration-numberofresults
	//
	NumberOfResults *float64 `field:"optional" json:"numberOfResults" yaml:"numberOfResults"`
	// The configuration for orchestrating the retrieval and generation process in the knowledge base node.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-knowledgebaseflownodeconfiguration.html#cfn-bedrock-flowversion-knowledgebaseflownodeconfiguration-orchestrationconfiguration
	//
	OrchestrationConfiguration interface{} `field:"optional" json:"orchestrationConfiguration" yaml:"orchestrationConfiguration"`
	// A custom prompt template to use with the knowledge base for generating responses.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-knowledgebaseflownodeconfiguration.html#cfn-bedrock-flowversion-knowledgebaseflownodeconfiguration-prompttemplate
	//
	PromptTemplate interface{} `field:"optional" json:"promptTemplate" yaml:"promptTemplate"`
	// The configuration for reranking the retrieved results from the knowledge base to improve relevance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-knowledgebaseflownodeconfiguration.html#cfn-bedrock-flowversion-knowledgebaseflownodeconfiguration-rerankingconfiguration
	//
	RerankingConfiguration interface{} `field:"optional" json:"rerankingConfiguration" yaml:"rerankingConfiguration"`
}

