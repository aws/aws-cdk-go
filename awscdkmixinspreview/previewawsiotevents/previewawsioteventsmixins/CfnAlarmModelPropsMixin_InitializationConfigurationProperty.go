package previewawsioteventsmixins


// Specifies the default alarm state.
//
// The configuration applies to all alarms that were created based on this alarm model.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   initializationConfigurationProperty := &InitializationConfigurationProperty{
//   	DisabledOnInitialization: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-alarmmodel-initializationconfiguration.html
//
type CfnAlarmModelPropsMixin_InitializationConfigurationProperty struct {
	// The value must be `TRUE` or `FALSE` .
	//
	// If `FALSE` , all alarm instances created based on the alarm model are activated. The default value is `TRUE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-alarmmodel-initializationconfiguration.html#cfn-iotevents-alarmmodel-initializationconfiguration-disabledoninitialization
	//
	DisabledOnInitialization interface{} `field:"optional" json:"disabledOnInitialization" yaml:"disabledOnInitialization"`
}

