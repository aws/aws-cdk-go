package awsec2


// A trust provider is a third-party entity that creates, maintains, and manages identity information for users and devices.
//
// When an application request is made, the identity information sent by the trust provider is evaluated by Verified Access before allowing or denying the application request.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccessinstance-verifiedaccesstrustprovider.html
//
type CfnVerifiedAccessInstance_VerifiedAccessTrustProviderProperty struct {
	// A description for the AWS Verified Access trust provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccessinstance-verifiedaccesstrustprovider.html#cfn-ec2-verifiedaccessinstance-verifiedaccesstrustprovider-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The type of device-based trust provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccessinstance-verifiedaccesstrustprovider.html#cfn-ec2-verifiedaccessinstance-verifiedaccesstrustprovider-devicetrustprovidertype
	//
	DeviceTrustProviderType *string `field:"optional" json:"deviceTrustProviderType" yaml:"deviceTrustProviderType"`
	// The type of Verified Access trust provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccessinstance-verifiedaccesstrustprovider.html#cfn-ec2-verifiedaccessinstance-verifiedaccesstrustprovider-trustprovidertype
	//
	TrustProviderType *string `field:"optional" json:"trustProviderType" yaml:"trustProviderType"`
	// The type of user-based trust provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccessinstance-verifiedaccesstrustprovider.html#cfn-ec2-verifiedaccessinstance-verifiedaccesstrustprovider-usertrustprovidertype
	//
	UserTrustProviderType *string `field:"optional" json:"userTrustProviderType" yaml:"userTrustProviderType"`
	// The ID of the AWS Verified Access trust provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccessinstance-verifiedaccesstrustprovider.html#cfn-ec2-verifiedaccessinstance-verifiedaccesstrustprovider-verifiedaccesstrustproviderid
	//
	VerifiedAccessTrustProviderId *string `field:"optional" json:"verifiedAccessTrustProviderId" yaml:"verifiedAccessTrustProviderId"`
}

