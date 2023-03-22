package awsvpclattice

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnServiceNetworkServiceAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnServiceNetworkServiceAssociationProps := &CfnServiceNetworkServiceAssociationProps{
//   	DnsEntry: &DnsEntryProperty{
//   		DomainName: jsii.String("domainName"),
//   		HostedZoneId: jsii.String("hostedZoneId"),
//   	},
//   	ServiceIdentifier: jsii.String("serviceIdentifier"),
//   	ServiceNetworkIdentifier: jsii.String("serviceNetworkIdentifier"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnServiceNetworkServiceAssociationProps struct {
	// `AWS::VpcLattice::ServiceNetworkServiceAssociation.DnsEntry`.
	DnsEntry interface{} `field:"optional" json:"dnsEntry" yaml:"dnsEntry"`
	// `AWS::VpcLattice::ServiceNetworkServiceAssociation.ServiceIdentifier`.
	ServiceIdentifier *string `field:"optional" json:"serviceIdentifier" yaml:"serviceIdentifier"`
	// `AWS::VpcLattice::ServiceNetworkServiceAssociation.ServiceNetworkIdentifier`.
	ServiceNetworkIdentifier *string `field:"optional" json:"serviceNetworkIdentifier" yaml:"serviceNetworkIdentifier"`
	// `AWS::VpcLattice::ServiceNetworkServiceAssociation.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

