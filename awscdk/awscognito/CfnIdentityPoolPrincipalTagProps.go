package awscognito


// Properties for defining a `CfnIdentityPoolPrincipalTag`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var principalTags interface{}
//
//   cfnIdentityPoolPrincipalTagProps := &CfnIdentityPoolPrincipalTagProps{
//   	IdentityPoolId: jsii.String("identityPoolId"),
//   	IdentityProviderName: jsii.String("identityProviderName"),
//
//   	// the properties below are optional
//   	PrincipalTags: principalTags,
//   	UseDefaults: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypoolprincipaltag.html
//
type CfnIdentityPoolPrincipalTagProps struct {
	// The identity pool that you want to associate with this principal tag map.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypoolprincipaltag.html#cfn-cognito-identitypoolprincipaltag-identitypoolid
	//
	IdentityPoolId *string `field:"required" json:"identityPoolId" yaml:"identityPoolId"`
	// The identity pool identity provider (IdP) that you want to associate with this principal tag map.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypoolprincipaltag.html#cfn-cognito-identitypoolprincipaltag-identityprovidername
	//
	IdentityProviderName *string `field:"required" json:"identityProviderName" yaml:"identityProviderName"`
	// A JSON-formatted list of user claims and the principal tags that you want to associate with them.
	//
	// When Amazon Cognito requests credentials, it sets the value of the principal tag to the value of the user's claim.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypoolprincipaltag.html#cfn-cognito-identitypoolprincipaltag-principaltags
	//
	PrincipalTags interface{} `field:"optional" json:"principalTags" yaml:"principalTags"`
	// Use a default set of mappings between claims and tags for this provider, instead of a custom map.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypoolprincipaltag.html#cfn-cognito-identitypoolprincipaltag-usedefaults
	//
	UseDefaults interface{} `field:"optional" json:"useDefaults" yaml:"useDefaults"`
}

