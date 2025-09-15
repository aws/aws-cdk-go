package awsec2


// A reference to a IpPoolRouteTableAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ipPoolRouteTableAssociationReference := &IpPoolRouteTableAssociationReference{
//   	AssociationId: jsii.String("associationId"),
//   }
//
type IpPoolRouteTableAssociationReference struct {
	// The AssociationId of the IpPoolRouteTableAssociation resource.
	AssociationId *string `field:"required" json:"associationId" yaml:"associationId"`
}

