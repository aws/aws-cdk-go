package awssagemaker


// Identifies a Amazon Cognito user group.
//
// A user group can be used in on or more work teams.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cognitoMemberDefinitionProperty := &cognitoMemberDefinitionProperty{
//   	cognitoClientId: jsii.String("cognitoClientId"),
//   	cognitoUserGroup: jsii.String("cognitoUserGroup"),
//   	cognitoUserPool: jsii.String("cognitoUserPool"),
//   }
//
type CfnWorkteam_CognitoMemberDefinitionProperty struct {
	// An identifier for an application client.
	//
	// You must create the app client ID using Amazon Cognito.
	CognitoClientId *string `field:"required" json:"cognitoClientId" yaml:"cognitoClientId"`
	// An identifier for a user group.
	CognitoUserGroup *string `field:"required" json:"cognitoUserGroup" yaml:"cognitoUserGroup"`
	// An identifier for a user pool.
	//
	// The user pool must be in the same region as the service that you are calling.
	CognitoUserPool *string `field:"required" json:"cognitoUserPool" yaml:"cognitoUserPool"`
}

