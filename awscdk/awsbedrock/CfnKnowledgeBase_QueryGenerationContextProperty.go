package awsbedrock


// Context used to improve query generation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   queryGenerationContextProperty := &QueryGenerationContextProperty{
//   	CuratedQueries: []interface{}{
//   		&CuratedQueryProperty{
//   			NaturalLanguage: jsii.String("naturalLanguage"),
//   			Sql: jsii.String("sql"),
//   		},
//   	},
//   	Tables: []interface{}{
//   		&QueryGenerationTableProperty{
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			Columns: []interface{}{
//   				&QueryGenerationColumnProperty{
//   					Description: jsii.String("description"),
//   					Inclusion: jsii.String("inclusion"),
//   					Name: jsii.String("name"),
//   				},
//   			},
//   			Description: jsii.String("description"),
//   			Inclusion: jsii.String("inclusion"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-querygenerationcontext.html
//
type CfnKnowledgeBase_QueryGenerationContextProperty struct {
	// List of example queries and results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-querygenerationcontext.html#cfn-bedrock-knowledgebase-querygenerationcontext-curatedqueries
	//
	CuratedQueries interface{} `field:"optional" json:"curatedQueries" yaml:"curatedQueries"`
	// List of tables used for Redshift query generation context.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-querygenerationcontext.html#cfn-bedrock-knowledgebase-querygenerationcontext-tables
	//
	Tables interface{} `field:"optional" json:"tables" yaml:"tables"`
}

