package awsfms


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   securityServicePolicyDataProperty := &securityServicePolicyDataProperty{
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	managedServiceData: jsii.String("managedServiceData"),
//   	policyOption: &policyOptionProperty{
//   		networkFirewallPolicy: &networkFirewallPolicyProperty{
//   			firewallDeploymentModel: jsii.String("firewallDeploymentModel"),
//   		},
//   		thirdPartyFirewallPolicy: &thirdPartyFirewallPolicyProperty{
//   			firewallDeploymentModel: jsii.String("firewallDeploymentModel"),
//   		},
//   	},
//   }
//
type CfnPolicy_SecurityServicePolicyDataProperty struct {
	// `CfnPolicy.SecurityServicePolicyDataProperty.Type`.
	Type *string `field:"required" json:"type" yaml:"type"`
	// `CfnPolicy.SecurityServicePolicyDataProperty.ManagedServiceData`.
	ManagedServiceData *string `field:"optional" json:"managedServiceData" yaml:"managedServiceData"`
	// `CfnPolicy.SecurityServicePolicyDataProperty.PolicyOption`.
	PolicyOption interface{} `field:"optional" json:"policyOption" yaml:"policyOption"`
}

