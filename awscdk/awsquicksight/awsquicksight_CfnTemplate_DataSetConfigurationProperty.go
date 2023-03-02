package awsquicksight


// Dataset configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataSetConfigurationProperty := &dataSetConfigurationProperty{
//   	columnGroupSchemaList: []interface{}{
//   		&columnGroupSchemaProperty{
//   			columnGroupColumnSchemaList: []interface{}{
//   				&columnGroupColumnSchemaProperty{
//   					name: jsii.String("name"),
//   				},
//   			},
//   			name: jsii.String("name"),
//   		},
//   	},
//   	dataSetSchema: &dataSetSchemaProperty{
//   		columnSchemaList: []interface{}{
//   			&columnSchemaProperty{
//   				dataType: jsii.String("dataType"),
//   				geographicRole: jsii.String("geographicRole"),
//   				name: jsii.String("name"),
//   			},
//   		},
//   	},
//   	placeholder: jsii.String("placeholder"),
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

