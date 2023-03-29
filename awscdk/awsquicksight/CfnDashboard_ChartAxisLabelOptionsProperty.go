package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   chartAxisLabelOptionsProperty := &ChartAxisLabelOptionsProperty{
//   	AxisLabelOptions: []interface{}{
//   		&AxisLabelOptionsProperty{
//   			ApplyTo: &AxisLabelReferenceOptionsProperty{
//   				Column: &ColumnIdentifierProperty{
//   					ColumnName: jsii.String("columnName"),
//   					DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   				},
//   				FieldId: jsii.String("fieldId"),
//   			},
//   			CustomLabel: jsii.String("customLabel"),
//   			FontConfiguration: &FontConfigurationProperty{
//   				FontColor: jsii.String("fontColor"),
//   				FontDecoration: jsii.String("fontDecoration"),
//   				FontSize: &FontSizeProperty{
//   					Relative: jsii.String("relative"),
//   				},
//   				FontStyle: jsii.String("fontStyle"),
//   				FontWeight: &FontWeightProperty{
//   					Name: jsii.String("name"),
//   				},
//   			},
//   		},
//   	},
//   	SortIconVisibility: jsii.String("sortIconVisibility"),
//   	Visibility: jsii.String("visibility"),
//   }
//
type CfnDashboard_ChartAxisLabelOptionsProperty struct {
	// `CfnDashboard.ChartAxisLabelOptionsProperty.AxisLabelOptions`.
	AxisLabelOptions interface{} `field:"optional" json:"axisLabelOptions" yaml:"axisLabelOptions"`
	// `CfnDashboard.ChartAxisLabelOptionsProperty.SortIconVisibility`.
	SortIconVisibility *string `field:"optional" json:"sortIconVisibility" yaml:"sortIconVisibility"`
	// `CfnDashboard.ChartAxisLabelOptionsProperty.Visibility`.
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

