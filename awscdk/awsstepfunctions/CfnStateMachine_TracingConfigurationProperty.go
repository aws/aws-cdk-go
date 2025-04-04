package awsstepfunctions


// Selects whether or not the state machine's AWS X-Ray tracing is enabled.
//
// To configure your state machine to send trace data to X-Ray, set `Enabled` to `true` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tracingConfigurationProperty := &TracingConfigurationProperty{
//   	Enabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachine-tracingconfiguration.html
//
type CfnStateMachine_TracingConfigurationProperty struct {
	// When set to `true` , X-Ray tracing is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachine-tracingconfiguration.html#cfn-stepfunctions-statemachine-tracingconfiguration-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

