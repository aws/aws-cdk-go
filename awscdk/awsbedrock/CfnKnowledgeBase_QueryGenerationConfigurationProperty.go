package awsbedrock


// Contains configurations for query generation.
//
// For more information, see [Build a knowledge base by connecting to a structured data source](https://docs.aws.amazon.com/bedrock/latest/userguide/knowledge-base-build-structured.html) in the Amazon Bedrock User Guide..
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   queryGenerationConfigurationProperty := &QueryGenerationConfigurationProperty{
//   	ExecutionTimeoutSeconds: jsii.Number(123),
//   	GenerationContext: &QueryGenerationContextProperty{
//   		CuratedQueries: []interface{}{
//   			&CuratedQueryProperty{
//   				NaturalLanguage: jsii.String("naturalLanguage"),
//   				Sql: jsii.String("sql"),
//   			},
//   		},
//   		Tables: []interface{}{
//   			&QueryGenerationTableProperty{
//   				Name: jsii.String("name"),
//
//   				// the properties below are optional
//   				Columns: []interface{}{
//   					&QueryGenerationColumnProperty{
//   						Description: jsii.String("description"),
//   						Inclusion: jsii.String("inclusion"),
//   						Name: jsii.String("name"),
//   					},
//   				},
//   				Description: jsii.String("description"),
//   				Inclusion: jsii.String("inclusion"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-querygenerationconfiguration.html
//
type CfnKnowledgeBase_QueryGenerationConfigurationProperty struct {
	// The time after which query generation will time out.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-querygenerationconfiguration.html#cfn-bedrock-knowledgebase-querygenerationconfiguration-executiontimeoutseconds
	//
	ExecutionTimeoutSeconds *float64 `field:"optional" json:"executionTimeoutSeconds" yaml:"executionTimeoutSeconds"`
	// Specifies configurations for context to use during query generation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-querygenerationconfiguration.html#cfn-bedrock-knowledgebase-querygenerationconfiguration-generationcontext
	//
	GenerationContext interface{} `field:"optional" json:"generationContext" yaml:"generationContext"`
}

