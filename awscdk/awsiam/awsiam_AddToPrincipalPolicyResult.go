package awsiam

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Result of calling `addToPrincipalPolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var dependable iDependable
//
//   addToPrincipalPolicyResult := &addToPrincipalPolicyResult{
//   	statementAdded: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	policyDependable: dependable,
//   }
//
// Experimental.
type AddToPrincipalPolicyResult struct {
	// Whether the statement was added to the identity's policies.
	// Experimental.
	StatementAdded *bool `field:"required" json:"statementAdded" yaml:"statementAdded"`
	// Dependable which allows depending on the policy change being applied.
	// Experimental.
	PolicyDependable awscdk.IDependable `field:"optional" json:"policyDependable" yaml:"policyDependable"`
}

