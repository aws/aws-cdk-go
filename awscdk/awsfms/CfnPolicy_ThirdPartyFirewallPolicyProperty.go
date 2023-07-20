package awsfms


// Configures the deployment model for the third-party firewall.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   thirdPartyFirewallPolicyProperty := &ThirdPartyFirewallPolicyProperty{
//   	FirewallDeploymentModel: jsii.String("firewallDeploymentModel"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fms-policy-thirdpartyfirewallpolicy.html
//
type CfnPolicy_ThirdPartyFirewallPolicyProperty struct {
	// Defines the deployment model to use for the third-party firewall policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fms-policy-thirdpartyfirewallpolicy.html#cfn-fms-policy-thirdpartyfirewallpolicy-firewalldeploymentmodel
	//
	FirewallDeploymentModel *string `field:"required" json:"firewallDeploymentModel" yaml:"firewallDeploymentModel"`
}

