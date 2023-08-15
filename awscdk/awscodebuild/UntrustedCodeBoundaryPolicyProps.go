package awscodebuild

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
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
//   untrustedCodeBoundaryPolicyProps := &UntrustedCodeBoundaryPolicyProps{
//   	AdditionalStatements: []*policyStatement{
//   		policyStatement,
//   	},
//   	ManagedPolicyName: jsii.String("managedPolicyName"),
//   }
//
type UntrustedCodeBoundaryPolicyProps struct {
	// Additional statements to add to the default set of statements.
	// Default: - No additional statements.
	//
	AdditionalStatements *[]awsiam.PolicyStatement `field:"optional" json:"additionalStatements" yaml:"additionalStatements"`
	// The name of the managed policy.
	// Default: - A name is automatically generated.
	//
	ManagedPolicyName *string `field:"optional" json:"managedPolicyName" yaml:"managedPolicyName"`
}

