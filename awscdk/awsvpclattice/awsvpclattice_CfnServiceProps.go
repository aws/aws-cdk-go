package awsvpclattice

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnService`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnServiceProps := &cfnServiceProps{
//   	authType: jsii.String("authType"),
//   	certificateArn: jsii.String("certificateArn"),
//   	customDomainName: jsii.String("customDomainName"),
//   	dnsEntry: &dnsEntryProperty{
//   		domainName: jsii.String("domainName"),
//   		hostedZoneId: jsii.String("hostedZoneId"),
//   	},
//   	name: jsii.String("name"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnServiceProps struct {
	// `AWS::VpcLattice::Service.AuthType`.
	AuthType *string `field:"optional" json:"authType" yaml:"authType"`
	// `AWS::VpcLattice::Service.CertificateArn`.
	CertificateArn *string `field:"optional" json:"certificateArn" yaml:"certificateArn"`
	// `AWS::VpcLattice::Service.CustomDomainName`.
	CustomDomainName *string `field:"optional" json:"customDomainName" yaml:"customDomainName"`
	// `AWS::VpcLattice::Service.DnsEntry`.
	DnsEntry interface{} `field:"optional" json:"dnsEntry" yaml:"dnsEntry"`
	// `AWS::VpcLattice::Service.Name`.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// `AWS::VpcLattice::Service.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

