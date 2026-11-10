package awsnetworkfirewall


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   addressProperty := &AddressProperty{
//   	AddressDefinition: jsii.String("addressDefinition"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-firewallpolicy-address.html
//
type CfnFirewallPolicyPropsMixin_AddressProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-firewallpolicy-address.html#cfn-networkfirewall-firewallpolicy-address-addressdefinition
	//
	AddressDefinition *string `field:"optional" json:"addressDefinition" yaml:"addressDefinition"`
}

