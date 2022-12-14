package awsquicksight


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
	// `CfnTemplate.DataSetConfigurationProperty.ColumnGroupSchemaList`.
	ColumnGroupSchemaList interface{} `field:"optional" json:"columnGroupSchemaList" yaml:"columnGroupSchemaList"`
	// `CfnTemplate.DataSetConfigurationProperty.DataSetSchema`.
	DataSetSchema interface{} `field:"optional" json:"dataSetSchema" yaml:"dataSetSchema"`
	// `CfnTemplate.DataSetConfigurationProperty.Placeholder`.
	Placeholder *string `field:"optional" json:"placeholder" yaml:"placeholder"`
}

