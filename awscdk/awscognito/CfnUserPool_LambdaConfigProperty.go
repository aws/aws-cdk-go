package awscognito


// A collection of user pool Lambda triggers.
//
// Amazon Cognito invokes triggers at several possible stages of user pool operations. Triggers can modify the outcome of the operations that invoked them.
//
// This data type is a request and response parameter of [CreateUserPool](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_CreateUserPool.html) and [UpdateUserPool](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_UpdateUserPool.html) , and a response parameter of [DescribeUserPool](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_DescribeUserPool.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lambdaConfigProperty := &LambdaConfigProperty{
//   	CreateAuthChallenge: jsii.String("createAuthChallenge"),
//   	CustomEmailSender: &CustomEmailSenderProperty{
//   		LambdaArn: jsii.String("lambdaArn"),
//   		LambdaVersion: jsii.String("lambdaVersion"),
//   	},
//   	CustomMessage: jsii.String("customMessage"),
//   	CustomSmsSender: &CustomSMSSenderProperty{
//   		LambdaArn: jsii.String("lambdaArn"),
//   		LambdaVersion: jsii.String("lambdaVersion"),
//   	},
//   	DefineAuthChallenge: jsii.String("defineAuthChallenge"),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	PostAuthentication: jsii.String("postAuthentication"),
//   	PostConfirmation: jsii.String("postConfirmation"),
//   	PreAuthentication: jsii.String("preAuthentication"),
//   	PreSignUp: jsii.String("preSignUp"),
//   	PreTokenGeneration: jsii.String("preTokenGeneration"),
//   	PreTokenGenerationConfig: &PreTokenGenerationConfigProperty{
//   		LambdaArn: jsii.String("lambdaArn"),
//   		LambdaVersion: jsii.String("lambdaVersion"),
//   	},
//   	UserMigration: jsii.String("userMigration"),
//   	VerifyAuthChallengeResponse: jsii.String("verifyAuthChallengeResponse"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-lambdaconfig.html
//
type CfnUserPool_LambdaConfigProperty struct {
	// The configuration of a create auth challenge Lambda trigger, one of three triggers in the sequence of the [custom authentication challenge triggers](https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-lambda-challenge.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-lambdaconfig.html#cfn-cognito-userpool-lambdaconfig-createauthchallenge
	//
	CreateAuthChallenge *string `field:"optional" json:"createAuthChallenge" yaml:"createAuthChallenge"`
	// A custom email sender AWS Lambda trigger.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-lambdaconfig.html#cfn-cognito-userpool-lambdaconfig-customemailsender
	//
	CustomEmailSender interface{} `field:"optional" json:"customEmailSender" yaml:"customEmailSender"`
	// A custom message Lambda trigger.
	//
	// This trigger is an opportunity to customize all SMS and email messages from your user pool. When a custom message trigger is active, your user pool routes all messages to a Lambda function that returns a runtime-customized message subject and body for your user pool to deliver to a user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-lambdaconfig.html#cfn-cognito-userpool-lambdaconfig-custommessage
	//
	CustomMessage *string `field:"optional" json:"customMessage" yaml:"customMessage"`
	// A custom SMS sender AWS Lambda trigger.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-lambdaconfig.html#cfn-cognito-userpool-lambdaconfig-customsmssender
	//
	CustomSmsSender interface{} `field:"optional" json:"customSmsSender" yaml:"customSmsSender"`
	// The configuration of a define auth challenge Lambda trigger, one of three triggers in the sequence of the [custom authentication challenge triggers](https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-lambda-challenge.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-lambdaconfig.html#cfn-cognito-userpool-lambdaconfig-defineauthchallenge
	//
	DefineAuthChallenge *string `field:"optional" json:"defineAuthChallenge" yaml:"defineAuthChallenge"`
	// The Amazon Resource Name of a AWS Key Management Service ( AWS KMS ) key.
	//
	// Amazon Cognito uses the key to encrypt codes and temporary passwords sent to `CustomEmailSender` and `CustomSMSSender` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-lambdaconfig.html#cfn-cognito-userpool-lambdaconfig-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The configuration of a [post authentication Lambda trigger](https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-lambda-post-authentication.html) in a user pool. This trigger can take custom actions after a user signs in.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-lambdaconfig.html#cfn-cognito-userpool-lambdaconfig-postauthentication
	//
	PostAuthentication *string `field:"optional" json:"postAuthentication" yaml:"postAuthentication"`
	// The configuration of a [post confirmation Lambda trigger](https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-lambda-post-confirmation.html) in a user pool. This trigger can take custom actions after a user confirms their user account and their email address or phone number.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-lambdaconfig.html#cfn-cognito-userpool-lambdaconfig-postconfirmation
	//
	PostConfirmation *string `field:"optional" json:"postConfirmation" yaml:"postConfirmation"`
	// The configuration of a [pre authentication trigger](https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-lambda-pre-authentication.html) in a user pool. This trigger can evaluate and modify user sign-in events.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-lambdaconfig.html#cfn-cognito-userpool-lambdaconfig-preauthentication
	//
	PreAuthentication *string `field:"optional" json:"preAuthentication" yaml:"preAuthentication"`
	// The configuration of a [pre sign-up Lambda trigger](https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-lambda-pre-sign-up.html) in a user pool. This trigger evaluates new users and can bypass confirmation, [link a federated user profile](https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-identity-federation-consolidate-users.html) , or block sign-up requests.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-lambdaconfig.html#cfn-cognito-userpool-lambdaconfig-presignup
	//
	PreSignUp *string `field:"optional" json:"preSignUp" yaml:"preSignUp"`
	// The legacy configuration of a [pre token generation Lambda trigger](https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-lambda-pre-token-generation.html) in a user pool.
	//
	// Set this parameter for legacy purposes. If you also set an ARN in `PreTokenGenerationConfig` , its value must be identical to `PreTokenGeneration` . For new instances of pre token generation triggers, set the `LambdaArn` of `PreTokenGenerationConfig` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-lambdaconfig.html#cfn-cognito-userpool-lambdaconfig-pretokengeneration
	//
	PreTokenGeneration *string `field:"optional" json:"preTokenGeneration" yaml:"preTokenGeneration"`
	// The detailed configuration of a [pre token generation Lambda trigger](https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-lambda-pre-token-generation.html) in a user pool. If you also set an ARN in `PreTokenGeneration` , its value must be identical to `PreTokenGenerationConfig` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-lambdaconfig.html#cfn-cognito-userpool-lambdaconfig-pretokengenerationconfig
	//
	PreTokenGenerationConfig interface{} `field:"optional" json:"preTokenGenerationConfig" yaml:"preTokenGenerationConfig"`
	// The configuration of a [migrate user Lambda trigger](https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-lambda-migrate-user.html) in a user pool. This trigger can create user profiles when users sign in or attempt to reset their password with credentials that don't exist yet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-lambdaconfig.html#cfn-cognito-userpool-lambdaconfig-usermigration
	//
	UserMigration *string `field:"optional" json:"userMigration" yaml:"userMigration"`
	// The configuration of a verify auth challenge Lambda trigger, one of three triggers in the sequence of the [custom authentication challenge triggers](https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-lambda-challenge.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-lambdaconfig.html#cfn-cognito-userpool-lambdaconfig-verifyauthchallengeresponse
	//
	VerifyAuthChallengeResponse *string `field:"optional" json:"verifyAuthChallengeResponse" yaml:"verifyAuthChallengeResponse"`
}

