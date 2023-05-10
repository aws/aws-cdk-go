package awsquicksight


// The arc configuration of a `GaugeChartVisual` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   arcConfigurationProperty := &ArcConfigurationProperty{
//   	ArcAngle: jsii.Number(123),
//   	ArcThickness: jsii.String("arcThickness"),
//   }
//
type CfnTemplate_ArcConfigurationProperty struct {
	// The option that determines the arc angle of a `GaugeChartVisual` .
	ArcAngle *float64 `field:"optional" json:"arcAngle" yaml:"arcAngle"`
	// The options that determine the arc thickness of a `GaugeChartVisual` .
	ArcThickness *string `field:"optional" json:"arcThickness" yaml:"arcThickness"`
}

