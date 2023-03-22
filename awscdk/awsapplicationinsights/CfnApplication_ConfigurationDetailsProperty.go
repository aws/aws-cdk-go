package awsapplicationinsights


// The `AWS::ApplicationInsights::Application ConfigurationDetails` property type specifies the configuration settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   configurationDetailsProperty := &ConfigurationDetailsProperty{
//   	AlarmMetrics: []interface{}{
//   		&AlarmMetricProperty{
//   			AlarmMetricName: jsii.String("alarmMetricName"),
//   		},
//   	},
//   	Alarms: []interface{}{
//   		&AlarmProperty{
//   			AlarmName: jsii.String("alarmName"),
//
//   			// the properties below are optional
//   			Severity: jsii.String("severity"),
//   		},
//   	},
//   	HaClusterPrometheusExporter: &HAClusterPrometheusExporterProperty{
//   		PrometheusPort: jsii.String("prometheusPort"),
//   	},
//   	HanaPrometheusExporter: &HANAPrometheusExporterProperty{
//   		AgreeToInstallHanadbClient: jsii.Boolean(false),
//   		HanaPort: jsii.String("hanaPort"),
//   		HanaSecretName: jsii.String("hanaSecretName"),
//   		Hanasid: jsii.String("hanasid"),
//
//   		// the properties below are optional
//   		PrometheusPort: jsii.String("prometheusPort"),
//   	},
//   	JmxPrometheusExporter: &JMXPrometheusExporterProperty{
//   		HostPort: jsii.String("hostPort"),
//   		Jmxurl: jsii.String("jmxurl"),
//   		PrometheusPort: jsii.String("prometheusPort"),
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
type CfnApplication_ConfigurationDetailsProperty struct {
	// A list of metrics to monitor for the component.
	//
	// All component types can use `AlarmMetrics` .
	AlarmMetrics interface{} `field:"optional" json:"alarmMetrics" yaml:"alarmMetrics"`
	// A list of alarms to monitor for the component.
	//
	// All component types can use `Alarm` .
	Alarms interface{} `field:"optional" json:"alarms" yaml:"alarms"`
	// The HA cluster Prometheus Exporter settings.
	HaClusterPrometheusExporter interface{} `field:"optional" json:"haClusterPrometheusExporter" yaml:"haClusterPrometheusExporter"`
	// The HANA DB Prometheus Exporter settings.
	HanaPrometheusExporter interface{} `field:"optional" json:"hanaPrometheusExporter" yaml:"hanaPrometheusExporter"`
	// A list of Java metrics to monitor for the component.
	JmxPrometheusExporter interface{} `field:"optional" json:"jmxPrometheusExporter" yaml:"jmxPrometheusExporter"`
	// A list of logs to monitor for the component.
	//
	// Only Amazon EC2 instances can use `Logs` .
	Logs interface{} `field:"optional" json:"logs" yaml:"logs"`
	// A list of Windows Events to monitor for the component.
	//
	// Only Amazon EC2 instances running on Windows can use `WindowsEvents` .
	WindowsEvents interface{} `field:"optional" json:"windowsEvents" yaml:"windowsEvents"`
}

