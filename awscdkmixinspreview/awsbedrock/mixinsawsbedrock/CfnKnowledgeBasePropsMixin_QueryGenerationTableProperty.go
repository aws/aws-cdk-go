package mixinsawsbedrock


// Contains information about a table for the query engine to consider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   queryGenerationTableProperty := &QueryGenerationTableProperty{
//   	Columns: []interface{}{
//   		&QueryGenerationColumnProperty{
//   			Description: jsii.String("description"),
//   			Inclusion: jsii.String("inclusion"),
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	Inclusion: jsii.String("inclusion"),
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-querygenerationtable.html
//
type CfnKnowledgeBasePropsMixin_QueryGenerationTableProperty struct {
	// An array of objects, each of which defines information about a column in the table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-querygenerationtable.html#cfn-bedrock-knowledgebase-querygenerationtable-columns
	//
	Columns interface{} `field:"optional" json:"columns" yaml:"columns"`
	// A description of the table that helps the query engine understand the contents of the table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-querygenerationtable.html#cfn-bedrock-knowledgebase-querygenerationtable-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies whether to include or exclude the table during query generation.
	//
	// If you specify `EXCLUDE` , the table will be ignored. If you specify `INCLUDE` , all other tables will be ignored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-querygenerationtable.html#cfn-bedrock-knowledgebase-querygenerationtable-inclusion
	//
	Inclusion *string `field:"optional" json:"inclusion" yaml:"inclusion"`
	// The name of the table for which the other fields in this object apply.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-querygenerationtable.html#cfn-bedrock-knowledgebase-querygenerationtable-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

