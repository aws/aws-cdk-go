package awsec2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   devicePostureOptionsProperty := &DevicePostureOptionsProperty{
//   	Enabled: jsii.Boolean(false),
//   	TrustProviders: []interface{}{
//   		&DevicePostureTrustProviderProperty{
//   			PublicSigningKeyUrl: jsii.String("publicSigningKeyUrl"),
//   			TenantId: jsii.String("tenantId"),
//   			TrustProviderType: jsii.String("trustProviderType"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-devicepostureoptions.html
//
type CfnClientVpnEndpointPropsMixin_DevicePostureOptionsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-devicepostureoptions.html#cfn-ec2-clientvpnendpoint-devicepostureoptions-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-devicepostureoptions.html#cfn-ec2-clientvpnendpoint-devicepostureoptions-trustproviders
	//
	TrustProviders interface{} `field:"optional" json:"trustProviders" yaml:"trustProviders"`
}

