package awsmediatailor


// Schedule configuration parameters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scheduleConfigurationProperty := &ScheduleConfigurationProperty{
//   	Transition: &TransitionProperty{
//   		RelativePosition: jsii.String("relativePosition"),
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
//   		DurationMillis: jsii.Number(123),
//   		RelativeProgram: jsii.String("relativeProgram"),
//   		ScheduledStartTimeMillis: jsii.Number(123),
//   	},
//
//   	// the properties below are optional
//   	ClipRange: &ClipRangeProperty{
//   		EndOffsetMillis: jsii.Number(123),
//   		StartOffsetMillis: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-program-scheduleconfiguration.html
//
type CfnProgram_ScheduleConfigurationProperty struct {
	// Program transition configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-program-scheduleconfiguration.html#cfn-mediatailor-program-scheduleconfiguration-transition
	//
	Transition interface{} `field:"required" json:"transition" yaml:"transition"`
	// Clip range configuration for the VOD source associated with the program.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-program-scheduleconfiguration.html#cfn-mediatailor-program-scheduleconfiguration-cliprange
	//
	ClipRange interface{} `field:"optional" json:"clipRange" yaml:"clipRange"`
}

