package awsappstream


// Properties for defining a `CfnApplicationEntitlementAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationEntitlementAssociationProps := &CfnApplicationEntitlementAssociationProps{
//   	ApplicationIdentifier: jsii.String("applicationIdentifier"),
//   	EntitlementName: jsii.String("entitlementName"),
//   	StackName: jsii.String("stackName"),
//   }
//
type CfnApplicationEntitlementAssociationProps struct {
	// The identifier of the application.
	ApplicationIdentifier *string `field:"required" json:"applicationIdentifier" yaml:"applicationIdentifier"`
	// The name of the entitlement.
	EntitlementName *string `field:"required" json:"entitlementName" yaml:"entitlementName"`
	// The name of the stack.
	StackName *string `field:"required" json:"stackName" yaml:"stackName"`
}

