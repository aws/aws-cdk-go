package awsapplicationinsights


// The `AWS::ApplicationInsights::Application SubComponentTypeConfiguration` property type specifies the sub-component configurations for a component.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   subComponentTypeConfigurationProperty := &SubComponentTypeConfigurationProperty{
//   	SubComponentConfigurationDetails: &SubComponentConfigurationDetailsProperty{
//   		AlarmMetrics: []interface{}{
//   			&AlarmMetricProperty{
//   				AlarmMetricName: jsii.String("alarmMetricName"),
//   			},
//   		},
//   		Logs: []interface{}{
//   			&LogProperty{
//   				LogType: jsii.String("logType"),
//
//   				// the properties below are optional
//   				Encoding: jsii.String("encoding"),
//   				LogGroupName: jsii.String("logGroupName"),
//   				LogPath: jsii.String("logPath"),
//   				PatternSet: jsii.String("patternSet"),
//   			},
//   		},
//   		Processes: []interface{}{
//   			&ProcessProperty{
//   				AlarmMetrics: []interface{}{
//   					&AlarmMetricProperty{
//   						AlarmMetricName: jsii.String("alarmMetricName"),
//   					},
//   				},
//   				ProcessName: jsii.String("processName"),
//   			},
//   		},
//   		WindowsEvents: []interface{}{
//   			&WindowsEventProperty{
//   				EventLevels: []*string{
//   					jsii.String("eventLevels"),
//   				},
//   				EventName: jsii.String("eventName"),
//   				LogGroupName: jsii.String("logGroupName"),
//
//   				// the properties below are optional
//   				PatternSet: jsii.String("patternSet"),
//   			},
//   		},
//   	},
//   	SubComponentType: jsii.String("subComponentType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-subcomponenttypeconfiguration.html
//
type CfnApplication_SubComponentTypeConfigurationProperty struct {
	// The configuration settings of the sub-components.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-subcomponenttypeconfiguration.html#cfn-applicationinsights-application-subcomponenttypeconfiguration-subcomponentconfigurationdetails
	//
	SubComponentConfigurationDetails interface{} `field:"required" json:"subComponentConfigurationDetails" yaml:"subComponentConfigurationDetails"`
	// The sub-component type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-subcomponenttypeconfiguration.html#cfn-applicationinsights-application-subcomponenttypeconfiguration-subcomponenttype
	//
	SubComponentType *string `field:"required" json:"subComponentType" yaml:"subComponentType"`
}

