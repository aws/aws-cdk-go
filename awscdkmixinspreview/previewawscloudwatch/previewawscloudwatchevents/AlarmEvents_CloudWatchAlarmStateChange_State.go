package previewawscloudwatchevents


// Type definition for State.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   state := &State{
//   	ActionsSuppressedBy: []*string{
//   		jsii.String("actionsSuppressedBy"),
//   	},
//   	ActionsSuppressedReason: []*string{
//   		jsii.String("actionsSuppressedReason"),
//   	},
//   	EvaluationState: []*string{
//   		jsii.String("evaluationState"),
//   	},
//   	Reason: []*string{
//   		jsii.String("reason"),
//   	},
//   	ReasonData: []*string{
//   		jsii.String("reasonData"),
//   	},
//   	Timestamp: []*string{
//   		jsii.String("timestamp"),
//   	},
//   	Value: []*string{
//   		jsii.String("value"),
//   	},
//   }
//
// Experimental.
type AlarmEvents_CloudWatchAlarmStateChange_State struct {
	// actionsSuppressedBy property.
	//
	// Specify an array of string values to match this event if the actual value of actionsSuppressedBy is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ActionsSuppressedBy *[]*string `field:"optional" json:"actionsSuppressedBy" yaml:"actionsSuppressedBy"`
	// actionsSuppressedReason property.
	//
	// Specify an array of string values to match this event if the actual value of actionsSuppressedReason is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ActionsSuppressedReason *[]*string `field:"optional" json:"actionsSuppressedReason" yaml:"actionsSuppressedReason"`
	// evaluationState property.
	//
	// Specify an array of string values to match this event if the actual value of evaluationState is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EvaluationState *[]*string `field:"optional" json:"evaluationState" yaml:"evaluationState"`
	// reason property.
	//
	// Specify an array of string values to match this event if the actual value of reason is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Reason *[]*string `field:"optional" json:"reason" yaml:"reason"`
	// reasonData property.
	//
	// Specify an array of string values to match this event if the actual value of reasonData is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ReasonData *[]*string `field:"optional" json:"reasonData" yaml:"reasonData"`
	// timestamp property.
	//
	// Specify an array of string values to match this event if the actual value of timestamp is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Timestamp *[]*string `field:"optional" json:"timestamp" yaml:"timestamp"`
	// value property.
	//
	// Specify an array of string values to match this event if the actual value of value is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Value *[]*string `field:"optional" json:"value" yaml:"value"`
}

