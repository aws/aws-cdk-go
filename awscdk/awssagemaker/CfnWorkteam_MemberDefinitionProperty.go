package awssagemaker


// Defines an Amazon Cognito or your own OIDC IdP user group that is part of a work team.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   memberDefinitionProperty := &MemberDefinitionProperty{
//   	CognitoMemberDefinition: &CognitoMemberDefinitionProperty{
//   		CognitoClientId: jsii.String("cognitoClientId"),
//   		CognitoUserGroup: jsii.String("cognitoUserGroup"),
//   		CognitoUserPool: jsii.String("cognitoUserPool"),
//   	},
//   	OidcMemberDefinition: &OidcMemberDefinitionProperty{
//   		OidcGroups: []*string{
//   			jsii.String("oidcGroups"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-workteam-memberdefinition.html
//
type CfnWorkteam_MemberDefinitionProperty struct {
	// The Amazon Cognito user group that is part of the work team.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-workteam-memberdefinition.html#cfn-sagemaker-workteam-memberdefinition-cognitomemberdefinition
	//
	CognitoMemberDefinition interface{} `field:"optional" json:"cognitoMemberDefinition" yaml:"cognitoMemberDefinition"`
	// A list user groups that exist in your OIDC Identity Provider (IdP).
	//
	// One to ten groups can be used to create a single private work team. When you add a user group to the list of `Groups` , you can add that user group to one or more private work teams. If you add a user group to a private work team, all workers in that user group are added to the work team.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-workteam-memberdefinition.html#cfn-sagemaker-workteam-memberdefinition-oidcmemberdefinition
	//
	OidcMemberDefinition interface{} `field:"optional" json:"oidcMemberDefinition" yaml:"oidcMemberDefinition"`
}

