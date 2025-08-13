package awsnetworkfirewall


// Defines the mapping between an Availability Zone and a firewall endpoint for a transit gateway-attached firewall.
//
// Each mapping represents where the firewall can process traffic. You use these mappings when calling `CreateFirewall` , `AssociateAvailabilityZones` , and `DisassociateAvailabilityZones` .
//
// To retrieve the current Availability Zone mappings for a firewall, use `DescribeFirewall` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   availabilityZoneMappingProperty := &AvailabilityZoneMappingProperty{
//   	AvailabilityZone: jsii.String("availabilityZone"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-firewall-availabilityzonemapping.html
//
type CfnFirewall_AvailabilityZoneMappingProperty struct {
	// The ID of the Availability Zone where the firewall endpoint is located.
	//
	// For example, `us-east-2a` . The Availability Zone must be in the same Region as the transit gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-firewall-availabilityzonemapping.html#cfn-networkfirewall-firewall-availabilityzonemapping-availabilityzone
	//
	AvailabilityZone *string `field:"required" json:"availabilityZone" yaml:"availabilityZone"`
}

