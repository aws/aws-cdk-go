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
type CfnTopic_CellValueSynonymProperty struct {
	// The cell value.
	CellValue *string `field:"optional" json:"cellValue" yaml:"cellValue"`
	// Other names or aliases for the cell value.
	Synonyms *[]*string `field:"optional" json:"synonyms" yaml:"synonyms"`
}

