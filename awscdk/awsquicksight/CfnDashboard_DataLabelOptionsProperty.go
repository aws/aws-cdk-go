package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataLabelOptionsProperty := &DataLabelOptionsProperty{
//   	CategoryLabelVisibility: jsii.String("categoryLabelVisibility"),
//   	DataLabelTypes: []interface{}{
//   		&DataLabelTypeProperty{
//   			DataPathLabelType: &DataPathLabelTypeProperty{
//   				FieldId: jsii.String("fieldId"),
//   				FieldValue: jsii.String("fieldValue"),
//   				Visibility: jsii.String("visibility"),
//   			},
//   			FieldLabelType: &FieldLabelTypeProperty{
//   				FieldId: jsii.String("fieldId"),
//   				Visibility: jsii.String("visibility"),
//   			},
//   			MaximumLabelType: &MaximumLabelTypeProperty{
//   				Visibility: jsii.String("visibility"),
//   			},
//   			MinimumLabelType: &MinimumLabelTypeProperty{
//   				Visibility: jsii.String("visibility"),
//   			},
//   			RangeEndsLabelType: &RangeEndsLabelTypeProperty{
//   				Visibility: jsii.String("visibility"),
//   			},
//   		},
//   	},
//   	LabelColor: jsii.String("labelColor"),
//   	LabelContent: jsii.String("labelContent"),
//   	LabelFontConfiguration: &FontConfigurationProperty{
//   		FontColor: jsii.String("fontColor"),
//   		FontDecoration: jsii.String("fontDecoration"),
//   		FontSize: &FontSizeProperty{
//   			Relative: jsii.String("relative"),
//   		},
//   		FontStyle: jsii.String("fontStyle"),
//   		FontWeight: &FontWeightProperty{
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	MeasureLabelVisibility: jsii.String("measureLabelVisibility"),
//   	Overlap: jsii.String("overlap"),
//   	Position: jsii.String("position"),
//   	Visibility: jsii.String("visibility"),
//   }
//
type CfnDashboard_DataLabelOptionsProperty struct {
	// `CfnDashboard.DataLabelOptionsProperty.CategoryLabelVisibility`.
	CategoryLabelVisibility *string `field:"optional" json:"categoryLabelVisibility" yaml:"categoryLabelVisibility"`
	// `CfnDashboard.DataLabelOptionsProperty.DataLabelTypes`.
	DataLabelTypes interface{} `field:"optional" json:"dataLabelTypes" yaml:"dataLabelTypes"`
	// `CfnDashboard.DataLabelOptionsProperty.LabelColor`.
	LabelColor *string `field:"optional" json:"labelColor" yaml:"labelColor"`
	// `CfnDashboard.DataLabelOptionsProperty.LabelContent`.
	LabelContent *string `field:"optional" json:"labelContent" yaml:"labelContent"`
	// `CfnDashboard.DataLabelOptionsProperty.LabelFontConfiguration`.
	LabelFontConfiguration interface{} `field:"optional" json:"labelFontConfiguration" yaml:"labelFontConfiguration"`
	// `CfnDashboard.DataLabelOptionsProperty.MeasureLabelVisibility`.
	MeasureLabelVisibility *string `field:"optional" json:"measureLabelVisibility" yaml:"measureLabelVisibility"`
	// `CfnDashboard.DataLabelOptionsProperty.Overlap`.
	Overlap *string `field:"optional" json:"overlap" yaml:"overlap"`
	// `CfnDashboard.DataLabelOptionsProperty.Position`.
	Position *string `field:"optional" json:"position" yaml:"position"`
	// `CfnDashboard.DataLabelOptionsProperty.Visibility`.
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

