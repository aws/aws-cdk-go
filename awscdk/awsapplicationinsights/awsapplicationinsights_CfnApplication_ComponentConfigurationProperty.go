package awsapplicationinsights


// The `AWS::ApplicationInsights::Application ComponentConfiguration` property type defines the configuration settings of the component.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   componentConfigurationProperty := &componentConfigurationProperty{
//   	configurationDetails: &configurationDetailsProperty{
//   		alarmMetrics: []interface{}{
//   			&alarmMetricProperty{
//   				alarmMetricName: jsii.String("alarmMetricName"),
//   			},
//   		},
//   		alarms: []interface{}{
//   			&alarmProperty{
//   				alarmName: jsii.String("alarmName"),
//
//   				// the properties below are optional
//   				severity: jsii.String("severity"),
//   			},
//   		},
//   		haClusterPrometheusExporter: &hAClusterPrometheusExporterProperty{
//   			prometheusPort: jsii.String("prometheusPort"),
//   		},
//   		hanaPrometheusExporter: &hANAPrometheusExporterProperty{
//   			agreeToInstallHanadbClient: jsii.Boolean(false),
//   			hanaPort: jsii.String("hanaPort"),
//   			hanaSecretName: jsii.String("hanaSecretName"),
//   			hanasid: jsii.String("hanasid"),
//
//   			// the properties below are optional
//   			prometheusPort: jsii.String("prometheusPort"),
//   		},
//   		jmxPrometheusExporter: &jMXPrometheusExporterProperty{
//   			hostPort: jsii.String("hostPort"),
//   			jmxurl: jsii.String("jmxurl"),
//   			prometheusPort: jsii.String("prometheusPort"),
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
//   	subComponentTypeConfigurations: []interface{}{
//   		&subComponentTypeConfigurationProperty{
//   			subComponentConfigurationDetails: &subComponentConfigurationDetailsProperty{
//   				alarmMetrics: []interface{}{
//   					&alarmMetricProperty{
//   						alarmMetricName: jsii.String("alarmMetricName"),
//   					},
//   				},
//   				logs: []interface{}{
//   					&logProperty{
//   						logType: jsii.String("logType"),
//
//   						// the properties below are optional
//   						encoding: jsii.String("encoding"),
//   						logGroupName: jsii.String("logGroupName"),
//   						logPath: jsii.String("logPath"),
//   						patternSet: jsii.String("patternSet"),
//   					},
//   				},
//   				windowsEvents: []interface{}{
//   					&windowsEventProperty{
//   						eventLevels: []*string{
//   							jsii.String("eventLevels"),
//   						},
//   						eventName: jsii.String("eventName"),
//   						logGroupName: jsii.String("logGroupName"),
//
//   						// the properties below are optional
//   						patternSet: jsii.String("patternSet"),
//   					},
//   				},
//   			},
//   			subComponentType: jsii.String("subComponentType"),
//   		},
//   	},
//   }
//
type CfnApplication_ComponentConfigurationProperty struct {
	// The configuration settings.
	ConfigurationDetails interface{} `field:"optional" json:"configurationDetails" yaml:"configurationDetails"`
	// Sub-component configurations of the component.
	SubComponentTypeConfigurations interface{} `field:"optional" json:"subComponentTypeConfigurations" yaml:"subComponentTypeConfigurations"`
}

