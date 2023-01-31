package awsquicksight


// A version of a template.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   templateVersionProperty := &templateVersionProperty{
//   	createdTime: jsii.String("createdTime"),
//   	dataSetConfigurations: []interface{}{
//   		&dataSetConfigurationProperty{
//   			columnGroupSchemaList: []interface{}{
//   				&columnGroupSchemaProperty{
//   					columnGroupColumnSchemaList: []interface{}{
//   						&columnGroupColumnSchemaProperty{
//   							name: jsii.String("name"),
//   						},
//   					},
//   					name: jsii.String("name"),
//   				},
//   			},
//   			dataSetSchema: &dataSetSchemaProperty{
//   				columnSchemaList: []interface{}{
//   					&columnSchemaProperty{
//   						dataType: jsii.String("dataType"),
//   						geographicRole: jsii.String("geographicRole"),
//   						name: jsii.String("name"),
//   					},
//   				},
//   			},
//   			placeholder: jsii.String("placeholder"),
//   		},
//   	},
//   	description: jsii.String("description"),
//   	errors: []interface{}{
//   		&templateErrorProperty{
//   			message: jsii.String("message"),
//   			type: jsii.String("type"),
//   		},
//   	},
//   	sheets: []interface{}{
//   		&sheetProperty{
//   			name: jsii.String("name"),
//   			sheetId: jsii.String("sheetId"),
//   		},
//   	},
//   	sourceEntityArn: jsii.String("sourceEntityArn"),
//   	status: jsii.String("status"),
//   	themeArn: jsii.String("themeArn"),
//   	versionNumber: jsii.Number(123),
//   }
//
type CfnTemplate_TemplateVersionProperty struct {
	// The time that this template version was created.
	CreatedTime *string `field:"optional" json:"createdTime" yaml:"createdTime"`
	// Schema of the dataset identified by the placeholder.
	//
	// Any dashboard created from this template should be bound to new datasets matching the same schema described through this API operation.
	DataSetConfigurations interface{} `field:"optional" json:"dataSetConfigurations" yaml:"dataSetConfigurations"`
	// The description of the template.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Errors associated with this template version.
	Errors interface{} `field:"optional" json:"errors" yaml:"errors"`
	// A list of the associated sheets with the unique identifier and name of each sheet.
	Sheets interface{} `field:"optional" json:"sheets" yaml:"sheets"`
	// The Amazon Resource Name (ARN) of an analysis or template that was used to create this template.
	SourceEntityArn *string `field:"optional" json:"sourceEntityArn" yaml:"sourceEntityArn"`
	// The status that is associated with the template.
	//
	// - `CREATION_IN_PROGRESS`
	// - `CREATION_SUCCESSFUL`
	// - `CREATION_FAILED`
	// - `UPDATE_IN_PROGRESS`
	// - `UPDATE_SUCCESSFUL`
	// - `UPDATE_FAILED`
	// - `DELETED`.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// The ARN of the theme associated with this version of the template.
	ThemeArn *string `field:"optional" json:"themeArn" yaml:"themeArn"`
	// The version number of the template version.
	VersionNumber *float64 `field:"optional" json:"versionNumber" yaml:"versionNumber"`
}

