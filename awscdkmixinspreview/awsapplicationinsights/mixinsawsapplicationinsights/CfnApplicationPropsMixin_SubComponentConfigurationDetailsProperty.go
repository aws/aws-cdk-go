package mixinsawsapplicationinsights


// The `AWS::ApplicationInsights::Application SubComponentConfigurationDetails` property type specifies the configuration settings of the sub-components.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   subComponentConfigurationDetailsProperty := &SubComponentConfigurationDetailsProperty{
//   	AlarmMetrics: []interface{}{
//   		&AlarmMetricProperty{
//   			AlarmMetricName: jsii.String("alarmMetricName"),
//   		},
//   	},
//   	Logs: []interface{}{
//   		&LogProperty{
//   			Encoding: jsii.String("encoding"),
//   			LogGroupName: jsii.String("logGroupName"),
//   			LogPath: jsii.String("logPath"),
//   			LogType: jsii.String("logType"),
//   			PatternSet: jsii.String("patternSet"),
//   		},
//   	},
//   	Processes: []interface{}{
//   		&ProcessProperty{
//   			AlarmMetrics: []interface{}{
//   				&AlarmMetricProperty{
//   					AlarmMetricName: jsii.String("alarmMetricName"),
//   				},
//   			},
//   			ProcessName: jsii.String("processName"),
//   		},
//   	},
//   	WindowsEvents: []interface{}{
//   		&WindowsEventProperty{
//   			EventLevels: []*string{
//   				jsii.String("eventLevels"),
//   			},
//   			EventName: jsii.String("eventName"),
//   			LogGroupName: jsii.String("logGroupName"),
//   			PatternSet: jsii.String("patternSet"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-subcomponentconfigurationdetails.html
//
type CfnApplicationPropsMixin_SubComponentConfigurationDetailsProperty struct {
	// A list of metrics to monitor for the component.
	//
	// All component types can use `AlarmMetrics` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-subcomponentconfigurationdetails.html#cfn-applicationinsights-application-subcomponentconfigurationdetails-alarmmetrics
	//
	AlarmMetrics interface{} `field:"optional" json:"alarmMetrics" yaml:"alarmMetrics"`
	// A list of logs to monitor for the component.
	//
	// Only Amazon EC2 instances can use `Logs` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-subcomponentconfigurationdetails.html#cfn-applicationinsights-application-subcomponentconfigurationdetails-logs
	//
	Logs interface{} `field:"optional" json:"logs" yaml:"logs"`
	// A list of processes to monitor for the component.
	//
	// Only Windows EC2 instances can have a processes section.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-subcomponentconfigurationdetails.html#cfn-applicationinsights-application-subcomponentconfigurationdetails-processes
	//
	Processes interface{} `field:"optional" json:"processes" yaml:"processes"`
	// A list of Windows Events to monitor for the component.
	//
	// Only Amazon EC2 instances running on Windows can use `WindowsEvents` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-subcomponentconfigurationdetails.html#cfn-applicationinsights-application-subcomponentconfigurationdetails-windowsevents
	//
	WindowsEvents interface{} `field:"optional" json:"windowsEvents" yaml:"windowsEvents"`
}

