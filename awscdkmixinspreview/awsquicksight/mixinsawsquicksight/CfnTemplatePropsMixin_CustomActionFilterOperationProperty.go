package mixinsawsquicksight


// The filter operation that filters data included in a visual or in an entire sheet.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   customActionFilterOperationProperty := &CustomActionFilterOperationProperty{
//   	SelectedFieldsConfiguration: &FilterOperationSelectedFieldsConfigurationProperty{
//   		SelectedColumns: []interface{}{
//   			&ColumnIdentifierProperty{
//   				ColumnName: jsii.String("columnName"),
//   				DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   			},
//   		},
//   		SelectedFieldOptions: jsii.String("selectedFieldOptions"),
//   		SelectedFields: []*string{
//   			jsii.String("selectedFields"),
//   		},
//   	},
//   	TargetVisualsConfiguration: &FilterOperationTargetVisualsConfigurationProperty{
//   		SameSheetTargetVisualConfiguration: &SameSheetTargetVisualConfigurationProperty{
//   			TargetVisualOptions: jsii.String("targetVisualOptions"),
//   			TargetVisuals: []*string{
//   				jsii.String("targetVisuals"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-customactionfilteroperation.html
//
type CfnTemplatePropsMixin_CustomActionFilterOperationProperty struct {
	// The configuration that chooses the fields to be filtered.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-customactionfilteroperation.html#cfn-quicksight-template-customactionfilteroperation-selectedfieldsconfiguration
	//
	SelectedFieldsConfiguration interface{} `field:"optional" json:"selectedFieldsConfiguration" yaml:"selectedFieldsConfiguration"`
	// The configuration that chooses the target visuals to be filtered.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-customactionfilteroperation.html#cfn-quicksight-template-customactionfilteroperation-targetvisualsconfiguration
	//
	TargetVisualsConfiguration interface{} `field:"optional" json:"targetVisualsConfiguration" yaml:"targetVisualsConfiguration"`
}

