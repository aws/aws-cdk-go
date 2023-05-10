package awsec2


// Describes a Verified Access trust provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   verifiedAccessTrustProviderProperty := &VerifiedAccessTrustProviderProperty{
//   	Description: jsii.String("description"),
//   	DeviceTrustProviderType: jsii.String("deviceTrustProviderType"),
//   	TrustProviderType: jsii.String("trustProviderType"),
//   	UserTrustProviderType: jsii.String("userTrustProviderType"),
//   	VerifiedAccessTrustProviderId: jsii.String("verifiedAccessTrustProviderId"),
//   }
//
type CfnVerifiedAccessInstance_VerifiedAccessTrustProviderProperty struct {
	// A description for the AWS Verified Access trust provider.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The type of device-based trust provider.
	DeviceTrustProviderType *string `field:"optional" json:"deviceTrustProviderType" yaml:"deviceTrustProviderType"`
	// The type of Verified Access trust provider.
	TrustProviderType *string `field:"optional" json:"trustProviderType" yaml:"trustProviderType"`
	// The type of user-based trust provider.
	UserTrustProviderType *string `field:"optional" json:"userTrustProviderType" yaml:"userTrustProviderType"`
	// The ID of the AWS Verified Access trust provider.
	VerifiedAccessTrustProviderId *string `field:"optional" json:"verifiedAccessTrustProviderId" yaml:"verifiedAccessTrustProviderId"`
}

