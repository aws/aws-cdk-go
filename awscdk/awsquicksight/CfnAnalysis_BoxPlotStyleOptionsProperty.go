package awsquicksight


// The style options of the box plot.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   boxPlotStyleOptionsProperty := &BoxPlotStyleOptionsProperty{
//   	FillStyle: jsii.String("fillStyle"),
//   }
//
type CfnAnalysis_BoxPlotStyleOptionsProperty struct {
	// The fill styles (solid, transparent) of the box plot.
	FillStyle *string `field:"optional" json:"fillStyle" yaml:"fillStyle"`
}

