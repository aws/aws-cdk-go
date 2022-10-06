package awssagemaker


// Defines an Amazon Cognito or your own OIDC IdP user group that is part of a work team.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   memberDefinitionProperty := &memberDefinitionProperty{
//   	cognitoMemberDefinition: &cognitoMemberDefinitionProperty{
//   		cognitoClientId: jsii.String("cognitoClientId"),
//   		cognitoUserGroup: jsii.String("cognitoUserGroup"),
//   		cognitoUserPool: jsii.String("cognitoUserPool"),
//   	},
//   	oidcMemberDefinition: &oidcMemberDefinitionProperty{
//   		oidcGroups: []*string{
//   			jsii.String("oidcGroups"),
//   		},
//   	},
//   }
//
type CfnWorkteam_MemberDefinitionProperty struct {
	// The Amazon Cognito user group that is part of the work team.
	CognitoMemberDefinition interface{} `field:"optional" json:"cognitoMemberDefinition" yaml:"cognitoMemberDefinition"`
	// `CfnWorkteam.MemberDefinitionProperty.OidcMemberDefinition`.
	OidcMemberDefinition interface{} `field:"optional" json:"oidcMemberDefinition" yaml:"oidcMemberDefinition"`
}

