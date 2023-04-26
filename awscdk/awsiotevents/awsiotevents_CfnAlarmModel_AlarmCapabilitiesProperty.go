package awsiotevents


// Contains the configuration information of alarm state changes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   alarmCapabilitiesProperty := &AlarmCapabilitiesProperty{
//   	AcknowledgeFlow: &AcknowledgeFlowProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	InitializationConfiguration: &InitializationConfigurationProperty{
//   		DisabledOnInitialization: jsii.Boolean(false),
//   	},
//   }
//
type CfnAlarmModel_AlarmCapabilitiesProperty struct {
	// Specifies whether to get notified for alarm state changes.
	AcknowledgeFlow interface{} `field:"optional" json:"acknowledgeFlow" yaml:"acknowledgeFlow"`
	// Specifies the default alarm state.
	//
	// The configuration applies to all alarms that were created based on this alarm model.
	InitializationConfiguration interface{} `field:"optional" json:"initializationConfiguration" yaml:"initializationConfiguration"`
}

