package awscognito


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   refreshTokenRotationProperty := &RefreshTokenRotationProperty{
//   	Feature: jsii.String("feature"),
//   	RetryGracePeriodSeconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolclient-refreshtokenrotation.html
//
type CfnUserPoolClient_RefreshTokenRotationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolclient-refreshtokenrotation.html#cfn-cognito-userpoolclient-refreshtokenrotation-feature
	//
	Feature *string `field:"optional" json:"feature" yaml:"feature"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolclient-refreshtokenrotation.html#cfn-cognito-userpoolclient-refreshtokenrotation-retrygraceperiodseconds
	//
	RetryGracePeriodSeconds *float64 `field:"optional" json:"retryGracePeriodSeconds" yaml:"retryGracePeriodSeconds"`
}

