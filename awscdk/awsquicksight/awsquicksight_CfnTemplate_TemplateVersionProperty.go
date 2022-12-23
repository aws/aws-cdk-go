package awsquicksight


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
	// `CfnTemplate.TemplateVersionProperty.CreatedTime`.
	CreatedTime *string `field:"optional" json:"createdTime" yaml:"createdTime"`
	// `CfnTemplate.TemplateVersionProperty.DataSetConfigurations`.
	DataSetConfigurations interface{} `field:"optional" json:"dataSetConfigurations" yaml:"dataSetConfigurations"`
	// `CfnTemplate.TemplateVersionProperty.Description`.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// `CfnTemplate.TemplateVersionProperty.Errors`.
	Errors interface{} `field:"optional" json:"errors" yaml:"errors"`
	// `CfnTemplate.TemplateVersionProperty.Sheets`.
	Sheets interface{} `field:"optional" json:"sheets" yaml:"sheets"`
	// `CfnTemplate.TemplateVersionProperty.SourceEntityArn`.
	SourceEntityArn *string `field:"optional" json:"sourceEntityArn" yaml:"sourceEntityArn"`
	// `CfnTemplate.TemplateVersionProperty.Status`.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// `CfnTemplate.TemplateVersionProperty.ThemeArn`.
	ThemeArn *string `field:"optional" json:"themeArn" yaml:"themeArn"`
	// `CfnTemplate.TemplateVersionProperty.VersionNumber`.
	VersionNumber *float64 `field:"optional" json:"versionNumber" yaml:"versionNumber"`
}

