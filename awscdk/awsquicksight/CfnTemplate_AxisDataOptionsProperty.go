package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var dataDriven interface{}
//
//   axisDataOptionsProperty := &AxisDataOptionsProperty{
//   	DateAxisOptions: &DateAxisOptionsProperty{
//   		MissingDateVisibility: jsii.String("missingDateVisibility"),
//   	},
//   	NumericAxisOptions: &NumericAxisOptionsProperty{
//   		Range: &AxisDisplayRangeProperty{
//   			DataDriven: dataDriven,
//   			MinMax: &AxisDisplayMinMaxRangeProperty{
//   				Maximum: jsii.Number(123),
//   				Minimum: jsii.Number(123),
//   			},
//   		},
//   		Scale: &AxisScaleProperty{
//   			Linear: &AxisLinearScaleProperty{
//   				StepCount: jsii.Number(123),
//   				StepSize: jsii.Number(123),
//   			},
//   			Logarithmic: &AxisLogarithmicScaleProperty{
//   				Base: jsii.Number(123),
//   			},
//   		},
//   	},
//   }
//
type CfnTemplate_AxisDataOptionsProperty struct {
	// `CfnTemplate.AxisDataOptionsProperty.DateAxisOptions`.
	DateAxisOptions interface{} `field:"optional" json:"dateAxisOptions" yaml:"dateAxisOptions"`
	// `CfnTemplate.AxisDataOptionsProperty.NumericAxisOptions`.
	NumericAxisOptions interface{} `field:"optional" json:"numericAxisOptions" yaml:"numericAxisOptions"`
}

