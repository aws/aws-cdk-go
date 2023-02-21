package awscognito


// The units in which the validity times are represented.
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
type CfnUserPoolClient_TokenValidityUnitsProperty struct {
	// A time unit of `seconds` , `minutes` , `hours` , or `days` for the value that you set in the `AccessTokenValidity` parameter.
	//
	// The default `AccessTokenValidity` time unit is hours.
	AccessToken *string `field:"optional" json:"accessToken" yaml:"accessToken"`
	// A time unit of `seconds` , `minutes` , `hours` , or `days` for the value that you set in the `IdTokenValidity` parameter.
	//
	// The default `IdTokenValidity` time unit is hours.
	IdToken *string `field:"optional" json:"idToken" yaml:"idToken"`
	// A time unit of `seconds` , `minutes` , `hours` , or `days` for the value that you set in the `RefreshTokenValidity` parameter.
	//
	// The default `RefreshTokenValidity` time unit is days.
	RefreshToken *string `field:"optional" json:"refreshToken" yaml:"refreshToken"`
}

