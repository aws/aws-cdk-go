package awscognito

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

// Triggers for a user pool.
//
// Example:
//   authChallengeFn := lambda.NewFunction(this, jsii.String("authChallengeFn"), &FunctionProps{
//   	Runtime: lambda.Runtime_NODEJS_LATEST(),
//   	Handler: jsii.String("index.handler"),
//   	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("path/to/asset"))),
//   })
//
//   userpool := cognito.NewUserPool(this, jsii.String("myuserpool"), &UserPoolProps{
//   	// ...
//   	LambdaTriggers: &UserPoolTriggers{
//   		CreateAuthChallenge: authChallengeFn,
//   	},
//   })
//
//   userpool.AddTrigger(cognito.UserPoolOperation_USER_MIGRATION(), lambda.NewFunction(this, jsii.String("userMigrationFn"), &FunctionProps{
//   	Runtime: lambda.Runtime_NODEJS_LATEST(),
//   	Handler: jsii.String("index.handler"),
//   	Code: lambda.Code_*FromAsset(path.join(__dirname, jsii.String("path/to/asset"))),
//   }))
//
// See: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-identity-pools-working-with-aws-lambda-triggers.html
//
type UserPoolTriggers struct {
	// Creates an authentication challenge.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-lambda-create-auth-challenge.html
	//
	// Default: - no trigger configured.
	//
	CreateAuthChallenge awslambda.IFunction `field:"optional" json:"createAuthChallenge" yaml:"createAuthChallenge"`
	// Amazon Cognito invokes this trigger to send email notifications to users.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-lambda-custom-email-sender.html
	//
	// Default: - no trigger configured.
	//
	CustomEmailSender awslambda.IFunction `field:"optional" json:"customEmailSender" yaml:"customEmailSender"`
	// A custom Message AWS Lambda trigger.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-lambda-custom-message.html
	//
	// Default: - no trigger configured.
	//
	CustomMessage awslambda.IFunction `field:"optional" json:"customMessage" yaml:"customMessage"`
	// Amazon Cognito invokes this trigger to send SMS notifications to users.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-lambda-custom-sms-sender.html
	//
	// Default: - no trigger configured.
	//
	CustomSmsSender awslambda.IFunction `field:"optional" json:"customSmsSender" yaml:"customSmsSender"`
	// Defines the authentication challenge.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-lambda-define-auth-challenge.html
	//
	// Default: - no trigger configured.
	//
	DefineAuthChallenge awslambda.IFunction `field:"optional" json:"defineAuthChallenge" yaml:"defineAuthChallenge"`
	// A post-authentication AWS Lambda trigger.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-lambda-post-authentication.html
	//
	// Default: - no trigger configured.
	//
	PostAuthentication awslambda.IFunction `field:"optional" json:"postAuthentication" yaml:"postAuthentication"`
	// A post-confirmation AWS Lambda trigger.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-lambda-post-confirmation.html
	//
	// Default: - no trigger configured.
	//
	PostConfirmation awslambda.IFunction `field:"optional" json:"postConfirmation" yaml:"postConfirmation"`
	// A pre-authentication AWS Lambda trigger.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-lambda-pre-authentication.html
	//
	// Default: - no trigger configured.
	//
	PreAuthentication awslambda.IFunction `field:"optional" json:"preAuthentication" yaml:"preAuthentication"`
	// A pre-registration AWS Lambda trigger.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-lambda-pre-sign-up.html
	//
	// Default: - no trigger configured.
	//
	PreSignUp awslambda.IFunction `field:"optional" json:"preSignUp" yaml:"preSignUp"`
	// A pre-token-generation AWS Lambda trigger.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-lambda-pre-token-generation.html
	//
	// Default: - no trigger configured.
	//
	PreTokenGeneration awslambda.IFunction `field:"optional" json:"preTokenGeneration" yaml:"preTokenGeneration"`
	// A user-migration AWS Lambda trigger.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-lambda-migrate-user.html
	//
	// Default: - no trigger configured.
	//
	UserMigration awslambda.IFunction `field:"optional" json:"userMigration" yaml:"userMigration"`
	// Verifies the authentication challenge response.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-lambda-verify-auth-challenge-response.html
	//
	// Default: - no trigger configured.
	//
	VerifyAuthChallengeResponse awslambda.IFunction `field:"optional" json:"verifyAuthChallengeResponse" yaml:"verifyAuthChallengeResponse"`
}

