package awscognito


// The time units you use when you set the duration of ID, access, and refresh tokens.
//
// The default unit for RefreshToken is days, and the default for ID and access tokens is hours.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tokenValidityUnitsProperty := &TokenValidityUnitsProperty{
//   	AccessToken: jsii.String("accessToken"),
//   	IdToken: jsii.String("idToken"),
//   	RefreshToken: jsii.String("refreshToken"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolclient-tokenvalidityunits.html
//
type CfnUserPoolClient_TokenValidityUnitsProperty struct {
	// A time unit of `seconds` , `minutes` , `hours` , or `days` for the value that you set in the `AccessTokenValidity` parameter.
	//
	// The default `AccessTokenValidity` time unit is hours. `AccessTokenValidity` duration can range from five minutes to one day.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolclient-tokenvalidityunits.html#cfn-cognito-userpoolclient-tokenvalidityunits-accesstoken
	//
	AccessToken *string `field:"optional" json:"accessToken" yaml:"accessToken"`
	// A time unit of `seconds` , `minutes` , `hours` , or `days` for the value that you set in the `IdTokenValidity` parameter.
	//
	// The default `IdTokenValidity` time unit is hours. `IdTokenValidity` duration can range from five minutes to one day.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolclient-tokenvalidityunits.html#cfn-cognito-userpoolclient-tokenvalidityunits-idtoken
	//
	IdToken *string `field:"optional" json:"idToken" yaml:"idToken"`
	// A time unit of `seconds` , `minutes` , `hours` , or `days` for the value that you set in the `RefreshTokenValidity` parameter.
	//
	// The default `RefreshTokenValidity` time unit is days. `RefreshTokenValidity` duration can range from 60 minutes to 10 years.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolclient-tokenvalidityunits.html#cfn-cognito-userpoolclient-tokenvalidityunits-refreshtoken
	//
	RefreshToken *string `field:"optional" json:"refreshToken" yaml:"refreshToken"`
}

