package awsec2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   devicePostureOptionsProperty := &DevicePostureOptionsProperty{
//   	Enabled: jsii.Boolean(false),
//   	TrustProviders: []interface{}{
//   		&DevicePostureTrustProviderProperty{
//   			TenantId: jsii.String("tenantId"),
//   			TrustProviderType: jsii.String("trustProviderType"),
//
//   			// the properties below are optional
//   			PublicSigningKeyUrl: jsii.String("publicSigningKeyUrl"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-devicepostureoptions.html
//
type CfnClientVpnEndpoint_DevicePostureOptionsProperty struct {
	// Whether device posture is enabled for the Client VPN endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-devicepostureoptions.html#cfn-ec2-clientvpnendpoint-devicepostureoptions-enabled
	//
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The device-posture trust providers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-devicepostureoptions.html#cfn-ec2-clientvpnendpoint-devicepostureoptions-trustproviders
	//
	TrustProviders interface{} `field:"required" json:"trustProviders" yaml:"trustProviders"`
}

