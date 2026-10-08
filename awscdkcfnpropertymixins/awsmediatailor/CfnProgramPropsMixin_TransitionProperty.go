package awsmediatailor


// Program transition configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   transitionProperty := &TransitionProperty{
//   	DurationMillis: jsii.Number(123),
//   	RelativePosition: jsii.String("relativePosition"),
//   	RelativeProgram: jsii.String("relativeProgram"),
//   	ScheduledStartTimeMillis: jsii.Number(123),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-program-transition.html
//
type CfnProgramPropsMixin_TransitionProperty struct {
	// The duration of the live program in seconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-program-transition.html#cfn-mediatailor-program-transition-durationmillis
	//
	DurationMillis *float64 `field:"optional" json:"durationMillis" yaml:"durationMillis"`
	// The position where this program will be inserted relative to the RelativePosition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-program-transition.html#cfn-mediatailor-program-transition-relativeposition
	//
	RelativePosition *string `field:"optional" json:"relativePosition" yaml:"relativePosition"`
	// The name of the program that this program will be inserted next to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-program-transition.html#cfn-mediatailor-program-transition-relativeprogram
	//
	RelativeProgram *string `field:"optional" json:"relativeProgram" yaml:"relativeProgram"`
	// The date and time that the program is scheduled to start, in epoch milliseconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-program-transition.html#cfn-mediatailor-program-transition-scheduledstarttimemillis
	//
	ScheduledStartTimeMillis *float64 `field:"optional" json:"scheduledStartTimeMillis" yaml:"scheduledStartTimeMillis"`
	// Defines when the program plays in the schedule.
	//
	// You can set the value to ABSOLUTE or RELATIVE.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-program-transition.html#cfn-mediatailor-program-transition-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

