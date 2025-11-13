package interfacesawsappstream


// A reference to a ApplicationEntitlementAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   applicationEntitlementAssociationReference := &ApplicationEntitlementAssociationReference{
//   	ApplicationIdentifier: jsii.String("applicationIdentifier"),
//   	EntitlementName: jsii.String("entitlementName"),
//   	StackName: jsii.String("stackName"),
//   }
//
type ApplicationEntitlementAssociationReference struct {
	// The ApplicationIdentifier of the ApplicationEntitlementAssociation resource.
	ApplicationIdentifier *string `field:"required" json:"applicationIdentifier" yaml:"applicationIdentifier"`
	// The EntitlementName of the ApplicationEntitlementAssociation resource.
	EntitlementName *string `field:"required" json:"entitlementName" yaml:"entitlementName"`
	// The StackName of the ApplicationEntitlementAssociation resource.
	StackName *string `field:"required" json:"stackName" yaml:"stackName"`
}

