package mixinsawscognito


// The configuration of your app client for refresh token rotation.
//
// When enabled, your app client issues new ID, access, and refresh tokens when users renew their sessions with refresh tokens. When disabled, token refresh issues only ID and access tokens.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   refreshTokenRotationProperty := &RefreshTokenRotationProperty{
//   	Feature: jsii.String("feature"),
//   	RetryGracePeriodSeconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolclient-refreshtokenrotation.html
//
type CfnUserPoolClientPropsMixin_RefreshTokenRotationProperty struct {
	// The state of refresh token rotation for the current app client.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolclient-refreshtokenrotation.html#cfn-cognito-userpoolclient-refreshtokenrotation-feature
	//
	Feature *string `field:"optional" json:"feature" yaml:"feature"`
	// When you request a token refresh with `GetTokensFromRefreshToken` , the original refresh token that you're rotating out can remain valid for a period of time of up to 60 seconds.
	//
	// This allows for client-side retries. When `RetryGracePeriodSeconds` is `0` , the grace period is disabled and a successful request immediately invalidates the submitted refresh token.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolclient-refreshtokenrotation.html#cfn-cognito-userpoolclient-refreshtokenrotation-retrygraceperiodseconds
	//
	RetryGracePeriodSeconds *float64 `field:"optional" json:"retryGracePeriodSeconds" yaml:"retryGracePeriodSeconds"`
}

