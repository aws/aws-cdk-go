package interfacesawsnetworkfirewall


// A reference to a VpcEndpointAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcEndpointAssociationReference := &VpcEndpointAssociationReference{
//   	VpcEndpointAssociationArn: jsii.String("vpcEndpointAssociationArn"),
//   }
//
type VpcEndpointAssociationReference struct {
	// The VpcEndpointAssociationArn of the VpcEndpointAssociation resource.
	VpcEndpointAssociationArn *string `field:"required" json:"vpcEndpointAssociationArn" yaml:"vpcEndpointAssociationArn"`
}

