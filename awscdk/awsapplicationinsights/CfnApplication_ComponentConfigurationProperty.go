package awsapplicationinsights


// The `AWS::ApplicationInsights::Application ComponentConfiguration` property type defines the configuration settings of the component.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   componentConfigurationProperty := &ComponentConfigurationProperty{
//   	ConfigurationDetails: &ConfigurationDetailsProperty{
//   		AlarmMetrics: []interface{}{
//   			&AlarmMetricProperty{
//   				AlarmMetricName: jsii.String("alarmMetricName"),
//   			},
//   		},
//   		Alarms: []interface{}{
//   			&AlarmProperty{
//   				AlarmName: jsii.String("alarmName"),
//
//   				// the properties below are optional
//   				Severity: jsii.String("severity"),
//   			},
//   		},
//   		HaClusterPrometheusExporter: &HAClusterPrometheusExporterProperty{
//   			PrometheusPort: jsii.String("prometheusPort"),
//   		},
//   		HanaPrometheusExporter: &HANAPrometheusExporterProperty{
//   			AgreeToInstallHanadbClient: jsii.Boolean(false),
//   			HanaPort: jsii.String("hanaPort"),
//   			HanaSecretName: jsii.String("hanaSecretName"),
//   			Hanasid: jsii.String("hanasid"),
//
//   			// the properties below are optional
//   			PrometheusPort: jsii.String("prometheusPort"),
//   		},
//   		JmxPrometheusExporter: &JMXPrometheusExporterProperty{
//   			HostPort: jsii.String("hostPort"),
//   			Jmxurl: jsii.String("jmxurl"),
//   			PrometheusPort: jsii.String("prometheusPort"),
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
//   		NetWeaverPrometheusExporter: &NetWeaverPrometheusExporterProperty{
//   			InstanceNumbers: []*string{
//   				jsii.String("instanceNumbers"),
//   			},
//   			Sapsid: jsii.String("sapsid"),
//
//   			// the properties below are optional
//   			PrometheusPort: jsii.String("prometheusPort"),
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
//   		SqlServerPrometheusExporter: &SQLServerPrometheusExporterProperty{
//   			PrometheusPort: jsii.String("prometheusPort"),
//   			SqlSecretName: jsii.String("sqlSecretName"),
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
//   	SubComponentTypeConfigurations: []interface{}{
//   		&SubComponentTypeConfigurationProperty{
//   			SubComponentConfigurationDetails: &SubComponentConfigurationDetailsProperty{
//   				AlarmMetrics: []interface{}{
//   					&AlarmMetricProperty{
//   						AlarmMetricName: jsii.String("alarmMetricName"),
//   					},
//   				},
//   				Logs: []interface{}{
//   					&LogProperty{
//   						LogType: jsii.String("logType"),
//
//   						// the properties below are optional
//   						Encoding: jsii.String("encoding"),
//   						LogGroupName: jsii.String("logGroupName"),
//   						LogPath: jsii.String("logPath"),
//   						PatternSet: jsii.String("patternSet"),
//   					},
//   				},
//   				Processes: []interface{}{
//   					&ProcessProperty{
//   						AlarmMetrics: []interface{}{
//   							&AlarmMetricProperty{
//   								AlarmMetricName: jsii.String("alarmMetricName"),
//   							},
//   						},
//   						ProcessName: jsii.String("processName"),
//   					},
//   				},
//   				WindowsEvents: []interface{}{
//   					&WindowsEventProperty{
//   						EventLevels: []*string{
//   							jsii.String("eventLevels"),
//   						},
//   						EventName: jsii.String("eventName"),
//   						LogGroupName: jsii.String("logGroupName"),
//
//   						// the properties below are optional
//   						PatternSet: jsii.String("patternSet"),
//   					},
//   				},
//   			},
//   			SubComponentType: jsii.String("subComponentType"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-componentconfiguration.html
//
type CfnApplication_ComponentConfigurationProperty struct {
	// The configuration settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-componentconfiguration.html#cfn-applicationinsights-application-componentconfiguration-configurationdetails
	//
	ConfigurationDetails interface{} `field:"optional" json:"configurationDetails" yaml:"configurationDetails"`
	// Sub-component configurations of the component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-componentconfiguration.html#cfn-applicationinsights-application-componentconfiguration-subcomponenttypeconfigurations
	//
	SubComponentTypeConfigurations interface{} `field:"optional" json:"subComponentTypeConfigurations" yaml:"subComponentTypeConfigurations"`
}

