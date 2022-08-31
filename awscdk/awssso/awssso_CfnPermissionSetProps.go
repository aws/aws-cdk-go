package awssso

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnPermissionSet`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var inlinePolicy interface{}
//
//   cfnPermissionSetProps := &cfnPermissionSetProps{
//   	instanceArn: jsii.String("instanceArn"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	customerManagedPolicyReferences: []interface{}{
//   		&customerManagedPolicyReferenceProperty{
//   			name: jsii.String("name"),
//
//   			// the properties below are optional
//   			path: jsii.String("path"),
//   		},
//   	},
//   	description: jsii.String("description"),
//   	inlinePolicy: inlinePolicy,
//   	managedPolicies: []*string{
//   		jsii.String("managedPolicies"),
//   	},
//   	permissionsBoundary: &permissionsBoundaryProperty{
//   		customerManagedPolicyReference: &customerManagedPolicyReferenceProperty{
//   			name: jsii.String("name"),
//
//   			// the properties below are optional
//   			path: jsii.String("path"),
//   		},
//   		managedPolicyArn: jsii.String("managedPolicyArn"),
//   	},
//   	relayStateType: jsii.String("relayStateType"),
//   	sessionDuration: jsii.String("sessionDuration"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnPermissionSetProps struct {
	// The ARN of the SSO instance under which the operation will be executed.
	//
	// For more information about ARNs, see [Amazon Resource Names (ARNs) and AWS Service Namespaces](https://docs.aws.amazon.com//general/latest/gr/aws-arns-and-namespaces.html) in the *AWS General Reference* .
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// The name of the permission set.
	Name *string `field:"required" json:"name" yaml:"name"`
	// `AWS::SSO::PermissionSet.CustomerManagedPolicyReferences`.
	CustomerManagedPolicyReferences interface{} `field:"optional" json:"customerManagedPolicyReferences" yaml:"customerManagedPolicyReferences"`
	// The description of the `PermissionSet` .
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The IAM inline policy that is attached to the permission set.
	InlinePolicy interface{} `field:"optional" json:"inlinePolicy" yaml:"inlinePolicy"`
	// A structure that stores the details of the IAM managed policy.
	ManagedPolicies *[]*string `field:"optional" json:"managedPolicies" yaml:"managedPolicies"`
	// `AWS::SSO::PermissionSet.PermissionsBoundary`.
	PermissionsBoundary interface{} `field:"optional" json:"permissionsBoundary" yaml:"permissionsBoundary"`
	// Used to redirect users within the application during the federation authentication process.
	RelayStateType *string `field:"optional" json:"relayStateType" yaml:"relayStateType"`
	// The length of time that the application user sessions are valid for in the ISO-8601 standard.
	SessionDuration *string `field:"optional" json:"sessionDuration" yaml:"sessionDuration"`
	// The tags to attach to the new `PermissionSet` .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

