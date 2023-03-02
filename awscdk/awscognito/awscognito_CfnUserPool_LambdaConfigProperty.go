package awscognito


// Specifies the configuration for AWS Lambda triggers.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lambdaConfigProperty := &lambdaConfigProperty{
//   	createAuthChallenge: jsii.String("createAuthChallenge"),
//   	customEmailSender: &customEmailSenderProperty{
//   		lambdaArn: jsii.String("lambdaArn"),
//   		lambdaVersion: jsii.String("lambdaVersion"),
//   	},
//   	customMessage: jsii.String("customMessage"),
//   	customSmsSender: &customSMSSenderProperty{
//   		lambdaArn: jsii.String("lambdaArn"),
//   		lambdaVersion: jsii.String("lambdaVersion"),
//   	},
//   	defineAuthChallenge: jsii.String("defineAuthChallenge"),
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	postAuthentication: jsii.String("postAuthentication"),
//   	postConfirmation: jsii.String("postConfirmation"),
//   	preAuthentication: jsii.String("preAuthentication"),
//   	preSignUp: jsii.String("preSignUp"),
//   	preTokenGeneration: jsii.String("preTokenGeneration"),
//   	userMigration: jsii.String("userMigration"),
//   	verifyAuthChallengeResponse: jsii.String("verifyAuthChallengeResponse"),
//   }
//
type CfnUserPool_LambdaConfigProperty struct {
	// Creates an authentication challenge.
	CreateAuthChallenge *string `field:"optional" json:"createAuthChallenge" yaml:"createAuthChallenge"`
	// A custom email sender AWS Lambda trigger.
	CustomEmailSender interface{} `field:"optional" json:"customEmailSender" yaml:"customEmailSender"`
	// A custom Message AWS Lambda trigger.
	CustomMessage *string `field:"optional" json:"customMessage" yaml:"customMessage"`
	// A custom SMS sender AWS Lambda trigger.
	CustomSmsSender interface{} `field:"optional" json:"customSmsSender" yaml:"customSmsSender"`
	// Defines the authentication challenge.
	DefineAuthChallenge *string `field:"optional" json:"defineAuthChallenge" yaml:"defineAuthChallenge"`
	// The Amazon Resource Name of a AWS Key Management Service ( AWS KMS ) key.
	//
	// Amazon Cognito uses the key to encrypt codes and temporary passwords sent to `CustomEmailSender` and `CustomSMSSender` .
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// A post-authentication AWS Lambda trigger.
	PostAuthentication *string `field:"optional" json:"postAuthentication" yaml:"postAuthentication"`
	// A post-confirmation AWS Lambda trigger.
	PostConfirmation *string `field:"optional" json:"postConfirmation" yaml:"postConfirmation"`
	// A pre-authentication AWS Lambda trigger.
	PreAuthentication *string `field:"optional" json:"preAuthentication" yaml:"preAuthentication"`
	// A pre-registration AWS Lambda trigger.
	PreSignUp *string `field:"optional" json:"preSignUp" yaml:"preSignUp"`
	// A Lambda trigger that is invoked before token generation.
	PreTokenGeneration *string `field:"optional" json:"preTokenGeneration" yaml:"preTokenGeneration"`
	// The user migration Lambda config type.
	UserMigration *string `field:"optional" json:"userMigration" yaml:"userMigration"`
	// Verifies the authentication challenge response.
	VerifyAuthChallengeResponse *string `field:"optional" json:"verifyAuthChallengeResponse" yaml:"verifyAuthChallengeResponse"`
}

