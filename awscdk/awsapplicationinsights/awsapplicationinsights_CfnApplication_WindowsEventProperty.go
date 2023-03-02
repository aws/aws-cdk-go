package awsapplicationinsights


// The `AWS::ApplicationInsights::Application WindowsEvent` property type specifies a Windows Event to monitor for the component.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   windowsEventProperty := &windowsEventProperty{
//   	eventLevels: []*string{
//   		jsii.String("eventLevels"),
//   	},
//   	eventName: jsii.String("eventName"),
//   	logGroupName: jsii.String("logGroupName"),
//
//   	// the properties below are optional
//   	patternSet: jsii.String("patternSet"),
//   }
//
type CfnApplication_WindowsEventProperty struct {
	// The levels of event to log.
	//
	// You must specify each level to log. Possible values include `INFORMATION` , `WARNING` , `ERROR` , `CRITICAL` , and `VERBOSE` . This field is required for each type of Windows Event to log.
	EventLevels *[]*string `field:"required" json:"eventLevels" yaml:"eventLevels"`
	// The type of Windows Events to log, equivalent to the Windows Event log channel name.
	//
	// For example, System, Security, CustomEventName, and so on. This field is required for each type of Windows event to log.
	EventName *string `field:"required" json:"eventName" yaml:"eventName"`
	// The CloudWatch log group name to be associated with the monitored log.
	LogGroupName *string `field:"required" json:"logGroupName" yaml:"logGroupName"`
	// The log pattern set.
	PatternSet *string `field:"optional" json:"patternSet" yaml:"patternSet"`
}

