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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53resolver-firewallconfig.html#cfn-route53resolver-firewallconfig-firewallfailopen
	//
	FirewallFailOpen *string `field:"optional" json:"firewallFailOpen" yaml:"firewallFailOpen"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53resolver-firewallconfig.html#cfn-route53resolver-firewallconfig-resourceid
	//
	ResourceId *string `field:"optional" json:"resourceId" yaml:"resourceId"`
}

