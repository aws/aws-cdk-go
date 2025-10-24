package awscognitoidentitypool

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Represents an Identity Pool Role Attachment role mapping rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role Role
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
type RoleMappingRule struct {
	// The key sent in the token by the federated Identity Provider.
	Claim *string `field:"required" json:"claim" yaml:"claim"`
	// The value of the claim that must be matched.
	ClaimValue *string `field:"required" json:"claimValue" yaml:"claimValue"`
	// The role to be assumed when the claim value is matched.
	MappedRole awsiam.IRole `field:"required" json:"mappedRole" yaml:"mappedRole"`
	// How to match with the claim value.
	// Default: RoleMappingMatchType.EQUALS
	//
	MatchType RoleMappingMatchType `field:"optional" json:"matchType" yaml:"matchType"`
}

