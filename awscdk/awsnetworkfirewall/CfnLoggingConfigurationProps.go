package awsnetworkfirewall


// Properties for defining a `CfnLoggingConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLoggingConfigurationProps := &CfnLoggingConfigurationProps{
//   	FirewallArn: jsii.String("firewallArn"),
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
//
//   	// the properties below are optional
//   	FirewallName: jsii.String("firewallName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-loggingconfiguration.html
//
type CfnLoggingConfigurationProps struct {
	// The Amazon Resource Name (ARN) of the `Firewall` that the logging configuration is associated with.
	//
	// You can't change the firewall specification after you create the logging configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-loggingconfiguration.html#cfn-networkfirewall-loggingconfiguration-firewallarn
	//
	FirewallArn *string `field:"required" json:"firewallArn" yaml:"firewallArn"`
	// Defines how AWS Network Firewall performs logging for a `Firewall` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-loggingconfiguration.html#cfn-networkfirewall-loggingconfiguration-loggingconfiguration
	//
	LoggingConfiguration interface{} `field:"required" json:"loggingConfiguration" yaml:"loggingConfiguration"`
	// The name of the firewall that the logging configuration is associated with.
	//
	// You can't change the firewall specification after you create the logging configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-loggingconfiguration.html#cfn-networkfirewall-loggingconfiguration-firewallname
	//
	FirewallName *string `field:"optional" json:"firewallName" yaml:"firewallName"`
}

