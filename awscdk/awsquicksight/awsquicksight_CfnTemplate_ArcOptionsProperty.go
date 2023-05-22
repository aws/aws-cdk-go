package awsquicksight


// The options that determine the arc thickness of a `GaugeChartVisual` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   arcOptionsProperty := &ArcOptionsProperty{
//   	ArcThickness: jsii.String("arcThickness"),
//   }
//
type CfnTemplate_ArcOptionsProperty struct {
	// The arc thickness of a `GaugeChartVisual` .
	ArcThickness *string `field:"optional" json:"arcThickness" yaml:"arcThickness"`
}

