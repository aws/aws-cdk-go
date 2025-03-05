package awsquicksight


// Options that determine the layout and display options of a chart's small multiples.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   smallMultiplesOptionsProperty := &SmallMultiplesOptionsProperty{
//   	MaxVisibleColumns: jsii.Number(123),
//   	MaxVisibleRows: jsii.Number(123),
//   	PanelConfiguration: &PanelConfigurationProperty{
//   		BackgroundColor: jsii.String("backgroundColor"),
//   		BackgroundVisibility: jsii.String("backgroundVisibility"),
//   		BorderColor: jsii.String("borderColor"),
//   		BorderStyle: jsii.String("borderStyle"),
//   		BorderThickness: jsii.String("borderThickness"),
//   		BorderVisibility: jsii.String("borderVisibility"),
//   		GutterSpacing: jsii.String("gutterSpacing"),
//   		GutterVisibility: jsii.String("gutterVisibility"),
//   		Title: &PanelTitleOptionsProperty{
//   			FontConfiguration: &FontConfigurationProperty{
//   				FontColor: jsii.String("fontColor"),
//   				FontDecoration: jsii.String("fontDecoration"),
//   				FontSize: &FontSizeProperty{
//   					Absolute: jsii.String("absolute"),
//   					Relative: jsii.String("relative"),
//   				},
//   				FontStyle: jsii.String("fontStyle"),
//   				FontWeight: &FontWeightProperty{
//   					Name: jsii.String("name"),
//   				},
//   			},
//   			HorizontalTextAlignment: jsii.String("horizontalTextAlignment"),
//   			Visibility: jsii.String("visibility"),
//   		},
//   	},
//   	XAxis: &SmallMultiplesAxisPropertiesProperty{
//   		Placement: jsii.String("placement"),
//   		Scale: jsii.String("scale"),
//   	},
//   	YAxis: &SmallMultiplesAxisPropertiesProperty{
//   		Placement: jsii.String("placement"),
//   		Scale: jsii.String("scale"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-smallmultiplesoptions.html
//
type CfnAnalysis_SmallMultiplesOptionsProperty struct {
	// Sets the maximum number of visible columns to display in the grid of small multiples panels.
	//
	// The default is `Auto` , which automatically adjusts the columns in the grid to fit the overall layout and size of the given chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-smallmultiplesoptions.html#cfn-quicksight-analysis-smallmultiplesoptions-maxvisiblecolumns
	//
	MaxVisibleColumns *float64 `field:"optional" json:"maxVisibleColumns" yaml:"maxVisibleColumns"`
	// Sets the maximum number of visible rows to display in the grid of small multiples panels.
	//
	// The default value is `Auto` , which automatically adjusts the rows in the grid to fit the overall layout and size of the given chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-smallmultiplesoptions.html#cfn-quicksight-analysis-smallmultiplesoptions-maxvisiblerows
	//
	MaxVisibleRows *float64 `field:"optional" json:"maxVisibleRows" yaml:"maxVisibleRows"`
	// Configures the display options for each small multiples panel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-smallmultiplesoptions.html#cfn-quicksight-analysis-smallmultiplesoptions-panelconfiguration
	//
	PanelConfiguration interface{} `field:"optional" json:"panelConfiguration" yaml:"panelConfiguration"`
	// The properties of a small multiples X axis.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-smallmultiplesoptions.html#cfn-quicksight-analysis-smallmultiplesoptions-xaxis
	//
	XAxis interface{} `field:"optional" json:"xAxis" yaml:"xAxis"`
	// The properties of a small multiples Y axis.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-smallmultiplesoptions.html#cfn-quicksight-analysis-smallmultiplesoptions-yaxis
	//
	YAxis interface{} `field:"optional" json:"yAxis" yaml:"yAxis"`
}

