package awsnetworkfirewall


// Defines how AWS Network Firewall performs logging for a firewall.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loggingConfigurationProperty := &LoggingConfigurationProperty{
//   	LogDestinationConfigs: []interface{}{
//   		&LogDestinationConfigProperty{
//   			LogDestination: map[string]*string{
//   				"logDestinationKey": jsii.String("logDestination"),
//   			},
//   			LogDestinationType: jsii.String("logDestinationType"),
//   			LogType: jsii.String("logType"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-loggingconfiguration-loggingconfiguration.html
//
type CfnLoggingConfiguration_LoggingConfigurationProperty struct {
	// Defines the logging destinations for the logs for a firewall.
	//
	// Network Firewall generates logs for stateful rule groups.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-loggingconfiguration-loggingconfiguration.html#cfn-networkfirewall-loggingconfiguration-loggingconfiguration-logdestinationconfigs
	//
	LogDestinationConfigs interface{} `field:"required" json:"logDestinationConfigs" yaml:"logDestinationConfigs"`
}

