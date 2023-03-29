package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   axisTickLabelOptionsProperty := &AxisTickLabelOptionsProperty{
//   	LabelOptions: &LabelOptionsProperty{
//   		CustomLabel: jsii.String("customLabel"),
//   		FontConfiguration: &FontConfigurationProperty{
//   			FontColor: jsii.String("fontColor"),
//   			FontDecoration: jsii.String("fontDecoration"),
//   			FontSize: &FontSizeProperty{
//   				Relative: jsii.String("relative"),
//   			},
//   			FontStyle: jsii.String("fontStyle"),
//   			FontWeight: &FontWeightProperty{
//   				Name: jsii.String("name"),
//   			},
//   		},
//   		Visibility: jsii.String("visibility"),
//   	},
//   	RotationAngle: jsii.Number(123),
//   }
//
type CfnAnalysis_AxisTickLabelOptionsProperty struct {
	// `CfnAnalysis.AxisTickLabelOptionsProperty.LabelOptions`.
	LabelOptions interface{} `field:"optional" json:"labelOptions" yaml:"labelOptions"`
	// `CfnAnalysis.AxisTickLabelOptionsProperty.RotationAngle`.
	RotationAngle *float64 `field:"optional" json:"rotationAngle" yaml:"rotationAngle"`
}

