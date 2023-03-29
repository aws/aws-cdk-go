package awsvpclattice

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnServiceNetworkVpcAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnServiceNetworkVpcAssociationProps := &CfnServiceNetworkVpcAssociationProps{
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	ServiceNetworkIdentifier: jsii.String("serviceNetworkIdentifier"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VpcIdentifier: jsii.String("vpcIdentifier"),
//   }
//
type CfnServiceNetworkVpcAssociationProps struct {
	// `AWS::VpcLattice::ServiceNetworkVpcAssociation.SecurityGroupIds`.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// `AWS::VpcLattice::ServiceNetworkVpcAssociation.ServiceNetworkIdentifier`.
	ServiceNetworkIdentifier *string `field:"optional" json:"serviceNetworkIdentifier" yaml:"serviceNetworkIdentifier"`
	// `AWS::VpcLattice::ServiceNetworkVpcAssociation.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// `AWS::VpcLattice::ServiceNetworkVpcAssociation.VpcIdentifier`.
	VpcIdentifier *string `field:"optional" json:"vpcIdentifier" yaml:"vpcIdentifier"`
}

