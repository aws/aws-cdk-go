package awsapplicationinsights


// The `AWS::ApplicationInsights::Application SubComponentConfigurationDetails` property type specifies the configuration settings of the sub-components.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   subComponentConfigurationDetailsProperty := &SubComponentConfigurationDetailsProperty{
//   	AlarmMetrics: []interface{}{
//   		&AlarmMetricProperty{
//   			AlarmMetricName: jsii.String("alarmMetricName"),
//   		},
//   	},
//   	Logs: []interface{}{
//   		&LogProperty{
//   			LogType: jsii.String("logType"),
//
//   			// the properties below are optional
//   			Encoding: jsii.String("encoding"),
//   			LogGroupName: jsii.String("logGroupName"),
//   			LogPath: jsii.String("logPath"),
//   			PatternSet: jsii.String("patternSet"),
//   		},
//   	},
//   	WindowsEvents: []interface{}{
//   		&WindowsEventProperty{
//   			EventLevels: []*string{
//   				jsii.String("eventLevels"),
//   			},
//   			EventName: jsii.String("eventName"),
//   			LogGroupName: jsii.String("logGroupName"),
//
//   			// the properties below are optional
//   			PatternSet: jsii.String("patternSet"),
//   		},
//   	},
//   }
//
type CfnApplication_SubComponentConfigurationDetailsProperty struct {
	// A list of metrics to monitor for the component.
	//
	// All component types can use `AlarmMetrics` .
	AlarmMetrics interface{} `field:"optional" json:"alarmMetrics" yaml:"alarmMetrics"`
	// A list of logs to monitor for the component.
	//
	// Only Amazon EC2 instances can use `Logs` .
	Logs interface{} `field:"optional" json:"logs" yaml:"logs"`
	// A list of Windows Events to monitor for the component.
	//
	// Only Amazon EC2 instances running on Windows can use `WindowsEvents` .
	WindowsEvents interface{} `field:"optional" json:"windowsEvents" yaml:"windowsEvents"`
}

