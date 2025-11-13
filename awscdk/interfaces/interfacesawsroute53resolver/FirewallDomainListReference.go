package interfacesawsroute53resolver


// A reference to a FirewallDomainList resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   firewallDomainListReference := &FirewallDomainListReference{
//   	FirewallDomainListArn: jsii.String("firewallDomainListArn"),
//   	FirewallDomainListId: jsii.String("firewallDomainListId"),
//   }
//
type FirewallDomainListReference struct {
	// The ARN of the FirewallDomainList resource.
	FirewallDomainListArn *string `field:"required" json:"firewallDomainListArn" yaml:"firewallDomainListArn"`
	// The Id of the FirewallDomainList resource.
	FirewallDomainListId *string `field:"required" json:"firewallDomainListId" yaml:"firewallDomainListId"`
}

