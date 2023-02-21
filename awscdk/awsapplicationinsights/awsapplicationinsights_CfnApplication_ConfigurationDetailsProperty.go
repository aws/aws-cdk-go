package awsapplicationinsights


// The `AWS::ApplicationInsights::Application ConfigurationDetails` property type specifies the configuration settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   configurationDetailsProperty := &configurationDetailsProperty{
//   	alarmMetrics: []interface{}{
//   		&alarmMetricProperty{
//   			alarmMetricName: jsii.String("alarmMetricName"),
//   		},
//   	},
//   	alarms: []interface{}{
//   		&alarmProperty{
//   			alarmName: jsii.String("alarmName"),
//
//   			// the properties below are optional
//   			severity: jsii.String("severity"),
//   		},
//   	},
//   	haClusterPrometheusExporter: &hAClusterPrometheusExporterProperty{
//   		prometheusPort: jsii.String("prometheusPort"),
//   	},
//   	hanaPrometheusExporter: &hANAPrometheusExporterProperty{
//   		agreeToInstallHanadbClient: jsii.Boolean(false),
//   		hanaPort: jsii.String("hanaPort"),
//   		hanaSecretName: jsii.String("hanaSecretName"),
//   		hanasid: jsii.String("hanasid"),
//
//   		// the properties below are optional
//   		prometheusPort: jsii.String("prometheusPort"),
//   	},
//   	jmxPrometheusExporter: &jMXPrometheusExporterProperty{
//   		hostPort: jsii.String("hostPort"),
//   		jmxurl: jsii.String("jmxurl"),
//   		prometheusPort: jsii.String("prometheusPort"),
//   	},
//   	logs: []interface{}{
//   		&logProperty{
//   			logType: jsii.String("logType"),
//
//   			// the properties below are optional
//   			encoding: jsii.String("encoding"),
//   			logGroupName: jsii.String("logGroupName"),
//   			logPath: jsii.String("logPath"),
//   			patternSet: jsii.String("patternSet"),
//   		},
//   	},
//   	windowsEvents: []interface{}{
//   		&windowsEventProperty{
//   			eventLevels: []*string{
//   				jsii.String("eventLevels"),
//   			},
//   			eventName: jsii.String("eventName"),
//   			logGroupName: jsii.String("logGroupName"),
//
//   			// the properties below are optional
//   			patternSet: jsii.String("patternSet"),
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

