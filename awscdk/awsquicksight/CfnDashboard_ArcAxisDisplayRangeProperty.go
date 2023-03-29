package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   arcAxisDisplayRangeProperty := &ArcAxisDisplayRangeProperty{
//   	Max: jsii.Number(123),
//   	Min: jsii.Number(123),
//   }
//
type CfnDashboard_ArcAxisDisplayRangeProperty struct {
	// `CfnDashboard.ArcAxisDisplayRangeProperty.Max`.
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// `CfnDashboard.ArcAxisDisplayRangeProperty.Min`.
	Min *float64 `field:"optional" json:"min" yaml:"min"`
}

