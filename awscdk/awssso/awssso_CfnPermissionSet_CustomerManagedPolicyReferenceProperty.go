package awssso


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customerManagedPolicyReferenceProperty := &customerManagedPolicyReferenceProperty{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	path: jsii.String("path"),
//   }
//
type CfnPermissionSet_CustomerManagedPolicyReferenceProperty struct {
	// `CfnPermissionSet.CustomerManagedPolicyReferenceProperty.Name`.
	Name *string `field:"required" json:"name" yaml:"name"`
	// `CfnPermissionSet.CustomerManagedPolicyReferenceProperty.Path`.
	Path *string `field:"optional" json:"path" yaml:"path"`
}

