package previewawsbedrockmixins


// Contains information about a column in the current table for the query engine to consider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   queryGenerationColumnProperty := &QueryGenerationColumnProperty{
//   	Description: jsii.String("description"),
//   	Inclusion: jsii.String("inclusion"),
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-querygenerationcolumn.html
//
type CfnKnowledgeBasePropsMixin_QueryGenerationColumnProperty struct {
	// A description of the column that helps the query engine understand the contents of the column.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-querygenerationcolumn.html#cfn-bedrock-knowledgebase-querygenerationcolumn-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies whether to include or exclude the column during query generation.
	//
	// If you specify `EXCLUDE` , the column will be ignored. If you specify `INCLUDE` , all other columns in the table will be ignored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-querygenerationcolumn.html#cfn-bedrock-knowledgebase-querygenerationcolumn-inclusion
	//
	Inclusion *string `field:"optional" json:"inclusion" yaml:"inclusion"`
	// The name of the column for which the other fields in this object apply.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-querygenerationcolumn.html#cfn-bedrock-knowledgebase-querygenerationcolumn-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

