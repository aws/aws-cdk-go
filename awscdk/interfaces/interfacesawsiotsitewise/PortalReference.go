package interfacesawsiotsitewise


// A reference to a Portal resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   portalReference := &PortalReference{
//   	PortalArn: jsii.String("portalArn"),
//   	PortalId: jsii.String("portalId"),
//   }
//
type PortalReference struct {
	// The ARN of the Portal resource.
	PortalArn *string `field:"required" json:"portalArn" yaml:"portalArn"`
	// The PortalId of the Portal resource.
	PortalId *string `field:"required" json:"portalId" yaml:"portalId"`
}

