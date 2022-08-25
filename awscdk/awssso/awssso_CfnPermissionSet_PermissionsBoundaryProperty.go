package awssso


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   permissionsBoundaryProperty := &permissionsBoundaryProperty{
//   	customerManagedPolicyReference: &customerManagedPolicyReferenceProperty{
//   		name: jsii.String("name"),
//
//   		// the properties below are optional
//   		path: jsii.String("path"),
//   	},
//   	managedPolicyArn: jsii.String("managedPolicyArn"),
//   }
//
type CfnPermissionSet_PermissionsBoundaryProperty struct {
	// `CfnPermissionSet.PermissionsBoundaryProperty.CustomerManagedPolicyReference`.
	CustomerManagedPolicyReference interface{} `field:"optional" json:"customerManagedPolicyReference" yaml:"customerManagedPolicyReference"`
	// `CfnPermissionSet.PermissionsBoundaryProperty.ManagedPolicyArn`.
	ManagedPolicyArn *string `field:"optional" json:"managedPolicyArn" yaml:"managedPolicyArn"`
}

