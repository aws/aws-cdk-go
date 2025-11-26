package previewawsquicksightmixins


// Determines the gradient color settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   gradientColorProperty := &GradientColorProperty{
//   	Stops: []interface{}{
//   		&GradientStopProperty{
//   			Color: jsii.String("color"),
//   			DataValue: jsii.Number(123),
//   			GradientOffset: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-gradientcolor.html
//
type CfnAnalysisPropsMixin_GradientColorProperty struct {
	// The list of gradient color stops.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-gradientcolor.html#cfn-quicksight-analysis-gradientcolor-stops
	//
	Stops interface{} `field:"optional" json:"stops" yaml:"stops"`
}

