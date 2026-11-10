package awsec2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   devicePostureTrustProviderProperty := &DevicePostureTrustProviderProperty{
//   	TenantId: jsii.String("tenantId"),
//   	TrustProviderType: jsii.String("trustProviderType"),
//
//   	// the properties below are optional
//   	PublicSigningKeyUrl: jsii.String("publicSigningKeyUrl"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-deviceposturetrustprovider.html
//
type CfnClientVpnEndpoint_DevicePostureTrustProviderProperty struct {
	// The tenant identifier registered with the trust provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-deviceposturetrustprovider.html#cfn-ec2-clientvpnendpoint-deviceposturetrustprovider-tenantid
	//
	TenantId *string `field:"required" json:"tenantId" yaml:"tenantId"`
	// The device-posture trust provider type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-deviceposturetrustprovider.html#cfn-ec2-clientvpnendpoint-deviceposturetrustprovider-trustprovidertype
	//
	TrustProviderType *string `field:"required" json:"trustProviderType" yaml:"trustProviderType"`
	// The URL of the public signing key.
	//
	// Conditionally required by the provider type and enforced server-side (crowdstrike/jumpcloud require it, jamf rejects it).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-deviceposturetrustprovider.html#cfn-ec2-clientvpnendpoint-deviceposturetrustprovider-publicsigningkeyurl
	//
	PublicSigningKeyUrl *string `field:"optional" json:"publicSigningKeyUrl" yaml:"publicSigningKeyUrl"`
}

