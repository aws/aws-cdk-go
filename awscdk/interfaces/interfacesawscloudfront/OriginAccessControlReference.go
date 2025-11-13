package interfacesawscloudfront


// A reference to a OriginAccessControl resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   originAccessControlReference := &OriginAccessControlReference{
//   	OriginAccessControlId: jsii.String("originAccessControlId"),
//   }
//
type OriginAccessControlReference struct {
	// The Id of the OriginAccessControl resource.
	OriginAccessControlId *string `field:"required" json:"originAccessControlId" yaml:"originAccessControlId"`
}

