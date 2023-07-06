package awsquicksight


// The configuration of selected fields in the `CustomActionFilterOperation` .
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filterOperationSelectedFieldsConfigurationProperty := &FilterOperationSelectedFieldsConfigurationProperty{
//   	SelectedColumns: []interface{}{
//   		&ColumnIdentifierProperty{
//   			ColumnName: jsii.String("columnName"),
//   			DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   		},
//   	},
//   	SelectedFieldOptions: jsii.String("selectedFieldOptions"),
//   	SelectedFields: []*string{
//   		jsii.String("selectedFields"),
//   	},
//   }
//
type CfnTemplate_FilterOperationSelectedFieldsConfigurationProperty struct {
	// The selected columns of a dataset.
	SelectedColumns interface{} `field:"optional" json:"selectedColumns" yaml:"selectedColumns"`
	// A structure that contains the options that choose which fields are filtered in the `CustomActionFilterOperation` .
	//
	// Valid values are defined as follows:
	//
	// - `ALL_FIELDS` : Applies the filter operation to all fields.
	SelectedFieldOptions *string `field:"optional" json:"selectedFieldOptions" yaml:"selectedFieldOptions"`
	// Chooses the fields that are filtered in `CustomActionFilterOperation` .
	SelectedFields *[]*string `field:"optional" json:"selectedFields" yaml:"selectedFields"`
}

