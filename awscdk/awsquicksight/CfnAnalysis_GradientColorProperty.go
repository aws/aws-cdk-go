package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gradientColorProperty := &GradientColorProperty{
//   	Stops: []interface{}{
//   		&GradientStopProperty{
//   			GradientOffset: jsii.Number(123),
//
//   			// the properties below are optional
//   			Color: jsii.String("color"),
//   			DataValue: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnAnalysis_GradientColorProperty struct {
	// `CfnAnalysis.GradientColorProperty.Stops`.
	Stops interface{} `field:"optional" json:"stops" yaml:"stops"`
}

