package awsbcmdataexports


// Includes basic information for a data column such as its description, name, and type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   columnProperty := &ColumnProperty{
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcmdataexports-table-column.html
//
type CfnTable_ColumnProperty struct {
	// The description for a column.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcmdataexports-table-column.html#cfn-bcmdataexports-table-column-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The column name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcmdataexports-table-column.html#cfn-bcmdataexports-table-column-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The kind of data a column stores.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcmdataexports-table-column.html#cfn-bcmdataexports-table-column-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

