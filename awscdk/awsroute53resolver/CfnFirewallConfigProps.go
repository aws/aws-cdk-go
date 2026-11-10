package awsroute53resolver


// Properties for defining a `CfnFirewallConfig`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFirewallConfigProps := &CfnFirewallConfigProps{
//   	FirewallFailOpen: jsii.String("firewallFailOpen"),
//   	ResourceId: jsii.String("resourceId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53resolver-firewallconfig.html
//
type CfnFirewallConfigProps struct {
	// Determines how DNS Firewall operates during failures.
	//
	// By default, fail open is disabled, which means the failure mode is closed. If you enable this option, the failure mode is open. This behavior is only enforced for VPCs that have at least one DNS Firewall rule group association.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53resolver-firewallconfig.html#cfn-route53resolver-firewallconfig-firewallfailopen
	//
	FirewallFailOpen *string `field:"required" json:"firewallFailOpen" yaml:"firewallFailOpen"`
	// The ID of the VPC that this firewall configuration applies to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53resolver-firewallconfig.html#cfn-route53resolver-firewallconfig-resourceid
	//
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
}

