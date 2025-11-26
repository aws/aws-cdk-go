package previewawsioteventsmixins


// Contains the configuration information of alarm state changes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-alarmmodel-alarmcapabilities.html
//
type CfnAlarmModelPropsMixin_AlarmCapabilitiesProperty struct {
	// Specifies whether to get notified for alarm state changes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-alarmmodel-alarmcapabilities.html#cfn-iotevents-alarmmodel-alarmcapabilities-acknowledgeflow
	//
	AcknowledgeFlow interface{} `field:"optional" json:"acknowledgeFlow" yaml:"acknowledgeFlow"`
	// Specifies the default alarm state.
	//
	// The configuration applies to all alarms that were created based on this alarm model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-alarmmodel-alarmcapabilities.html#cfn-iotevents-alarmmodel-alarmcapabilities-initializationconfiguration
	//
	InitializationConfiguration interface{} `field:"optional" json:"initializationConfiguration" yaml:"initializationConfiguration"`
}

