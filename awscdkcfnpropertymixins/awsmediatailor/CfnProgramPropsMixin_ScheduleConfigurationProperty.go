package awsmediatailor


// Schedule configuration parameters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   scheduleConfigurationProperty := &ScheduleConfigurationProperty{
//   	ClipRange: &ClipRangeProperty{
//   		EndOffsetMillis: jsii.Number(123),
//   		StartOffsetMillis: jsii.Number(123),
//   	},
//   	Transition: &TransitionProperty{
//   		DurationMillis: jsii.Number(123),
//   		RelativePosition: jsii.String("relativePosition"),
//   		RelativeProgram: jsii.String("relativeProgram"),
//   		ScheduledStartTimeMillis: jsii.Number(123),
//   		Type: jsii.String("type"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-program-scheduleconfiguration.html
//
type CfnProgramPropsMixin_ScheduleConfigurationProperty struct {
	// Clip range configuration for the VOD source associated with the program.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-program-scheduleconfiguration.html#cfn-mediatailor-program-scheduleconfiguration-cliprange
	//
	ClipRange interface{} `field:"optional" json:"clipRange" yaml:"clipRange"`
	// Program transition configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-program-scheduleconfiguration.html#cfn-mediatailor-program-scheduleconfiguration-transition
	//
	Transition interface{} `field:"optional" json:"transition" yaml:"transition"`
}

