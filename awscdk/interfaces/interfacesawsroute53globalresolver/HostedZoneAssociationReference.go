package interfacesawsroute53globalresolver


// A reference to a HostedZoneAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hostedZoneAssociationReference := &HostedZoneAssociationReference{
//   	HostedZoneAssociationId: jsii.String("hostedZoneAssociationId"),
//   }
//
type HostedZoneAssociationReference struct {
	// The HostedZoneAssociationId of the HostedZoneAssociation resource.
	HostedZoneAssociationId *string `field:"required" json:"hostedZoneAssociationId" yaml:"hostedZoneAssociationId"`
}

