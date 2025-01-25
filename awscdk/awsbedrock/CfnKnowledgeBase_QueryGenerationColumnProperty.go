package awsbedrock


// Redshift query generation column.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   queryGenerationColumnProperty := &QueryGenerationColumnProperty{
//   	Description: jsii.String("description"),
//   	Inclusion: jsii.String("inclusion"),
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-querygenerationcolumn.html
//
type CfnKnowledgeBase_QueryGenerationColumnProperty struct {
	// Description for the attached entity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-querygenerationcolumn.html#cfn-bedrock-knowledgebase-querygenerationcolumn-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Include or Exclude status for an entity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-querygenerationcolumn.html#cfn-bedrock-knowledgebase-querygenerationcolumn-inclusion
	//
	Inclusion *string `field:"optional" json:"inclusion" yaml:"inclusion"`
	// Query generation column name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-querygenerationcolumn.html#cfn-bedrock-knowledgebase-querygenerationcolumn-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

