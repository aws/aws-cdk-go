package awsquicksight


// A structure that represents the cell value synonym.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cellValueSynonymProperty := &CellValueSynonymProperty{
//   	CellValue: jsii.String("cellValue"),
//   	Synonyms: []*string{
//   		jsii.String("synonyms"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-cellvaluesynonym.html
//
type CfnTopic_CellValueSynonymProperty struct {
	// The cell value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-cellvaluesynonym.html#cfn-quicksight-topic-cellvaluesynonym-cellvalue
	//
	CellValue *string `field:"optional" json:"cellValue" yaml:"cellValue"`
	// Other names or aliases for the cell value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-cellvaluesynonym.html#cfn-quicksight-topic-cellvaluesynonym-synonyms
	//
	Synonyms *[]*string `field:"optional" json:"synonyms" yaml:"synonyms"`
}

