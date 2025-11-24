package mixinsawsfms


// Configures the firewall policy deployment model of AWS Network Firewall .
//
// For information about Network Firewall deployment models, see [AWS Network Firewall example architectures with routing](https://docs.aws.amazon.com/network-firewall/latest/developerguide/architectures.html) in the *Network Firewall Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   networkFirewallPolicyProperty := &NetworkFirewallPolicyProperty{
//   	FirewallDeploymentModel: jsii.String("firewallDeploymentModel"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fms-policy-networkfirewallpolicy.html
//
type CfnPolicyPropsMixin_NetworkFirewallPolicyProperty struct {
	// Defines the deployment model to use for the firewall policy.
	//
	// To use a distributed model, set [FirewallDeploymentModel](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fms-policy-thirdpartyfirewallpolicy.html) to `DISTRIBUTED` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fms-policy-networkfirewallpolicy.html#cfn-fms-policy-networkfirewallpolicy-firewalldeploymentmodel
	//
	FirewallDeploymentModel *string `field:"optional" json:"firewallDeploymentModel" yaml:"firewallDeploymentModel"`
}

