package awscognito


// Defines how to map a claim to a role ARN.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mappingRuleProperty := &mappingRuleProperty{
//   	claim: jsii.String("claim"),
//   	matchType: jsii.String("matchType"),
//   	roleArn: jsii.String("roleArn"),
//   	value: jsii.String("value"),
//   }
//
type CfnIdentityPoolRoleAttachment_MappingRuleProperty struct {
	// The claim name that must be present in the token.
	//
	// For example: "isAdmin" or "paid".
	Claim *string `field:"required" json:"claim" yaml:"claim"`
	// The match condition that specifies how closely the claim value in the IdP token must match `Value` .
	//
	// Valid values are: `Equals` , `Contains` , `StartsWith` , and `NotEqual` .
	MatchType *string `field:"required" json:"matchType" yaml:"matchType"`
	// The Amazon Resource Name (ARN) of the role.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// A brief string that the claim must match.
	//
	// For example, "paid" or "yes".
	Value *string `field:"required" json:"value" yaml:"value"`
}

