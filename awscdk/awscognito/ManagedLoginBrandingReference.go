package awscognito


// A reference to a ManagedLoginBranding resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   managedLoginBrandingReference := &ManagedLoginBrandingReference{
//   	ManagedLoginBrandingId: jsii.String("managedLoginBrandingId"),
//   	UserPoolId: jsii.String("userPoolId"),
//   }
//
type ManagedLoginBrandingReference struct {
	// The ManagedLoginBrandingId of the ManagedLoginBranding resource.
	ManagedLoginBrandingId *string `field:"required" json:"managedLoginBrandingId" yaml:"managedLoginBrandingId"`
	// The UserPoolId of the ManagedLoginBranding resource.
	UserPoolId *string `field:"required" json:"userPoolId" yaml:"userPoolId"`
}

