package awscloudfront


// The properties to import from the VPC origin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcOriginAttributes := &VpcOriginAttributes{
//   	DomainName: jsii.String("domainName"),
//   	VpcOriginArn: jsii.String("vpcOriginArn"),
//   	VpcOriginId: jsii.String("vpcOriginId"),
//   }
//
type VpcOriginAttributes struct {
	// The domain name of the CloudFront VPC origin endpoint configuration.
	// Default: - No domain name configured.
	//
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// The ARN of the VPC origin.
	//
	// At least one of vpcOriginArn and vpcOriginId must be provided.
	// Default: - derived from `vpcOriginId`.
	//
	VpcOriginArn *string `field:"optional" json:"vpcOriginArn" yaml:"vpcOriginArn"`
	// The ID of the VPC origin.
	//
	// At least one of vpcOriginArn and vpcOriginId must be provided.
	// Default: - derived from `vpcOriginArn`.
	//
	VpcOriginId *string `field:"optional" json:"vpcOriginId" yaml:"vpcOriginId"`
}

