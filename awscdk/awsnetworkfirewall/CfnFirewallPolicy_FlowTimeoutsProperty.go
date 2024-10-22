package awsnetworkfirewall


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   flowTimeoutsProperty := &FlowTimeoutsProperty{
//   	TcpIdleTimeoutSeconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-firewallpolicy-flowtimeouts.html
//
type CfnFirewallPolicy_FlowTimeoutsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-firewallpolicy-flowtimeouts.html#cfn-networkfirewall-firewallpolicy-flowtimeouts-tcpidletimeoutseconds
	//
	TcpIdleTimeoutSeconds *float64 `field:"optional" json:"tcpIdleTimeoutSeconds" yaml:"tcpIdleTimeoutSeconds"`
}

