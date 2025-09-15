package awsmediaconnect


// A reference to a FlowEntitlement resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   flowEntitlementReference := &FlowEntitlementReference{
//   	EntitlementArn: jsii.String("entitlementArn"),
//   }
//
type FlowEntitlementReference struct {
	// The EntitlementArn of the FlowEntitlement resource.
	EntitlementArn *string `field:"required" json:"entitlementArn" yaml:"entitlementArn"`
}

