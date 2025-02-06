package awsbedrock


// >Contains configurations for context to use during query generation.
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
	// An array of objects, each of which defines information about example queries to help the query engine generate appropriate SQL queries.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-querygenerationcontext.html#cfn-bedrock-knowledgebase-querygenerationcontext-curatedqueries
	//
	CuratedQueries interface{} `field:"optional" json:"curatedQueries" yaml:"curatedQueries"`
	// An array of objects, each of which defines information about a table in the database.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-querygenerationcontext.html#cfn-bedrock-knowledgebase-querygenerationcontext-tables
	//
	Tables interface{} `field:"optional" json:"tables" yaml:"tables"`
}

