package awsworkspacesweb


// A reference to a Portal resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   portalReference := &PortalReference{
//   	PortalArn: jsii.String("portalArn"),
//   }
//
type PortalReference struct {
	// The PortalArn of the Portal resource.
	PortalArn *string `field:"required" json:"portalArn" yaml:"portalArn"`
}

