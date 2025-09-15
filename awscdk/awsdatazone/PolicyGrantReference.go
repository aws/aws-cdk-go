package awsdatazone


// A reference to a PolicyGrant resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   policyGrantReference := &PolicyGrantReference{
//   	DomainIdentifier: jsii.String("domainIdentifier"),
//   	EntityIdentifier: jsii.String("entityIdentifier"),
//   	EntityType: jsii.String("entityType"),
//   	GrantId: jsii.String("grantId"),
//   	PolicyType: jsii.String("policyType"),
//   }
//
type PolicyGrantReference struct {
	// The DomainIdentifier of the PolicyGrant resource.
	DomainIdentifier *string `field:"required" json:"domainIdentifier" yaml:"domainIdentifier"`
	// The EntityIdentifier of the PolicyGrant resource.
	EntityIdentifier *string `field:"required" json:"entityIdentifier" yaml:"entityIdentifier"`
	// The EntityType of the PolicyGrant resource.
	EntityType *string `field:"required" json:"entityType" yaml:"entityType"`
	// The GrantId of the PolicyGrant resource.
	GrantId *string `field:"required" json:"grantId" yaml:"grantId"`
	// The PolicyType of the PolicyGrant resource.
	PolicyType *string `field:"required" json:"policyType" yaml:"policyType"`
}

