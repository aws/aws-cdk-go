package interfacesawsaccountaccess


// A reference to a Entitlement resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   entitlementReference := &EntitlementReference{
//   	ApplicationArn: jsii.String("applicationArn"),
//   	EntitlementId: jsii.String("entitlementId"),
//   }
//
type EntitlementReference struct {
	// The ApplicationArn of the Entitlement resource.
	ApplicationArn *string `field:"required" json:"applicationArn" yaml:"applicationArn"`
	// The EntitlementId of the Entitlement resource.
	EntitlementId *string `field:"required" json:"entitlementId" yaml:"entitlementId"`
}

