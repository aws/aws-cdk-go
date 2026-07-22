package awscognito


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
//   	InboundFederation: &InboundFederationProperty{
//   		LambdaArn: jsii.String("lambdaArn"),
//   		LambdaVersion: jsii.String("lambdaVersion"),
//   	},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolregionalconfigurationattachment-lambdaconfig.html
//
type CfnUserPoolRegionalConfigurationAttachment_LambdaConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolregionalconfigurationattachment-lambdaconfig.html#cfn-cognito-userpoolregionalconfigurationattachment-lambdaconfig-createauthchallenge
	//
	CreateAuthChallenge *string `field:"optional" json:"createAuthChallenge" yaml:"createAuthChallenge"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolregionalconfigurationattachment-lambdaconfig.html#cfn-cognito-userpoolregionalconfigurationattachment-lambdaconfig-customemailsender
	//
	CustomEmailSender interface{} `field:"optional" json:"customEmailSender" yaml:"customEmailSender"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolregionalconfigurationattachment-lambdaconfig.html#cfn-cognito-userpoolregionalconfigurationattachment-lambdaconfig-custommessage
	//
	CustomMessage *string `field:"optional" json:"customMessage" yaml:"customMessage"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolregionalconfigurationattachment-lambdaconfig.html#cfn-cognito-userpoolregionalconfigurationattachment-lambdaconfig-customsmssender
	//
	CustomSmsSender interface{} `field:"optional" json:"customSmsSender" yaml:"customSmsSender"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolregionalconfigurationattachment-lambdaconfig.html#cfn-cognito-userpoolregionalconfigurationattachment-lambdaconfig-defineauthchallenge
	//
	DefineAuthChallenge *string `field:"optional" json:"defineAuthChallenge" yaml:"defineAuthChallenge"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolregionalconfigurationattachment-lambdaconfig.html#cfn-cognito-userpoolregionalconfigurationattachment-lambdaconfig-inboundfederation
	//
	InboundFederation interface{} `field:"optional" json:"inboundFederation" yaml:"inboundFederation"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolregionalconfigurationattachment-lambdaconfig.html#cfn-cognito-userpoolregionalconfigurationattachment-lambdaconfig-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolregionalconfigurationattachment-lambdaconfig.html#cfn-cognito-userpoolregionalconfigurationattachment-lambdaconfig-postauthentication
	//
	PostAuthentication *string `field:"optional" json:"postAuthentication" yaml:"postAuthentication"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolregionalconfigurationattachment-lambdaconfig.html#cfn-cognito-userpoolregionalconfigurationattachment-lambdaconfig-postconfirmation
	//
	PostConfirmation *string `field:"optional" json:"postConfirmation" yaml:"postConfirmation"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolregionalconfigurationattachment-lambdaconfig.html#cfn-cognito-userpoolregionalconfigurationattachment-lambdaconfig-preauthentication
	//
	PreAuthentication *string `field:"optional" json:"preAuthentication" yaml:"preAuthentication"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolregionalconfigurationattachment-lambdaconfig.html#cfn-cognito-userpoolregionalconfigurationattachment-lambdaconfig-presignup
	//
	PreSignUp *string `field:"optional" json:"preSignUp" yaml:"preSignUp"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolregionalconfigurationattachment-lambdaconfig.html#cfn-cognito-userpoolregionalconfigurationattachment-lambdaconfig-pretokengeneration
	//
	PreTokenGeneration *string `field:"optional" json:"preTokenGeneration" yaml:"preTokenGeneration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolregionalconfigurationattachment-lambdaconfig.html#cfn-cognito-userpoolregionalconfigurationattachment-lambdaconfig-pretokengenerationconfig
	//
	PreTokenGenerationConfig interface{} `field:"optional" json:"preTokenGenerationConfig" yaml:"preTokenGenerationConfig"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolregionalconfigurationattachment-lambdaconfig.html#cfn-cognito-userpoolregionalconfigurationattachment-lambdaconfig-usermigration
	//
	UserMigration *string `field:"optional" json:"userMigration" yaml:"userMigration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolregionalconfigurationattachment-lambdaconfig.html#cfn-cognito-userpoolregionalconfigurationattachment-lambdaconfig-verifyauthchallengeresponse
	//
	VerifyAuthChallengeResponse *string `field:"optional" json:"verifyAuthChallengeResponse" yaml:"verifyAuthChallengeResponse"`
}

