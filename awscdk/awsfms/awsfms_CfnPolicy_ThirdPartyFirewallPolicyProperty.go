package awsfms


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   thirdPartyFirewallPolicyProperty := &thirdPartyFirewallPolicyProperty{
//   	firewallDeploymentModel: jsii.String("firewallDeploymentModel"),
//   }
//
type CfnPolicy_ThirdPartyFirewallPolicyProperty struct {
	// `CfnPolicy.ThirdPartyFirewallPolicyProperty.FirewallDeploymentModel`.
	FirewallDeploymentModel *string `field:"required" json:"firewallDeploymentModel" yaml:"firewallDeploymentModel"`
}

