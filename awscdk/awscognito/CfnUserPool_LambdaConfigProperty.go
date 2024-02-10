package awscognito


// Specifies the configuration for AWS Lambda triggers.
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
	// Creates an authentication challenge.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-lambdaconfig.html#cfn-cognito-userpool-lambdaconfig-createauthchallenge
	//
	CreateAuthChallenge *string `field:"optional" json:"createAuthChallenge" yaml:"createAuthChallenge"`
	// A custom email sender AWS Lambda trigger.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-lambdaconfig.html#cfn-cognito-userpool-lambdaconfig-customemailsender
	//
	CustomEmailSender interface{} `field:"optional" json:"customEmailSender" yaml:"customEmailSender"`
	// A custom Message AWS Lambda trigger.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-lambdaconfig.html#cfn-cognito-userpool-lambdaconfig-custommessage
	//
	CustomMessage *string `field:"optional" json:"customMessage" yaml:"customMessage"`
	// A custom SMS sender AWS Lambda trigger.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-lambdaconfig.html#cfn-cognito-userpool-lambdaconfig-customsmssender
	//
	CustomSmsSender interface{} `field:"optional" json:"customSmsSender" yaml:"customSmsSender"`
	// Defines the authentication challenge.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-lambdaconfig.html#cfn-cognito-userpool-lambdaconfig-defineauthchallenge
	//
	DefineAuthChallenge *string `field:"optional" json:"defineAuthChallenge" yaml:"defineAuthChallenge"`
	// The Amazon Resource Name of a AWS Key Management Service ( AWS KMS ) key.
	//
	// Amazon Cognito uses the key to encrypt codes and temporary passwords sent to `CustomEmailSender` and `CustomSMSSender` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-lambdaconfig.html#cfn-cognito-userpool-lambdaconfig-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// A post-authentication AWS Lambda trigger.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-lambdaconfig.html#cfn-cognito-userpool-lambdaconfig-postauthentication
	//
	PostAuthentication *string `field:"optional" json:"postAuthentication" yaml:"postAuthentication"`
	// A post-confirmation AWS Lambda trigger.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-lambdaconfig.html#cfn-cognito-userpool-lambdaconfig-postconfirmation
	//
	PostConfirmation *string `field:"optional" json:"postConfirmation" yaml:"postConfirmation"`
	// A pre-authentication AWS Lambda trigger.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-lambdaconfig.html#cfn-cognito-userpool-lambdaconfig-preauthentication
	//
	PreAuthentication *string `field:"optional" json:"preAuthentication" yaml:"preAuthentication"`
	// A pre-registration AWS Lambda trigger.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-lambdaconfig.html#cfn-cognito-userpool-lambdaconfig-presignup
	//
	PreSignUp *string `field:"optional" json:"preSignUp" yaml:"preSignUp"`
	// The Amazon Resource Name (ARN) of the function that you want to assign to your Lambda trigger.
	//
	// Set this parameter for legacy purposes. If you also set an ARN in `PreTokenGenerationConfig` , its value must be identical to `PreTokenGeneration` . For new instances of pre token generation triggers, set the `LambdaArn` of `PreTokenGenerationConfig` .
	//
	// You can set ``.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-lambdaconfig.html#cfn-cognito-userpool-lambdaconfig-pretokengeneration
	//
	PreTokenGeneration *string `field:"optional" json:"preTokenGeneration" yaml:"preTokenGeneration"`
	// The detailed configuration of a pre token generation trigger.
	//
	// If you also set an ARN in `PreTokenGeneration` , its value must be identical to `PreTokenGenerationConfig` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-lambdaconfig.html#cfn-cognito-userpool-lambdaconfig-pretokengenerationconfig
	//
	PreTokenGenerationConfig interface{} `field:"optional" json:"preTokenGenerationConfig" yaml:"preTokenGenerationConfig"`
	// The user migration Lambda config type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-lambdaconfig.html#cfn-cognito-userpool-lambdaconfig-usermigration
	//
	UserMigration *string `field:"optional" json:"userMigration" yaml:"userMigration"`
	// Verifies the authentication challenge response.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-lambdaconfig.html#cfn-cognito-userpool-lambdaconfig-verifyauthchallengeresponse
	//
	VerifyAuthChallengeResponse *string `field:"optional" json:"verifyAuthChallengeResponse" yaml:"verifyAuthChallengeResponse"`
}

