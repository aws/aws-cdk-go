package awsnetworkfirewall


// Defines how AWS Network Firewall performs logging for a `Firewall` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loggingConfigurationProperty := &loggingConfigurationProperty{
//   	logDestinationConfigs: []interface{}{
//   		&logDestinationConfigProperty{
//   			logDestination: map[string]*string{
//   				"logDestinationKey": jsii.String("logDestination"),
//   			},
//   			logDestinationType: jsii.String("logDestinationType"),
//   			logType: jsii.String("logType"),
//   		},
//   	},
//   }
//
type CfnLoggingConfiguration_LoggingConfigurationProperty struct {
	// Defines the logging destinations for the logs for a firewall.
	//
	// Network Firewall generates logs for stateful rule groups.
	LogDestinationConfigs interface{} `field:"required" json:"logDestinationConfigs" yaml:"logDestinationConfigs"`
}

