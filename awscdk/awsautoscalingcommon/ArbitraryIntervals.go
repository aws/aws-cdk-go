package awsautoscalingcommon


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   arbitraryIntervals := &ArbitraryIntervals{
//   	Absolute: jsii.Boolean(false),
//   	Intervals: []scalingInterval{
//   		&scalingInterval{
//   			Change: jsii.Number(123),
//
//   			// the properties below are optional
//   			Lower: jsii.Number(123),
//   			Upper: jsii.Number(123),
//   		},
//   	},
//   }
//
type ArbitraryIntervals struct {
	Absolute *bool `field:"required" json:"absolute" yaml:"absolute"`
	Intervals *[]*ScalingInterval `field:"required" json:"intervals" yaml:"intervals"`
}

