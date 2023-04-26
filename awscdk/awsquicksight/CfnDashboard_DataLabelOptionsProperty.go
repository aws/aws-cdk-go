package awsquicksight


// The options that determine the presentation of the data labels.
//
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
	// Determines the visibility of the category field labels.
	CategoryLabelVisibility *string `field:"optional" json:"categoryLabelVisibility" yaml:"categoryLabelVisibility"`
	// The option that determines the data label type.
	DataLabelTypes interface{} `field:"optional" json:"dataLabelTypes" yaml:"dataLabelTypes"`
	// Determines the color of the data labels.
	LabelColor *string `field:"optional" json:"labelColor" yaml:"labelColor"`
	// Determines the content of the data labels.
	LabelContent *string `field:"optional" json:"labelContent" yaml:"labelContent"`
	// Determines the font configuration of the data labels.
	LabelFontConfiguration interface{} `field:"optional" json:"labelFontConfiguration" yaml:"labelFontConfiguration"`
	// Determines the visibility of the measure field labels.
	MeasureLabelVisibility *string `field:"optional" json:"measureLabelVisibility" yaml:"measureLabelVisibility"`
	// Determines whether overlap is enabled or disabled for the data labels.
	Overlap *string `field:"optional" json:"overlap" yaml:"overlap"`
	// Determines the position of the data labels.
	Position *string `field:"optional" json:"position" yaml:"position"`
	// Determines the visibility of the data labels.
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

