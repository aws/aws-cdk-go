package interfacesawsnetworkfirewall


// A reference to a LoggingConfiguration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loggingConfigurationReference := &LoggingConfigurationReference{
//   	FirewallArn: jsii.String("firewallArn"),
//   }
//
type LoggingConfigurationReference struct {
	// The FirewallArn of the LoggingConfiguration resource.
	FirewallArn *string `field:"required" json:"firewallArn" yaml:"firewallArn"`
}

