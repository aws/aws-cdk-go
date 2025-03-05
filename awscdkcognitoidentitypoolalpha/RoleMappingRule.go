package awscdkcognitoidentitypoolalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Represents an Identity Pool Role Attachment role mapping rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cognito_identitypool_alpha "github.com/aws/aws-cdk-go/awscdkcognitoidentitypoolalpha"
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
//   	MatchType: cognito_identitypool_alpha.RoleMappingMatchType_EQUALS,
//   }
//
// Experimental.
type RoleMappingRule struct {
	// The key sent in the token by the federated Identity Provider.
	// Experimental.
	Claim *string `field:"required" json:"claim" yaml:"claim"`
	// The value of the claim that must be matched.
	// Experimental.
	ClaimValue *string `field:"required" json:"claimValue" yaml:"claimValue"`
	// The role to be assumed when the claim value is matched.
	// Experimental.
	MappedRole awsiam.IRole `field:"required" json:"mappedRole" yaml:"mappedRole"`
	// How to match with the claim value.
	// Default: RoleMappingMatchType.EQUALS
	//
	// Experimental.
	MatchType RoleMappingMatchType `field:"optional" json:"matchType" yaml:"matchType"`
}

