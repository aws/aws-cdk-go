package interfacesawsappstream


// A reference to a Entitlement resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   entitlementReference := &EntitlementReference{
//   	EntitlementName: jsii.String("entitlementName"),
//   	StackName: jsii.String("stackName"),
//   }
//
type EntitlementReference struct {
	// The Name of the Entitlement resource.
	EntitlementName *string `field:"required" json:"entitlementName" yaml:"entitlementName"`
	// The StackName of the Entitlement resource.
	StackName *string `field:"required" json:"stackName" yaml:"stackName"`
}

