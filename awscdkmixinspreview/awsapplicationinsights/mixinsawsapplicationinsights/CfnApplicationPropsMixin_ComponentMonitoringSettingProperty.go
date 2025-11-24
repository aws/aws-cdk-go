package mixinsawsapplicationinsights


// The `AWS::ApplicationInsights::Application ComponentMonitoringSetting` property type defines the monitoring setting of the component.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   componentMonitoringSettingProperty := &ComponentMonitoringSettingProperty{
//   	ComponentArn: jsii.String("componentArn"),
//   	ComponentConfigurationMode: jsii.String("componentConfigurationMode"),
//   	ComponentName: jsii.String("componentName"),
//   	CustomComponentConfiguration: &ComponentConfigurationProperty{
//   		ConfigurationDetails: &ConfigurationDetailsProperty{
//   			AlarmMetrics: []interface{}{
//   				&AlarmMetricProperty{
//   					AlarmMetricName: jsii.String("alarmMetricName"),
//   				},
//   			},
//   			Alarms: []interface{}{
//   				&AlarmProperty{
//   					AlarmName: jsii.String("alarmName"),
//   					Severity: jsii.String("severity"),
//   				},
//   			},
//   			HaClusterPrometheusExporter: &HAClusterPrometheusExporterProperty{
//   				PrometheusPort: jsii.String("prometheusPort"),
//   			},
//   			HanaPrometheusExporter: &HANAPrometheusExporterProperty{
//   				AgreeToInstallHanadbClient: jsii.Boolean(false),
//   				HanaPort: jsii.String("hanaPort"),
//   				HanaSecretName: jsii.String("hanaSecretName"),
//   				Hanasid: jsii.String("hanasid"),
//   				PrometheusPort: jsii.String("prometheusPort"),
//   			},
//   			JmxPrometheusExporter: &JMXPrometheusExporterProperty{
//   				HostPort: jsii.String("hostPort"),
//   				Jmxurl: jsii.String("jmxurl"),
//   				PrometheusPort: jsii.String("prometheusPort"),
//   			},
//   			Logs: []interface{}{
//   				&LogProperty{
//   					Encoding: jsii.String("encoding"),
//   					LogGroupName: jsii.String("logGroupName"),
//   					LogPath: jsii.String("logPath"),
//   					LogType: jsii.String("logType"),
//   					PatternSet: jsii.String("patternSet"),
//   				},
//   			},
//   			NetWeaverPrometheusExporter: &NetWeaverPrometheusExporterProperty{
//   				InstanceNumbers: []*string{
//   					jsii.String("instanceNumbers"),
//   				},
//   				PrometheusPort: jsii.String("prometheusPort"),
//   				Sapsid: jsii.String("sapsid"),
//   			},
//   			Processes: []interface{}{
//   				&ProcessProperty{
//   					AlarmMetrics: []interface{}{
//   						&AlarmMetricProperty{
//   							AlarmMetricName: jsii.String("alarmMetricName"),
//   						},
//   					},
//   					ProcessName: jsii.String("processName"),
//   				},
//   			},
//   			SqlServerPrometheusExporter: &SQLServerPrometheusExporterProperty{
//   				PrometheusPort: jsii.String("prometheusPort"),
//   				SqlSecretName: jsii.String("sqlSecretName"),
//   			},
//   			WindowsEvents: []interface{}{
//   				&WindowsEventProperty{
//   					EventLevels: []*string{
//   						jsii.String("eventLevels"),
//   					},
//   					EventName: jsii.String("eventName"),
//   					LogGroupName: jsii.String("logGroupName"),
//   					PatternSet: jsii.String("patternSet"),
//   				},
//   			},
//   		},
//   		SubComponentTypeConfigurations: []interface{}{
//   			&SubComponentTypeConfigurationProperty{
//   				SubComponentConfigurationDetails: &SubComponentConfigurationDetailsProperty{
//   					AlarmMetrics: []interface{}{
//   						&AlarmMetricProperty{
//   							AlarmMetricName: jsii.String("alarmMetricName"),
//   						},
//   					},
//   					Logs: []interface{}{
//   						&LogProperty{
//   							Encoding: jsii.String("encoding"),
//   							LogGroupName: jsii.String("logGroupName"),
//   							LogPath: jsii.String("logPath"),
//   							LogType: jsii.String("logType"),
//   							PatternSet: jsii.String("patternSet"),
//   						},
//   					},
//   					Processes: []interface{}{
//   						&ProcessProperty{
//   							AlarmMetrics: []interface{}{
//   								&AlarmMetricProperty{
//   									AlarmMetricName: jsii.String("alarmMetricName"),
//   								},
//   							},
//   							ProcessName: jsii.String("processName"),
//   						},
//   					},
//   					WindowsEvents: []interface{}{
//   						&WindowsEventProperty{
//   							EventLevels: []*string{
//   								jsii.String("eventLevels"),
//   							},
//   							EventName: jsii.String("eventName"),
//   							LogGroupName: jsii.String("logGroupName"),
//   							PatternSet: jsii.String("patternSet"),
//   						},
//   					},
//   				},
//   				SubComponentType: jsii.String("subComponentType"),
//   			},
//   		},
//   	},
//   	DefaultOverwriteComponentConfiguration: &ComponentConfigurationProperty{
//   		ConfigurationDetails: &ConfigurationDetailsProperty{
//   			AlarmMetrics: []interface{}{
//   				&AlarmMetricProperty{
//   					AlarmMetricName: jsii.String("alarmMetricName"),
//   				},
//   			},
//   			Alarms: []interface{}{
//   				&AlarmProperty{
//   					AlarmName: jsii.String("alarmName"),
//   					Severity: jsii.String("severity"),
//   				},
//   			},
//   			HaClusterPrometheusExporter: &HAClusterPrometheusExporterProperty{
//   				PrometheusPort: jsii.String("prometheusPort"),
//   			},
//   			HanaPrometheusExporter: &HANAPrometheusExporterProperty{
//   				AgreeToInstallHanadbClient: jsii.Boolean(false),
//   				HanaPort: jsii.String("hanaPort"),
//   				HanaSecretName: jsii.String("hanaSecretName"),
//   				Hanasid: jsii.String("hanasid"),
//   				PrometheusPort: jsii.String("prometheusPort"),
//   			},
//   			JmxPrometheusExporter: &JMXPrometheusExporterProperty{
//   				HostPort: jsii.String("hostPort"),
//   				Jmxurl: jsii.String("jmxurl"),
//   				PrometheusPort: jsii.String("prometheusPort"),
//   			},
//   			Logs: []interface{}{
//   				&LogProperty{
//   					Encoding: jsii.String("encoding"),
//   					LogGroupName: jsii.String("logGroupName"),
//   					LogPath: jsii.String("logPath"),
//   					LogType: jsii.String("logType"),
//   					PatternSet: jsii.String("patternSet"),
//   				},
//   			},
//   			NetWeaverPrometheusExporter: &NetWeaverPrometheusExporterProperty{
//   				InstanceNumbers: []*string{
//   					jsii.String("instanceNumbers"),
//   				},
//   				PrometheusPort: jsii.String("prometheusPort"),
//   				Sapsid: jsii.String("sapsid"),
//   			},
//   			Processes: []interface{}{
//   				&ProcessProperty{
//   					AlarmMetrics: []interface{}{
//   						&AlarmMetricProperty{
//   							AlarmMetricName: jsii.String("alarmMetricName"),
//   						},
//   					},
//   					ProcessName: jsii.String("processName"),
//   				},
//   			},
//   			SqlServerPrometheusExporter: &SQLServerPrometheusExporterProperty{
//   				PrometheusPort: jsii.String("prometheusPort"),
//   				SqlSecretName: jsii.String("sqlSecretName"),
//   			},
//   			WindowsEvents: []interface{}{
//   				&WindowsEventProperty{
//   					EventLevels: []*string{
//   						jsii.String("eventLevels"),
//   					},
//   					EventName: jsii.String("eventName"),
//   					LogGroupName: jsii.String("logGroupName"),
//   					PatternSet: jsii.String("patternSet"),
//   				},
//   			},
//   		},
//   		SubComponentTypeConfigurations: []interface{}{
//   			&SubComponentTypeConfigurationProperty{
//   				SubComponentConfigurationDetails: &SubComponentConfigurationDetailsProperty{
//   					AlarmMetrics: []interface{}{
//   						&AlarmMetricProperty{
//   							AlarmMetricName: jsii.String("alarmMetricName"),
//   						},
//   					},
//   					Logs: []interface{}{
//   						&LogProperty{
//   							Encoding: jsii.String("encoding"),
//   							LogGroupName: jsii.String("logGroupName"),
//   							LogPath: jsii.String("logPath"),
//   							LogType: jsii.String("logType"),
//   							PatternSet: jsii.String("patternSet"),
//   						},
//   					},
//   					Processes: []interface{}{
//   						&ProcessProperty{
//   							AlarmMetrics: []interface{}{
//   								&AlarmMetricProperty{
//   									AlarmMetricName: jsii.String("alarmMetricName"),
//   								},
//   							},
//   							ProcessName: jsii.String("processName"),
//   						},
//   					},
//   					WindowsEvents: []interface{}{
//   						&WindowsEventProperty{
//   							EventLevels: []*string{
//   								jsii.String("eventLevels"),
//   							},
//   							EventName: jsii.String("eventName"),
//   							LogGroupName: jsii.String("logGroupName"),
//   							PatternSet: jsii.String("patternSet"),
//   						},
//   					},
//   				},
//   				SubComponentType: jsii.String("subComponentType"),
//   			},
//   		},
//   	},
//   	Tier: jsii.String("tier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-componentmonitoringsetting.html
//
type CfnApplicationPropsMixin_ComponentMonitoringSettingProperty struct {
	// The ARN of the component.
	//
	// Either the component ARN or the component name is required.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-componentmonitoringsetting.html#cfn-applicationinsights-application-componentmonitoringsetting-componentarn
	//
	ComponentArn *string `field:"optional" json:"componentArn" yaml:"componentArn"`
	// Component monitoring can be configured in one of the following three modes:.
	//
	// - `DEFAULT` : The component will be configured with the recommended default monitoring settings of the selected `Tier` .
	// - `CUSTOM` : The component will be configured with the customized monitoring settings that are specified in `CustomComponentConfiguration` . If used, `CustomComponentConfiguration` must be provided.
	// - `DEFAULT_WITH_OVERWRITE` : The component will be configured with the recommended default monitoring settings of the selected `Tier` , and merged with customized overwrite settings that are specified in `DefaultOverwriteComponentConfiguration` . If used, `DefaultOverwriteComponentConfiguration` must be provided.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-componentmonitoringsetting.html#cfn-applicationinsights-application-componentmonitoringsetting-componentconfigurationmode
	//
	ComponentConfigurationMode *string `field:"optional" json:"componentConfigurationMode" yaml:"componentConfigurationMode"`
	// The name of the component.
	//
	// Either the component ARN or the component name is required.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-componentmonitoringsetting.html#cfn-applicationinsights-application-componentmonitoringsetting-componentname
	//
	ComponentName *string `field:"optional" json:"componentName" yaml:"componentName"`
	// Customized monitoring settings.
	//
	// Required if CUSTOM mode is configured in `ComponentConfigurationMode` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-componentmonitoringsetting.html#cfn-applicationinsights-application-componentmonitoringsetting-customcomponentconfiguration
	//
	CustomComponentConfiguration interface{} `field:"optional" json:"customComponentConfiguration" yaml:"customComponentConfiguration"`
	// Customized overwrite monitoring settings.
	//
	// Required if CUSTOM mode is configured in `ComponentConfigurationMode` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-componentmonitoringsetting.html#cfn-applicationinsights-application-componentmonitoringsetting-defaultoverwritecomponentconfiguration
	//
	DefaultOverwriteComponentConfiguration interface{} `field:"optional" json:"defaultOverwriteComponentConfiguration" yaml:"defaultOverwriteComponentConfiguration"`
	// The tier of the application component.
	//
	// Supported tiers include `DOT_NET_CORE` , `DOT_NET_WORKER` , `DOT_NET_WEB` , `SQL_SERVER` , `SQL_SERVER_ALWAYSON_AVAILABILITY_GROUP` , `SQL_SERVER_FAILOVER_CLUSTER_INSTANCE` , `MYSQL` , `POSTGRESQL` , `JAVA_JMX` , `ORACLE` , `SAP_HANA_MULTI_NODE` , `SAP_HANA_SINGLE_NODE` , `SAP_HANA_HIGH_AVAILABILITY` , `SHAREPOINT` . `ACTIVE_DIRECTORY` , and `DEFAULT` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-componentmonitoringsetting.html#cfn-applicationinsights-application-componentmonitoringsetting-tier
	//
	Tier *string `field:"optional" json:"tier" yaml:"tier"`
}

