package mixinsawsquicksight


// The symbol style for null data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   geospatialNullSymbolStyleProperty := &GeospatialNullSymbolStyleProperty{
//   	FillColor: jsii.String("fillColor"),
//   	StrokeColor: jsii.String("strokeColor"),
//   	StrokeWidth: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatialnullsymbolstyle.html
//
type CfnDashboardPropsMixin_GeospatialNullSymbolStyleProperty struct {
	// The color and opacity values for the fill color.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatialnullsymbolstyle.html#cfn-quicksight-dashboard-geospatialnullsymbolstyle-fillcolor
	//
	FillColor *string `field:"optional" json:"fillColor" yaml:"fillColor"`
	// The color and opacity values for the stroke color.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatialnullsymbolstyle.html#cfn-quicksight-dashboard-geospatialnullsymbolstyle-strokecolor
	//
	StrokeColor *string `field:"optional" json:"strokeColor" yaml:"strokeColor"`
	// The width of the border stroke.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatialnullsymbolstyle.html#cfn-quicksight-dashboard-geospatialnullsymbolstyle-strokewidth
	//
	StrokeWidth *float64 `field:"optional" json:"strokeWidth" yaml:"strokeWidth"`
}

