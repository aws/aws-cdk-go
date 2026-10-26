package awsnetworkfirewall


// A port and protocol that a proxy-mode firewall's proxy listener accepts traffic on.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   listenerPropertyProperty := &ListenerPropertyProperty{
//   	Port: jsii.Number(123),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-firewall-listenerproperty.html
//
type CfnFirewall_ListenerPropertyProperty struct {
	// The port the proxy listener accepts traffic on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-firewall-listenerproperty.html#cfn-networkfirewall-firewall-listenerproperty-port
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The protocol the proxy listener uses.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-firewall-listenerproperty.html#cfn-networkfirewall-firewall-listenerproperty-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

