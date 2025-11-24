package mixinsawsquicksight


// A version of a template.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   templateVersionProperty := &TemplateVersionProperty{
//   	CreatedTime: jsii.String("createdTime"),
//   	DataSetConfigurations: []interface{}{
//   		&DataSetConfigurationProperty{
//   			ColumnGroupSchemaList: []interface{}{
//   				&ColumnGroupSchemaProperty{
//   					ColumnGroupColumnSchemaList: []interface{}{
//   						&ColumnGroupColumnSchemaProperty{
//   							Name: jsii.String("name"),
//   						},
//   					},
//   					Name: jsii.String("name"),
//   				},
//   			},
//   			DataSetSchema: &DataSetSchemaProperty{
//   				ColumnSchemaList: []interface{}{
//   					&ColumnSchemaProperty{
//   						DataType: jsii.String("dataType"),
//   						GeographicRole: jsii.String("geographicRole"),
//   						Name: jsii.String("name"),
//   					},
//   				},
//   			},
//   			Placeholder: jsii.String("placeholder"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	Errors: []interface{}{
//   		&TemplateErrorProperty{
//   			Message: jsii.String("message"),
//   			Type: jsii.String("type"),
//   			ViolatedEntities: []interface{}{
//   				&EntityProperty{
//   					Path: jsii.String("path"),
//   				},
//   			},
//   		},
//   	},
//   	Sheets: []interface{}{
//   		&SheetProperty{
//   			Name: jsii.String("name"),
//   			SheetId: jsii.String("sheetId"),
//   		},
//   	},
//   	SourceEntityArn: jsii.String("sourceEntityArn"),
//   	Status: jsii.String("status"),
//   	ThemeArn: jsii.String("themeArn"),
//   	VersionNumber: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-templateversion.html
//
type CfnTemplatePropsMixin_TemplateVersionProperty struct {
	// The time that this template version was created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-templateversion.html#cfn-quicksight-template-templateversion-createdtime
	//
	CreatedTime *string `field:"optional" json:"createdTime" yaml:"createdTime"`
	// Schema of the dataset identified by the placeholder.
	//
	// Any dashboard created from this template should be bound to new datasets matching the same schema described through this API operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-templateversion.html#cfn-quicksight-template-templateversion-datasetconfigurations
	//
	DataSetConfigurations interface{} `field:"optional" json:"dataSetConfigurations" yaml:"dataSetConfigurations"`
	// The description of the template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-templateversion.html#cfn-quicksight-template-templateversion-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Errors associated with this template version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-templateversion.html#cfn-quicksight-template-templateversion-errors
	//
	Errors interface{} `field:"optional" json:"errors" yaml:"errors"`
	// A list of the associated sheets with the unique identifier and name of each sheet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-templateversion.html#cfn-quicksight-template-templateversion-sheets
	//
	Sheets interface{} `field:"optional" json:"sheets" yaml:"sheets"`
	// The Amazon Resource Name (ARN) of an analysis or template that was used to create this template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-templateversion.html#cfn-quicksight-template-templateversion-sourceentityarn
	//
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-templateversion.html#cfn-quicksight-template-templateversion-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// The ARN of the theme associated with this version of the template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-templateversion.html#cfn-quicksight-template-templateversion-themearn
	//
	ThemeArn *string `field:"optional" json:"themeArn" yaml:"themeArn"`
	// The version number of the template version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-templateversion.html#cfn-quicksight-template-templateversion-versionnumber
	//
	VersionNumber *float64 `field:"optional" json:"versionNumber" yaml:"versionNumber"`
}

