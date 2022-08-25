package awsautoscalingcommon


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   arbitraryIntervals := &arbitraryIntervals{
//   	absolute: jsii.Boolean(false),
//   	intervals: []scalingInterval{
//   		&scalingInterval{
//   			change: jsii.Number(123),
//
//   			// the properties below are optional
//   			lower: jsii.Number(123),
//   			upper: jsii.Number(123),
//   		},
//   	},
//   }
//
type ArbitraryIntervals struct {
	Absolute *bool `field:"required" json:"absolute" yaml:"absolute"`
	Intervals *[]*ScalingInterval `field:"required" json:"intervals" yaml:"intervals"`
}

