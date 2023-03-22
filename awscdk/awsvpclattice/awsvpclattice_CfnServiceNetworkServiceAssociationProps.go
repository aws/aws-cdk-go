package awsvpclattice

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnServiceNetworkServiceAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnServiceNetworkServiceAssociationProps := &cfnServiceNetworkServiceAssociationProps{
//   	dnsEntry: &dnsEntryProperty{
//   		domainName: jsii.String("domainName"),
//   		hostedZoneId: jsii.String("hostedZoneId"),
//   	},
//   	serviceIdentifier: jsii.String("serviceIdentifier"),
//   	serviceNetworkIdentifier: jsii.String("serviceNetworkIdentifier"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
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

