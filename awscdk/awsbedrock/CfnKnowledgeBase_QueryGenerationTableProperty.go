package awsbedrock


// Tables used for Redshift query generation context.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   queryGenerationTableProperty := &QueryGenerationTableProperty{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Columns: []interface{}{
//   		&QueryGenerationColumnProperty{
//   			Description: jsii.String("description"),
//   			Inclusion: jsii.String("inclusion"),
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	Inclusion: jsii.String("inclusion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-querygenerationtable.html
//
type CfnKnowledgeBase_QueryGenerationTableProperty struct {
	// Query generation table name.
	//
	// Must follow three-part notation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-querygenerationtable.html#cfn-bedrock-knowledgebase-querygenerationtable-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// List of Redshift query generation columns.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-querygenerationtable.html#cfn-bedrock-knowledgebase-querygenerationtable-columns
	//
	Columns interface{} `field:"optional" json:"columns" yaml:"columns"`
	// Description for the attached entity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-querygenerationtable.html#cfn-bedrock-knowledgebase-querygenerationtable-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Include or Exclude status for an entity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-querygenerationtable.html#cfn-bedrock-knowledgebase-querygenerationtable-inclusion
	//
	Inclusion *string `field:"optional" json:"inclusion" yaml:"inclusion"`
}

