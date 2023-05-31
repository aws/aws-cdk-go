package awsec2


// Describes the options for an AWS Verified Access device-identity based trust provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deviceOptionsProperty := &DeviceOptionsProperty{
//   	TenantId: jsii.String("tenantId"),
//   }
//
type CfnVerifiedAccessTrustProvider_DeviceOptionsProperty struct {
	// The ID of the tenant application with the device-identity provider.
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
}

