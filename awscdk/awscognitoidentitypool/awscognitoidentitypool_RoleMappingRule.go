package awscognitoidentitypool

import (
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// Represents an Identity Pool Role Attachment Role Mapping Rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role role
//
//   roleMappingRule := &RoleMappingRule{
//   	Claim: jsii.String("claim"),
//   	ClaimValue: jsii.String("claimValue"),
//   	MappedRole: role,
//
//   	// the properties below are optional
//   	MatchType: awscdk.Aws_cognito_identitypool.RoleMappingMatchType_EQUALS,
//   }
//
// Experimental.
type RoleMappingRule struct {
	// The key sent in the token by the federated identity provider.
	// Experimental.
	Claim *string `field:"required" json:"claim" yaml:"claim"`
	// The value of the claim that must be matched.
	// Experimental.
	ClaimValue *string `field:"required" json:"claimValue" yaml:"claimValue"`
	// The Role to be assumed when Claim Value is matched.
	// Experimental.
	MappedRole awsiam.IRole `field:"required" json:"mappedRole" yaml:"mappedRole"`
	// How to match with the Claim value.
	// Experimental.
	MatchType RoleMappingMatchType `field:"optional" json:"matchType" yaml:"matchType"`
}

