package previewawsapplicationinsightsmixins


// A process to be monitored for the component.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   processProperty := &ProcessProperty{
//   	AlarmMetrics: []interface{}{
//   		&AlarmMetricProperty{
//   			AlarmMetricName: jsii.String("alarmMetricName"),
//   		},
//   	},
//   	ProcessName: jsii.String("processName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-process.html
//
type CfnApplicationPropsMixin_ProcessProperty struct {
	// A list of metrics to monitor for the component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-process.html#cfn-applicationinsights-application-process-alarmmetrics
	//
	AlarmMetrics interface{} `field:"optional" json:"alarmMetrics" yaml:"alarmMetrics"`
	// The name of the process to be monitored for the component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-process.html#cfn-applicationinsights-application-process-processname
	//
	ProcessName *string `field:"optional" json:"processName" yaml:"processName"`
}

