package awsbedrock


// Curated query or question and answer pair.
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
	// Question for the curated query.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-curatedquery.html#cfn-bedrock-knowledgebase-curatedquery-naturallanguage
	//
	NaturalLanguage *string `field:"required" json:"naturalLanguage" yaml:"naturalLanguage"`
	// Answer for the curated query.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-curatedquery.html#cfn-bedrock-knowledgebase-curatedquery-sql
	//
	Sql *string `field:"required" json:"sql" yaml:"sql"`
}

