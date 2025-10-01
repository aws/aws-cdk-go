package awsappstream


// A reference to a StackFleetAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stackFleetAssociationReference := &StackFleetAssociationReference{
//   	StackFleetAssociationId: jsii.String("stackFleetAssociationId"),
//   }
//
type StackFleetAssociationReference struct {
	// The Id of the StackFleetAssociation resource.
	StackFleetAssociationId *string `field:"required" json:"stackFleetAssociationId" yaml:"stackFleetAssociationId"`
}

