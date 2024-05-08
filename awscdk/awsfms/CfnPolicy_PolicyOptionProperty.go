package awsfms


// Contains the settings to configure a network ACL policy, a AWS Network Firewall firewall policy deployment model, or a third-party firewall policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   policyOptionProperty := &PolicyOptionProperty{
//   	NetworkFirewallPolicy: &NetworkFirewallPolicyProperty{
//   		FirewallDeploymentModel: jsii.String("firewallDeploymentModel"),
//   	},
//   	ThirdPartyFirewallPolicy: &ThirdPartyFirewallPolicyProperty{
//   		FirewallDeploymentModel: jsii.String("firewallDeploymentModel"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fms-policy-policyoption.html
//
type CfnPolicy_PolicyOptionProperty struct {
	// Defines the deployment model to use for the firewall policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fms-policy-policyoption.html#cfn-fms-policy-policyoption-networkfirewallpolicy
	//
	NetworkFirewallPolicy interface{} `field:"optional" json:"networkFirewallPolicy" yaml:"networkFirewallPolicy"`
	// Defines the policy options for a third-party firewall policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fms-policy-policyoption.html#cfn-fms-policy-policyoption-thirdpartyfirewallpolicy
	//
	ThirdPartyFirewallPolicy interface{} `field:"optional" json:"thirdPartyFirewallPolicy" yaml:"thirdPartyFirewallPolicy"`
}

