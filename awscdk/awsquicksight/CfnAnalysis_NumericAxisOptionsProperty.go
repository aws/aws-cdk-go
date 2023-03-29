package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var dataDriven interface{}
//
//   numericAxisOptionsProperty := &NumericAxisOptionsProperty{
//   	Range: &AxisDisplayRangeProperty{
//   		DataDriven: dataDriven,
//   		MinMax: &AxisDisplayMinMaxRangeProperty{
//   			Maximum: jsii.Number(123),
//   			Minimum: jsii.Number(123),
//   		},
//   	},
//   	Scale: &AxisScaleProperty{
//   		Linear: &AxisLinearScaleProperty{
//   			StepCount: jsii.Number(123),
//   			StepSize: jsii.Number(123),
//   		},
//   		Logarithmic: &AxisLogarithmicScaleProperty{
//   			Base: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnAnalysis_NumericAxisOptionsProperty struct {
	// `CfnAnalysis.NumericAxisOptionsProperty.Range`.
	Range interface{} `field:"optional" json:"range" yaml:"range"`
	// `CfnAnalysis.NumericAxisOptionsProperty.Scale`.
	Scale interface{} `field:"optional" json:"scale" yaml:"scale"`
}

