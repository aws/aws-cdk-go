package awsiotevents


// Specifies the default alarm state.
//
// The configuration applies to all alarms that were created based on this alarm model.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   initializationConfigurationProperty := &initializationConfigurationProperty{
//   	disabledOnInitialization: jsii.Boolean(false),
//   }
//
type CfnAlarmModel_InitializationConfigurationProperty struct {
	// The value must be `TRUE` or `FALSE` .
	//
	// If `FALSE` , all alarm instances created based on the alarm model are activated. The default value is `TRUE` .
	DisabledOnInitialization interface{} `field:"required" json:"disabledOnInitialization" yaml:"disabledOnInitialization"`
}

