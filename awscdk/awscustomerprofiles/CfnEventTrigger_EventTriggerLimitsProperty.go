package awscustomerprofiles


// Defines limits controlling whether an event triggers the destination, based on ingestion latency and the number of invocations per profile over specific time periods.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventTriggerLimitsProperty := &EventTriggerLimitsProperty{
//   	EventExpiration: jsii.Number(123),
//   	Periods: []interface{}{
//   		&PeriodProperty{
//   			Unit: jsii.String("unit"),
//   			Value: jsii.Number(123),
//
//   			// the properties below are optional
//   			MaxInvocationsPerProfile: jsii.Number(123),
//   			Unlimited: jsii.Boolean(false),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-eventtrigger-eventtriggerlimits.html
//
type CfnEventTrigger_EventTriggerLimitsProperty struct {
	// Specifies that an event will only trigger the destination if it is processed within a certain latency period.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-eventtrigger-eventtriggerlimits.html#cfn-customerprofiles-eventtrigger-eventtriggerlimits-eventexpiration
	//
	EventExpiration *float64 `field:"optional" json:"eventExpiration" yaml:"eventExpiration"`
	// A list of time periods during which the limits apply.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-eventtrigger-eventtriggerlimits.html#cfn-customerprofiles-eventtrigger-eventtriggerlimits-periods
	//
	Periods interface{} `field:"optional" json:"periods" yaml:"periods"`
}

