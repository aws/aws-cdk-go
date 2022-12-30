package awsfms


// Configures the firewall policy deployment model of AWS Network Firewall .
//
// For information about Network Firewall deployment models, see [AWS Network Firewall example architectures with routing](https://docs.aws.amazon.com/network-firewall/latest/developerguide/architectures.html) in the *Network Firewall Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkFirewallPolicyProperty := &networkFirewallPolicyProperty{
//   	firewallDeploymentModel: jsii.String("firewallDeploymentModel"),
//   }
//
type CfnPolicy_NetworkFirewallPolicyProperty struct {
	// Defines the deployment model to use for the firewall policy.
	//
	// To use a distributed model, set [PolicyOption](https://docs.aws.amazon.com/fms/2018-01-01/APIReference/API_PolicyOption.html) to `NULL` .
	FirewallDeploymentModel *string `field:"required" json:"firewallDeploymentModel" yaml:"firewallDeploymentModel"`
}

