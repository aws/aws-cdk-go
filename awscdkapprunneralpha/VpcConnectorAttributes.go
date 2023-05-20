package awscdkapprunneralpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Attributes for the App Runner VPC Connector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import apprunner_alpha "github.com/aws/aws-cdk-go/awscdkapprunneralpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var securityGroup securityGroup
//
//   vpcConnectorAttributes := &VpcConnectorAttributes{
//   	SecurityGroups: []iSecurityGroup{
//   		securityGroup,
//   	},
//   	VpcConnectorArn: jsii.String("vpcConnectorArn"),
//   	VpcConnectorName: jsii.String("vpcConnectorName"),
//   	VpcConnectorRevision: jsii.Number(123),
//   }
//
// Experimental.
type VpcConnectorAttributes struct {
	// The security groups associated with the VPC connector.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"required" json:"securityGroups" yaml:"securityGroups"`
	// The ARN of the VPC connector.
	// Experimental.
	VpcConnectorArn *string `field:"required" json:"vpcConnectorArn" yaml:"vpcConnectorArn"`
	// The name of the VPC connector.
	// Experimental.
	VpcConnectorName *string `field:"required" json:"vpcConnectorName" yaml:"vpcConnectorName"`
	// The revision of the VPC connector.
	// Experimental.
	VpcConnectorRevision *float64 `field:"required" json:"vpcConnectorRevision" yaml:"vpcConnectorRevision"`
}

