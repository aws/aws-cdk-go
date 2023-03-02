package awsquicksight


// Metadata that contains a description for a column.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   columnDescriptionProperty := &columnDescriptionProperty{
//   	text: jsii.String("text"),
//   }
//
type CfnDataSet_ColumnDescriptionProperty struct {
	// The text of a description for a column.
	Text *string `field:"optional" json:"text" yaml:"text"`
}

