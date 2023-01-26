package awsiam

import (
	"github.com/aws/constructs-go/constructs/v10"
)

// Result of calling `addToPrincipalPolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import constructs "github.com/aws/constructs-go/constructs"
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
type AddToPrincipalPolicyResult struct {
	// Whether the statement was added to the identity's policies.
	StatementAdded *bool `field:"required" json:"statementAdded" yaml:"statementAdded"`
	// Dependable which allows depending on the policy change being applied.
	PolicyDependable constructs.IDependable `field:"optional" json:"policyDependable" yaml:"policyDependable"`
}

