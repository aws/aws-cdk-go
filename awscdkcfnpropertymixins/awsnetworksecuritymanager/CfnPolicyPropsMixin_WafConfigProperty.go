package awsnetworksecuritymanager


// WAF-specific policy settings.
//
// Populated only for WAF firewall type policies.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   wafConfigProperty := &WafConfigProperty{
//   	ConflictResolution: jsii.String("conflictResolution"),
//   	ExistingCustomerWebAclResolution: jsii.String("existingCustomerWebAclResolution"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networksecuritymanager-policy-wafconfig.html
//
type CfnPolicyPropsMixin_WafConfigProperty struct {
	// Conflict-resolution strategy applied to AWS WAF policies.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networksecuritymanager-policy-wafconfig.html#cfn-networksecuritymanager-policy-wafconfig-conflictresolution
	//
	ConflictResolution *string `field:"optional" json:"conflictResolution" yaml:"conflictResolution"`
	// Controls how Network Security Manager handles remediation when a resource already has a customer-created WebACL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networksecuritymanager-policy-wafconfig.html#cfn-networksecuritymanager-policy-wafconfig-existingcustomerwebaclresolution
	//
	ExistingCustomerWebAclResolution *string `field:"optional" json:"existingCustomerWebAclResolution" yaml:"existingCustomerWebAclResolution"`
}

