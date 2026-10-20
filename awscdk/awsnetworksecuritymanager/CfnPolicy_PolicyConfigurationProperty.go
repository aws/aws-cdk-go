package awsnetworksecuritymanager


// Configuration settings for policy behavior.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   policyConfigurationProperty := &PolicyConfigurationProperty{
//   	RemediationEnabled: jsii.Boolean(false),
//   	ResourcesCleanUp: jsii.Boolean(false),
//   	WafConfig: &WafConfigProperty{
//   		ConflictResolution: jsii.String("conflictResolution"),
//   		ExistingCustomerWebAclResolution: jsii.String("existingCustomerWebAclResolution"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networksecuritymanager-policy-policyconfiguration.html
//
type CfnPolicy_PolicyConfigurationProperty struct {
	// Controls automatic remediation of non-compliant resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networksecuritymanager-policy-policyconfiguration.html#cfn-networksecuritymanager-policy-policyconfiguration-remediationenabled
	//
	RemediationEnabled interface{} `field:"optional" json:"remediationEnabled" yaml:"remediationEnabled"`
	// Controls automatic cleanup of unused resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networksecuritymanager-policy-policyconfiguration.html#cfn-networksecuritymanager-policy-policyconfiguration-resourcescleanup
	//
	ResourcesCleanUp interface{} `field:"optional" json:"resourcesCleanUp" yaml:"resourcesCleanUp"`
	// WAF-specific policy settings.
	//
	// Populated only for WAF firewall type policies.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networksecuritymanager-policy-policyconfiguration.html#cfn-networksecuritymanager-policy-policyconfiguration-wafconfig
	//
	WafConfig interface{} `field:"optional" json:"wafConfig" yaml:"wafConfig"`
}

