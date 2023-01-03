package awsfms


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
	// `CfnPolicy.NetworkFirewallPolicyProperty.FirewallDeploymentModel`.
	FirewallDeploymentModel *string `field:"required" json:"firewallDeploymentModel" yaml:"firewallDeploymentModel"`
}

