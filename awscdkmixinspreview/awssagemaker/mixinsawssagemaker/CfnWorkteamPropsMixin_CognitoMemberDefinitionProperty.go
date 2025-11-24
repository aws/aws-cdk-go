package mixinsawssagemaker


// Identifies a Amazon Cognito user group.
//
// A user group can be used in on or more work teams.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cognitoMemberDefinitionProperty := &CognitoMemberDefinitionProperty{
//   	CognitoClientId: jsii.String("cognitoClientId"),
//   	CognitoUserGroup: jsii.String("cognitoUserGroup"),
//   	CognitoUserPool: jsii.String("cognitoUserPool"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-workteam-cognitomemberdefinition.html
//
type CfnWorkteamPropsMixin_CognitoMemberDefinitionProperty struct {
	// An identifier for an application client.
	//
	// You must create the app client ID using Amazon Cognito .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-workteam-cognitomemberdefinition.html#cfn-sagemaker-workteam-cognitomemberdefinition-cognitoclientid
	//
	CognitoClientId *string `field:"optional" json:"cognitoClientId" yaml:"cognitoClientId"`
	// An identifier for a user group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-workteam-cognitomemberdefinition.html#cfn-sagemaker-workteam-cognitomemberdefinition-cognitousergroup
	//
	CognitoUserGroup *string `field:"optional" json:"cognitoUserGroup" yaml:"cognitoUserGroup"`
	// An identifier for a user pool.
	//
	// The user pool must be in the same region as the service that you are calling.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-workteam-cognitomemberdefinition.html#cfn-sagemaker-workteam-cognitomemberdefinition-cognitouserpool
	//
	CognitoUserPool *string `field:"optional" json:"cognitoUserPool" yaml:"cognitoUserPool"`
}

