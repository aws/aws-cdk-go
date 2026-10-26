package awsnetworkfirewall


// The proxy listener configuration for a proxy-mode (no source preservation) firewall.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   proxySettingsProperty := &ProxySettingsProperty{
//   	ListenerProperties: []interface{}{
//   		&ListenerPropertyProperty{
//   			Port: jsii.Number(123),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-firewall-proxysettings.html
//
type CfnFirewall_ProxySettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-firewall-proxysettings.html#cfn-networkfirewall-firewall-proxysettings-listenerproperties
	//
	ListenerProperties interface{} `field:"optional" json:"listenerProperties" yaml:"listenerProperties"`
}

