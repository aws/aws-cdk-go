package awsquicksight


// Dataset configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataSetConfigurationProperty := &DataSetConfigurationProperty{
//   	ColumnGroupSchemaList: []interface{}{
//   		&ColumnGroupSchemaProperty{
//   			ColumnGroupColumnSchemaList: []interface{}{
//   				&ColumnGroupColumnSchemaProperty{
//   					Name: jsii.String("name"),
//   				},
//   			},
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	DataSetSchema: &DataSetSchemaProperty{
//   		ColumnSchemaList: []interface{}{
//   			&ColumnSchemaProperty{
//   				DataType: jsii.String("dataType"),
//   				GeographicRole: jsii.String("geographicRole"),
//   				Name: jsii.String("name"),
//   			},
//   		},
//   	},
//   	Placeholder: jsii.String("placeholder"),
//   }
//
type CfnTemplate_DataSetConfigurationProperty struct {
	// A structure containing the list of column group schemas.
	ColumnGroupSchemaList interface{} `field:"optional" json:"columnGroupSchemaList" yaml:"columnGroupSchemaList"`
	// Dataset schema.
	DataSetSchema interface{} `field:"optional" json:"dataSetSchema" yaml:"dataSetSchema"`
	// Placeholder.
	Placeholder *string `field:"optional" json:"placeholder" yaml:"placeholder"`
}

