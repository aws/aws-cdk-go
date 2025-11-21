package awsec2


// The tunnel options for a single VPN tunnel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpnTunnelOptionsSpecificationProperty := &VpnTunnelOptionsSpecificationProperty{
//   	DpdTimeoutAction: jsii.String("dpdTimeoutAction"),
//   	DpdTimeoutSeconds: jsii.Number(123),
//   	EnableTunnelLifecycleControl: jsii.Boolean(false),
//   	IkeVersions: []interface{}{
//   		map[string]*string{
//   			"value": jsii.String("value"),
//   		},
//   	},
//   	LogOptions: &VpnTunnelLogOptionsSpecificationProperty{
//   		CloudwatchLogOptions: &CloudwatchLogOptionsSpecificationProperty{
//   			BgpLogEnabled: jsii.Boolean(false),
//   			BgpLogGroupArn: jsii.String("bgpLogGroupArn"),
//   			BgpLogOutputFormat: jsii.String("bgpLogOutputFormat"),
//   			LogEnabled: jsii.Boolean(false),
//   			LogGroupArn: jsii.String("logGroupArn"),
//   			LogOutputFormat: jsii.String("logOutputFormat"),
//   		},
//   	},
//   	Phase1DhGroupNumbers: []interface{}{
//   		&Phase1DHGroupNumbersRequestListValueProperty{
//   			Value: jsii.Number(123),
//   		},
//   	},
//   	Phase1EncryptionAlgorithms: []interface{}{
//   		&Phase1EncryptionAlgorithmsRequestListValueProperty{
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Phase1IntegrityAlgorithms: []interface{}{
//   		&Phase1IntegrityAlgorithmsRequestListValueProperty{
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Phase1LifetimeSeconds: jsii.Number(123),
//   	Phase2DhGroupNumbers: []interface{}{
//   		&Phase2DHGroupNumbersRequestListValueProperty{
//   			Value: jsii.Number(123),
//   		},
//   	},
//   	Phase2EncryptionAlgorithms: []interface{}{
//   		&Phase2EncryptionAlgorithmsRequestListValueProperty{
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Phase2IntegrityAlgorithms: []interface{}{
//   		&Phase2IntegrityAlgorithmsRequestListValueProperty{
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Phase2LifetimeSeconds: jsii.Number(123),
//   	PreSharedKey: jsii.String("preSharedKey"),
//   	RekeyFuzzPercentage: jsii.Number(123),
//   	RekeyMarginTimeSeconds: jsii.Number(123),
//   	ReplayWindowSize: jsii.Number(123),
//   	StartupAction: jsii.String("startupAction"),
//   	TunnelInsideCidr: jsii.String("tunnelInsideCidr"),
//   	TunnelInsideIpv6Cidr: jsii.String("tunnelInsideIpv6Cidr"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-vpntunneloptionsspecification.html
//
type CfnVPNConnection_VpnTunnelOptionsSpecificationProperty struct {
	// The action to take after DPD timeout occurs.
	//
	// Specify `restart` to restart the IKE initiation. Specify `clear` to end the IKE session.
	//
	// Valid Values: `clear` | `none` | `restart`
	//
	// Default: `clear`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-vpntunneloptionsspecification.html#cfn-ec2-vpnconnection-vpntunneloptionsspecification-dpdtimeoutaction
	//
	DpdTimeoutAction *string `field:"optional" json:"dpdTimeoutAction" yaml:"dpdTimeoutAction"`
	// The number of seconds after which a DPD timeout occurs.
	//
	// Constraints: A value greater than or equal to 30.
	//
	// Default: `30`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-vpntunneloptionsspecification.html#cfn-ec2-vpnconnection-vpntunneloptionsspecification-dpdtimeoutseconds
	//
	DpdTimeoutSeconds *float64 `field:"optional" json:"dpdTimeoutSeconds" yaml:"dpdTimeoutSeconds"`
	// Turn on or off tunnel endpoint lifecycle control feature.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-vpntunneloptionsspecification.html#cfn-ec2-vpnconnection-vpntunneloptionsspecification-enabletunnellifecyclecontrol
	//
	EnableTunnelLifecycleControl interface{} `field:"optional" json:"enableTunnelLifecycleControl" yaml:"enableTunnelLifecycleControl"`
	// The IKE versions that are permitted for the VPN tunnel.
	//
	// Valid values: `ikev1` | `ikev2`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-vpntunneloptionsspecification.html#cfn-ec2-vpnconnection-vpntunneloptionsspecification-ikeversions
	//
	IkeVersions interface{} `field:"optional" json:"ikeVersions" yaml:"ikeVersions"`
	// Options for logging VPN tunnel activity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-vpntunneloptionsspecification.html#cfn-ec2-vpnconnection-vpntunneloptionsspecification-logoptions
	//
	LogOptions interface{} `field:"optional" json:"logOptions" yaml:"logOptions"`
	// One or more Diffie-Hellman group numbers that are permitted for the VPN tunnel for phase 1 IKE negotiations.
	//
	// Valid values: `2` | `14` | `15` | `16` | `17` | `18` | `19` | `20` | `21` | `22` | `23` | `24`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-vpntunneloptionsspecification.html#cfn-ec2-vpnconnection-vpntunneloptionsspecification-phase1dhgroupnumbers
	//
	Phase1DhGroupNumbers interface{} `field:"optional" json:"phase1DhGroupNumbers" yaml:"phase1DhGroupNumbers"`
	// One or more encryption algorithms that are permitted for the VPN tunnel for phase 1 IKE negotiations.
	//
	// Valid values: `AES128` | `AES256` | `AES128-GCM-16` | `AES256-GCM-16`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-vpntunneloptionsspecification.html#cfn-ec2-vpnconnection-vpntunneloptionsspecification-phase1encryptionalgorithms
	//
	Phase1EncryptionAlgorithms interface{} `field:"optional" json:"phase1EncryptionAlgorithms" yaml:"phase1EncryptionAlgorithms"`
	// One or more integrity algorithms that are permitted for the VPN tunnel for phase 1 IKE negotiations.
	//
	// Valid values: `SHA1` | `SHA2-256` | `SHA2-384` | `SHA2-512`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-vpntunneloptionsspecification.html#cfn-ec2-vpnconnection-vpntunneloptionsspecification-phase1integrityalgorithms
	//
	Phase1IntegrityAlgorithms interface{} `field:"optional" json:"phase1IntegrityAlgorithms" yaml:"phase1IntegrityAlgorithms"`
	// The lifetime for phase 1 of the IKE negotiation, in seconds.
	//
	// Constraints: A value between 900 and 28,800.
	//
	// Default: `28800`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-vpntunneloptionsspecification.html#cfn-ec2-vpnconnection-vpntunneloptionsspecification-phase1lifetimeseconds
	//
	Phase1LifetimeSeconds *float64 `field:"optional" json:"phase1LifetimeSeconds" yaml:"phase1LifetimeSeconds"`
	// One or more Diffie-Hellman group numbers that are permitted for the VPN tunnel for phase 2 IKE negotiations.
	//
	// Valid values: `2` | `5` | `14` | `15` | `16` | `17` | `18` | `19` | `20` | `21` | `22` | `23` | `24`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-vpntunneloptionsspecification.html#cfn-ec2-vpnconnection-vpntunneloptionsspecification-phase2dhgroupnumbers
	//
	Phase2DhGroupNumbers interface{} `field:"optional" json:"phase2DhGroupNumbers" yaml:"phase2DhGroupNumbers"`
	// One or more encryption algorithms that are permitted for the VPN tunnel for phase 2 IKE negotiations.
	//
	// Valid values: `AES128` | `AES256` | `AES128-GCM-16` | `AES256-GCM-16`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-vpntunneloptionsspecification.html#cfn-ec2-vpnconnection-vpntunneloptionsspecification-phase2encryptionalgorithms
	//
	Phase2EncryptionAlgorithms interface{} `field:"optional" json:"phase2EncryptionAlgorithms" yaml:"phase2EncryptionAlgorithms"`
	// One or more integrity algorithms that are permitted for the VPN tunnel for phase 2 IKE negotiations.
	//
	// Valid values: `SHA1` | `SHA2-256` | `SHA2-384` | `SHA2-512`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-vpntunneloptionsspecification.html#cfn-ec2-vpnconnection-vpntunneloptionsspecification-phase2integrityalgorithms
	//
	Phase2IntegrityAlgorithms interface{} `field:"optional" json:"phase2IntegrityAlgorithms" yaml:"phase2IntegrityAlgorithms"`
	// The lifetime for phase 2 of the IKE negotiation, in seconds.
	//
	// Constraints: A value between 900 and 3,600. The value must be less than the value for `Phase1LifetimeSeconds` .
	//
	// Default: `3600`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-vpntunneloptionsspecification.html#cfn-ec2-vpnconnection-vpntunneloptionsspecification-phase2lifetimeseconds
	//
	Phase2LifetimeSeconds *float64 `field:"optional" json:"phase2LifetimeSeconds" yaml:"phase2LifetimeSeconds"`
	// The pre-shared key (PSK) to establish initial authentication between the virtual private gateway and customer gateway.
	//
	// Constraints: Allowed characters are alphanumeric characters, periods (.), and underscores (_). Must be between 8 and 64 characters in length and cannot start with zero (0).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-vpntunneloptionsspecification.html#cfn-ec2-vpnconnection-vpntunneloptionsspecification-presharedkey
	//
	PreSharedKey *string `field:"optional" json:"preSharedKey" yaml:"preSharedKey"`
	// The percentage of the rekey window (determined by `RekeyMarginTimeSeconds` ) during which the rekey time is randomly selected.
	//
	// Constraints: A value between 0 and 100.
	//
	// Default: `100`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-vpntunneloptionsspecification.html#cfn-ec2-vpnconnection-vpntunneloptionsspecification-rekeyfuzzpercentage
	//
	RekeyFuzzPercentage *float64 `field:"optional" json:"rekeyFuzzPercentage" yaml:"rekeyFuzzPercentage"`
	// The margin time, in seconds, before the phase 2 lifetime expires, during which the AWS side of the VPN connection performs an IKE rekey.
	//
	// The exact time of the rekey is randomly selected based on the value for `RekeyFuzzPercentage` .
	//
	// Constraints: A value between 60 and half of `Phase2LifetimeSeconds` .
	//
	// Default: `270`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-vpntunneloptionsspecification.html#cfn-ec2-vpnconnection-vpntunneloptionsspecification-rekeymargintimeseconds
	//
	RekeyMarginTimeSeconds *float64 `field:"optional" json:"rekeyMarginTimeSeconds" yaml:"rekeyMarginTimeSeconds"`
	// The number of packets in an IKE replay window.
	//
	// Constraints: A value between 64 and 2048.
	//
	// Default: `1024`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-vpntunneloptionsspecification.html#cfn-ec2-vpnconnection-vpntunneloptionsspecification-replaywindowsize
	//
	ReplayWindowSize *float64 `field:"optional" json:"replayWindowSize" yaml:"replayWindowSize"`
	// The action to take when the establishing the tunnel for the VPN connection.
	//
	// By default, your customer gateway device must initiate the IKE negotiation and bring up the tunnel. Specify `start` for AWS to initiate the IKE negotiation.
	//
	// Valid Values: `add` | `start`
	//
	// Default: `add`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-vpntunneloptionsspecification.html#cfn-ec2-vpnconnection-vpntunneloptionsspecification-startupaction
	//
	StartupAction *string `field:"optional" json:"startupAction" yaml:"startupAction"`
	// The range of inside IP addresses for the tunnel.
	//
	// Any specified CIDR blocks must be unique across all VPN connections that use the same virtual private gateway.
	//
	// Constraints: A size /30 CIDR block from the `169.254.0.0/16` range. The following CIDR blocks are reserved and cannot be used:
	//
	// - `169.254.0.0/30`
	// - `169.254.1.0/30`
	// - `169.254.2.0/30`
	// - `169.254.3.0/30`
	// - `169.254.4.0/30`
	// - `169.254.5.0/30`
	// - `169.254.169.252/30`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-vpntunneloptionsspecification.html#cfn-ec2-vpnconnection-vpntunneloptionsspecification-tunnelinsidecidr
	//
	TunnelInsideCidr *string `field:"optional" json:"tunnelInsideCidr" yaml:"tunnelInsideCidr"`
	// The range of inside IPv6 addresses for the tunnel.
	//
	// Any specified CIDR blocks must be unique across all VPN connections that use the same transit gateway.
	//
	// Constraints: A size /126 CIDR block from the local `fd00::/8` range.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-vpntunneloptionsspecification.html#cfn-ec2-vpnconnection-vpntunneloptionsspecification-tunnelinsideipv6cidr
	//
	TunnelInsideIpv6Cidr *string `field:"optional" json:"tunnelInsideIpv6Cidr" yaml:"tunnelInsideIpv6Cidr"`
}

