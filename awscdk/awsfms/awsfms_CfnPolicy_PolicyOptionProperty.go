package awsfms


// Contains the AWS Network Firewall firewall policy options to configure the policy's deployment model and third-party firewall policy settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   policyOptionProperty := &policyOptionProperty{
//   	networkFirewallPolicy: &networkFirewallPolicyProperty{
//   		firewallDeploymentModel: jsii.String("firewallDeploymentModel"),
//   	},
//   	thirdPartyFirewallPolicy: &thirdPartyFirewallPolicyProperty{
//   		firewallDeploymentModel: jsii.String("firewallDeploymentModel"),
//   	},
//   }
//
type CfnPolicy_PolicyOptionProperty struct {
	// Defines the deployment model to use for the firewall policy.
	NetworkFirewallPolicy interface{} `field:"optional" json:"networkFirewallPolicy" yaml:"networkFirewallPolicy"`
	// Defines the policy options for a third-party firewall policy.
	ThirdPartyFirewallPolicy interface{} `field:"optional" json:"thirdPartyFirewallPolicy" yaml:"thirdPartyFirewallPolicy"`
}

