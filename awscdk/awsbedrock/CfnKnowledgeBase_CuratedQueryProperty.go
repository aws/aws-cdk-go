package awsbedrock


// Contains configurations for a query, each of which defines information about example queries to help the query engine generate appropriate SQL queries.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   curatedQueryProperty := &CuratedQueryProperty{
//   	NaturalLanguage: jsii.String("naturalLanguage"),
//   	Sql: jsii.String("sql"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-curatedquery.html
//
type CfnKnowledgeBase_CuratedQueryProperty struct {
	// An example natural language query.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-curatedquery.html#cfn-bedrock-knowledgebase-curatedquery-naturallanguage
	//
	NaturalLanguage *string `field:"required" json:"naturalLanguage" yaml:"naturalLanguage"`
	// The SQL equivalent of the natural language query.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-curatedquery.html#cfn-bedrock-knowledgebase-curatedquery-sql
	//
	Sql *string `field:"required" json:"sql" yaml:"sql"`
}

