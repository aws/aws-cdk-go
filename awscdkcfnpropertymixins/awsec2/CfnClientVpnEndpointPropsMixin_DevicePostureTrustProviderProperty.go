package awsec2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   devicePostureTrustProviderProperty := &DevicePostureTrustProviderProperty{
//   	PublicSigningKeyUrl: jsii.String("publicSigningKeyUrl"),
//   	TenantId: jsii.String("tenantId"),
//   	TrustProviderType: jsii.String("trustProviderType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-deviceposturetrustprovider.html
//
type CfnClientVpnEndpointPropsMixin_DevicePostureTrustProviderProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-deviceposturetrustprovider.html#cfn-ec2-clientvpnendpoint-deviceposturetrustprovider-publicsigningkeyurl
	//
	PublicSigningKeyUrl *string `field:"optional" json:"publicSigningKeyUrl" yaml:"publicSigningKeyUrl"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-deviceposturetrustprovider.html#cfn-ec2-clientvpnendpoint-deviceposturetrustprovider-tenantid
	//
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-deviceposturetrustprovider.html#cfn-ec2-clientvpnendpoint-deviceposturetrustprovider-trustprovidertype
	//
	TrustProviderType *string `field:"optional" json:"trustProviderType" yaml:"trustProviderType"`
}

