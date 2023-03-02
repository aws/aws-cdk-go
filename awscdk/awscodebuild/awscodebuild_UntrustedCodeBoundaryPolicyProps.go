package awscodebuild

import (
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// Construction properties for UntrustedCodeBoundaryPolicy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var policyStatement policyStatement
//
//   untrustedCodeBoundaryPolicyProps := &untrustedCodeBoundaryPolicyProps{
//   	additionalStatements: []*policyStatement{
//   		policyStatement,
//   	},
//   	managedPolicyName: jsii.String("managedPolicyName"),
//   }
//
// Experimental.
type UntrustedCodeBoundaryPolicyProps struct {
	// Additional statements to add to the default set of statements.
	// Experimental.
	AdditionalStatements *[]awsiam.PolicyStatement `field:"optional" json:"additionalStatements" yaml:"additionalStatements"`
	// The name of the managed policy.
	// Experimental.
	ManagedPolicyName *string `field:"optional" json:"managedPolicyName" yaml:"managedPolicyName"`
}

