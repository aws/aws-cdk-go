package awsnetworkfirewall


// Properties for defining a `CfnLoggingConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLoggingConfigurationProps := &cfnLoggingConfigurationProps{
//   	firewallArn: jsii.String("firewallArn"),
//   	loggingConfiguration: &loggingConfigurationProperty{
//   		logDestinationConfigs: []interface{}{
//   			&logDestinationConfigProperty{
//   				logDestination: map[string]*string{
//   					"logDestinationKey": jsii.String("logDestination"),
//   				},
//   				logDestinationType: jsii.String("logDestinationType"),
//   				logType: jsii.String("logType"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	firewallName: jsii.String("firewallName"),
//   }
//
type CfnLoggingConfigurationProps struct {
	// The Amazon Resource Name (ARN) of the `Firewall` that the logging configuration is associated with.
	//
	// You can't change the firewall specification after you create the logging configuration.
	FirewallArn *string `field:"required" json:"firewallArn" yaml:"firewallArn"`
	// Defines how AWS Network Firewall performs logging for a `Firewall` .
	LoggingConfiguration interface{} `field:"required" json:"loggingConfiguration" yaml:"loggingConfiguration"`
	// The name of the firewall that the logging configuration is associated with.
	//
	// You can't change the firewall specification after you create the logging configuration.
	FirewallName *string `field:"optional" json:"firewallName" yaml:"firewallName"`
}

