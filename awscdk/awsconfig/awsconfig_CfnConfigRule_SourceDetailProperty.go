package awsconfig


// Provides the source and the message types that trigger AWS Config to evaluate your AWS resources against a rule.
//
// It also provides the frequency with which you want AWS Config to run evaluations for the rule if the trigger type is periodic. You can specify the parameter values for `SourceDetail` only for custom rules.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceDetailProperty := &sourceDetailProperty{
//   	eventSource: jsii.String("eventSource"),
//   	messageType: jsii.String("messageType"),
//
//   	// the properties below are optional
//   	maximumExecutionFrequency: jsii.String("maximumExecutionFrequency"),
//   }
//
type CfnConfigRule_SourceDetailProperty struct {
	// The source of the event, such as an AWS service, that triggers AWS Config to evaluate your AWS resources.
	EventSource *string `field:"required" json:"eventSource" yaml:"eventSource"`
	// The type of notification that triggers AWS Config to run an evaluation for a rule.
	//
	// You can specify the following notification types:
	//
	// - `ConfigurationItemChangeNotification` - Triggers an evaluation when AWS Config delivers a configuration item as a result of a resource change.
	// - `OversizedConfigurationItemChangeNotification` - Triggers an evaluation when AWS Config delivers an oversized configuration item. AWS Config may generate this notification type when a resource changes and the notification exceeds the maximum size allowed by Amazon SNS.
	// - `ScheduledNotification` - Triggers a periodic evaluation at the frequency specified for `MaximumExecutionFrequency` .
	// - `ConfigurationSnapshotDeliveryCompleted` - Triggers a periodic evaluation when AWS Config delivers a configuration snapshot.
	//
	// If you want your custom rule to be triggered by configuration changes, specify two SourceDetail objects, one for `ConfigurationItemChangeNotification` and one for `OversizedConfigurationItemChangeNotification` .
	MessageType *string `field:"required" json:"messageType" yaml:"messageType"`
	// The frequency at which you want AWS Config to run evaluations for a custom rule with a periodic trigger.
	//
	// If you specify a value for `MaximumExecutionFrequency` , then `MessageType` must use the `ScheduledNotification` value.
	//
	// > By default, rules with a periodic trigger are evaluated every 24 hours. To change the frequency, specify a valid value for the `MaximumExecutionFrequency` parameter.
	// >
	// > Based on the valid value you choose, AWS Config runs evaluations once for each valid value. For example, if you choose `Three_Hours` , AWS Config runs evaluations once every three hours. In this case, `Three_Hours` is the frequency of this rule.
	MaximumExecutionFrequency *string `field:"optional" json:"maximumExecutionFrequency" yaml:"maximumExecutionFrequency"`
}

