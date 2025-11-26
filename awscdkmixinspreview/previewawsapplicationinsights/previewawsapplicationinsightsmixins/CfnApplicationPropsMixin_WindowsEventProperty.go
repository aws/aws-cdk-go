package previewawsapplicationinsightsmixins


// The `AWS::ApplicationInsights::Application WindowsEvent` property type specifies a Windows Event to monitor for the component.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   windowsEventProperty := &WindowsEventProperty{
//   	EventLevels: []*string{
//   		jsii.String("eventLevels"),
//   	},
//   	EventName: jsii.String("eventName"),
//   	LogGroupName: jsii.String("logGroupName"),
//   	PatternSet: jsii.String("patternSet"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-windowsevent.html
//
type CfnApplicationPropsMixin_WindowsEventProperty struct {
	// The levels of event to log.
	//
	// You must specify each level to log. Possible values include `INFORMATION` , `WARNING` , `ERROR` , `CRITICAL` , and `VERBOSE` . This field is required for each type of Windows Event to log.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-windowsevent.html#cfn-applicationinsights-application-windowsevent-eventlevels
	//
	EventLevels *[]*string `field:"optional" json:"eventLevels" yaml:"eventLevels"`
	// The type of Windows Events to log, equivalent to the Windows Event log channel name.
	//
	// For example, System, Security, CustomEventName, and so on. This field is required for each type of Windows event to log.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-windowsevent.html#cfn-applicationinsights-application-windowsevent-eventname
	//
	EventName *string `field:"optional" json:"eventName" yaml:"eventName"`
	// The CloudWatch log group name to be associated with the monitored log.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-windowsevent.html#cfn-applicationinsights-application-windowsevent-loggroupname
	//
	LogGroupName *string `field:"optional" json:"logGroupName" yaml:"logGroupName"`
	// The log pattern set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-windowsevent.html#cfn-applicationinsights-application-windowsevent-patternset
	//
	PatternSet *string `field:"optional" json:"patternSet" yaml:"patternSet"`
}

