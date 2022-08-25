package awsnetworkfirewall


// The ID for a subnet that you want to associate with the firewall.
//
// AWS Network Firewall creates an instance of the associated firewall in each subnet that you specify, to filter traffic in the subnet's Availability Zone.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   subnetMappingProperty := &subnetMappingProperty{
//   	subnetId: jsii.String("subnetId"),
//   }
//
type CfnFirewall_SubnetMappingProperty struct {
	// The unique identifier for the subnet.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
}

