package awsapplicationinsights


// The `AWS::ApplicationInsights::Application SubComponentTypeConfiguration` property type specifies the sub-component configurations for a component.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   subComponentTypeConfigurationProperty := &subComponentTypeConfigurationProperty{
//   	subComponentConfigurationDetails: &subComponentConfigurationDetailsProperty{
//   		alarmMetrics: []interface{}{
//   			&alarmMetricProperty{
//   				alarmMetricName: jsii.String("alarmMetricName"),
//   			},
//   		},
//   		logs: []interface{}{
//   			&logProperty{
//   				logType: jsii.String("logType"),
//
//   				// the properties below are optional
//   				encoding: jsii.String("encoding"),
//   				logGroupName: jsii.String("logGroupName"),
//   				logPath: jsii.String("logPath"),
//   				patternSet: jsii.String("patternSet"),
//   			},
//   		},
//   		windowsEvents: []interface{}{
//   			&windowsEventProperty{
//   				eventLevels: []*string{
//   					jsii.String("eventLevels"),
//   				},
//   				eventName: jsii.String("eventName"),
//   				logGroupName: jsii.String("logGroupName"),
//
//   				// the properties below are optional
//   				patternSet: jsii.String("patternSet"),
//   			},
//   		},
//   	},
//   	subComponentType: jsii.String("subComponentType"),
//   }
//
type CfnApplication_SubComponentTypeConfigurationProperty struct {
	// The configuration settings of the sub-components.
	SubComponentConfigurationDetails interface{} `field:"required" json:"subComponentConfigurationDetails" yaml:"subComponentConfigurationDetails"`
	// The sub-component type.
	SubComponentType *string `field:"required" json:"subComponentType" yaml:"subComponentType"`
}

