package interfacesawsec2


// A reference to a SubnetRouteTableAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   subnetRouteTableAssociationReference := &SubnetRouteTableAssociationReference{
//   	SubnetRouteTableAssociationId: jsii.String("subnetRouteTableAssociationId"),
//   }
//
type SubnetRouteTableAssociationReference struct {
	// The Id of the SubnetRouteTableAssociation resource.
	SubnetRouteTableAssociationId *string `field:"required" json:"subnetRouteTableAssociationId" yaml:"subnetRouteTableAssociationId"`
}

