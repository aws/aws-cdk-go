package previewawsnetworkfirewallmixins


// Properties for CfnLoggingConfigurationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLoggingConfigurationMixinProps := &CfnLoggingConfigurationMixinProps{
//   	EnableMonitoringDashboard: jsii.Boolean(false),
//   	FirewallArn: jsii.String("firewallArn"),
//   	FirewallName: jsii.String("firewallName"),
//   	LoggingConfiguration: &LoggingConfigurationProperty{
//   		LogDestinationConfigs: []interface{}{
//   			&LogDestinationConfigProperty{
//   				LogDestination: map[string]*string{
//   					"logDestinationKey": jsii.String("logDestination"),
//   				},
//   				LogDestinationType: jsii.String("logDestinationType"),
//   				LogType: jsii.String("logType"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-loggingconfiguration.html
//
type CfnLoggingConfigurationMixinProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-loggingconfiguration.html#cfn-networkfirewall-loggingconfiguration-enablemonitoringdashboard
	//
	EnableMonitoringDashboard interface{} `field:"optional" json:"enableMonitoringDashboard" yaml:"enableMonitoringDashboard"`
	// The Amazon Resource Name (ARN) of the firewallthat the logging configuration is associated with.
	//
	// You can't change the firewall specification after you create the logging configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-loggingconfiguration.html#cfn-networkfirewall-loggingconfiguration-firewallarn
	//
	FirewallArn *string `field:"optional" json:"firewallArn" yaml:"firewallArn"`
	// The name of the firewall that the logging configuration is associated with.
	//
	// You can't change the firewall specification after you create the logging configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-loggingconfiguration.html#cfn-networkfirewall-loggingconfiguration-firewallname
	//
	FirewallName *string `field:"optional" json:"firewallName" yaml:"firewallName"`
	// Defines how AWS Network Firewall performs logging for a firewall.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-loggingconfiguration.html#cfn-networkfirewall-loggingconfiguration-loggingconfiguration
	//
	LoggingConfiguration interface{} `field:"optional" json:"loggingConfiguration" yaml:"loggingConfiguration"`
}

