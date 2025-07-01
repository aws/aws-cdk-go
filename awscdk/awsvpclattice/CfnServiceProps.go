package awsvpclattice

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnService`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnServiceProps := &CfnServiceProps{
//   	AuthType: jsii.String("authType"),
//   	CertificateArn: jsii.String("certificateArn"),
//   	CustomDomainName: jsii.String("customDomainName"),
//   	DnsEntry: &DnsEntryProperty{
//   		DomainName: jsii.String("domainName"),
//   		HostedZoneId: jsii.String("hostedZoneId"),
//   	},
//   	Name: jsii.String("name"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-service.html
//
type CfnServiceProps struct {
	// The type of IAM policy.
	//
	// - `NONE` : The resource does not use an IAM policy. This is the default.
	// - `AWS_IAM` : The resource uses an IAM policy. When this type is used, auth is enabled and an auth policy is required.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-service.html#cfn-vpclattice-service-authtype
	//
	// Default: - "NONE".
	//
	AuthType *string `field:"optional" json:"authType" yaml:"authType"`
	// The Amazon Resource Name (ARN) of the certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-service.html#cfn-vpclattice-service-certificatearn
	//
	CertificateArn *string `field:"optional" json:"certificateArn" yaml:"certificateArn"`
	// The custom domain name of the service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-service.html#cfn-vpclattice-service-customdomainname
	//
	CustomDomainName *string `field:"optional" json:"customDomainName" yaml:"customDomainName"`
	// Describes the DNS information of the service.
	//
	// This field is read-only.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-service.html#cfn-vpclattice-service-dnsentry
	//
	DnsEntry interface{} `field:"optional" json:"dnsEntry" yaml:"dnsEntry"`
	// The name of the service.
	//
	// The name must be unique within the account. The valid characters are a-z, 0-9, and hyphens (-). You can't use a hyphen as the first or last character, or immediately after another hyphen.
	//
	// If you don't specify a name, CloudFormation generates one. However, if you specify a name, and later want to replace the resource, you must specify a new name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-service.html#cfn-vpclattice-service-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The tags for the service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-service.html#cfn-vpclattice-service-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

