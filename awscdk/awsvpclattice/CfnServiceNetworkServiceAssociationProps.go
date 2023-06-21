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
	// The ID or Amazon Resource Name (ARN) of the service.
	ServiceIdentifier *string `field:"optional" json:"serviceIdentifier" yaml:"serviceIdentifier"`
	// The ID or Amazon Resource Name (ARN) of the service network.
	//
	// You must use the ARN if the resources specified in the operation are in different accounts.
	ServiceNetworkIdentifier *string `field:"optional" json:"serviceNetworkIdentifier" yaml:"serviceNetworkIdentifier"`
	// The tags for the association.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

