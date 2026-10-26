package awslicensemanagerusersubscriptions


// A Remote Desktop Services license server that is accessed through the endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   licenseServerProperty := &LicenseServerProperty{
//   	HealthStatus: jsii.String("healthStatus"),
//   	Ipv4Address: jsii.String("ipv4Address"),
//   	Ipv6Address: jsii.String("ipv6Address"),
//   	ProvisioningStatus: jsii.String("provisioningStatus"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanagerusersubscriptions-licenseserverendpoint-licenseserver.html
//
type CfnLicenseServerEndpoint_LicenseServerProperty struct {
	// The health status of the license server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanagerusersubscriptions-licenseserverendpoint-licenseserver.html#cfn-licensemanagerusersubscriptions-licenseserverendpoint-licenseserver-healthstatus
	//
	HealthStatus *string `field:"optional" json:"healthStatus" yaml:"healthStatus"`
	// The domain IPv4 address of the license server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanagerusersubscriptions-licenseserverendpoint-licenseserver.html#cfn-licensemanagerusersubscriptions-licenseserverendpoint-licenseserver-ipv4address
	//
	Ipv4Address *string `field:"optional" json:"ipv4Address" yaml:"ipv4Address"`
	// The domain IPv6 address of the license server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanagerusersubscriptions-licenseserverendpoint-licenseserver.html#cfn-licensemanagerusersubscriptions-licenseserverendpoint-licenseserver-ipv6address
	//
	Ipv6Address *string `field:"optional" json:"ipv6Address" yaml:"ipv6Address"`
	// The current state of the provisioning process for the license server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanagerusersubscriptions-licenseserverendpoint-licenseserver.html#cfn-licensemanagerusersubscriptions-licenseserverendpoint-licenseserver-provisioningstatus
	//
	ProvisioningStatus *string `field:"optional" json:"provisioningStatus" yaml:"provisioningStatus"`
}

