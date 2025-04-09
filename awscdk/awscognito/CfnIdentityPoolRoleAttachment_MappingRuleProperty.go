package awscognito


// Defines how to map a claim to a role ARN.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mappingRuleProperty := &MappingRuleProperty{
//   	Claim: jsii.String("claim"),
//   	MatchType: jsii.String("matchType"),
//   	RoleArn: jsii.String("roleArn"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypoolroleattachment-mappingrule.html
//
type CfnIdentityPoolRoleAttachment_MappingRuleProperty struct {
	// The claim name that must be present in the token.
	//
	// For example: "isAdmin" or "paid".
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypoolroleattachment-mappingrule.html#cfn-cognito-identitypoolroleattachment-mappingrule-claim
	//
	Claim *string `field:"required" json:"claim" yaml:"claim"`
	// The match condition that specifies how closely the claim value in the IdP token must match `Value` .
	//
	// Valid values are: `Equals` , `Contains` , `StartsWith` , and `NotEqual` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypoolroleattachment-mappingrule.html#cfn-cognito-identitypoolroleattachment-mappingrule-matchtype
	//
	MatchType *string `field:"required" json:"matchType" yaml:"matchType"`
	// The Amazon Resource Name (ARN) of the role.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypoolroleattachment-mappingrule.html#cfn-cognito-identitypoolroleattachment-mappingrule-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// A brief string that the claim must match.
	//
	// For example, "paid" or "yes".
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypoolroleattachment-mappingrule.html#cfn-cognito-identitypoolroleattachment-mappingrule-value
	//
	Value *string `field:"required" json:"value" yaml:"value"`
}

