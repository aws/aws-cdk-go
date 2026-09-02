package interfacesawsapigatewayv2


// A reference to a PortalProduct resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   portalProductReference := &PortalProductReference{
//   	PortalProductArn: jsii.String("portalProductArn"),
//   }
//
type PortalProductReference struct {
	// The PortalProductArn of the PortalProduct resource.
	PortalProductArn *string `field:"required" json:"portalProductArn" yaml:"portalProductArn"`
}

