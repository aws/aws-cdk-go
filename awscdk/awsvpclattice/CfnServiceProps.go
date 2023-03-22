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

