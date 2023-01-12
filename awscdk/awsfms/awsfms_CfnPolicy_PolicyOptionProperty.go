package awsfms


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
	// `CfnPolicy.PolicyOptionProperty.NetworkFirewallPolicy`.
	NetworkFirewallPolicy interface{} `field:"optional" json:"networkFirewallPolicy" yaml:"networkFirewallPolicy"`
	// `CfnPolicy.PolicyOptionProperty.ThirdPartyFirewallPolicy`.
	ThirdPartyFirewallPolicy interface{} `field:"optional" json:"thirdPartyFirewallPolicy" yaml:"thirdPartyFirewallPolicy"`
}

