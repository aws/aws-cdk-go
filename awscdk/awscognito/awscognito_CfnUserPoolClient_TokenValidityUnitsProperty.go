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
//   tokenValidityUnitsProperty := &tokenValidityUnitsProperty{
//   	accessToken: jsii.String("accessToken"),
//   	idToken: jsii.String("idToken"),
//   	refreshToken: jsii.String("refreshToken"),
//   }
//
type CfnUserPoolClient_TokenValidityUnitsProperty struct {
	// A time unit in “seconds”, “minutes”, “hours”, or “days” for the value in AccessTokenValidity, defaulting to hours.
	AccessToken *string `field:"optional" json:"accessToken" yaml:"accessToken"`
	// A time unit in “seconds”, “minutes”, “hours”, or “days” for the value in IdTokenValidity, defaulting to hours.
	IdToken *string `field:"optional" json:"idToken" yaml:"idToken"`
	// A time unit in “seconds”, “minutes”, “hours”, or “days” for the value in RefreshTokenValidity, defaulting to days.
	RefreshToken *string `field:"optional" json:"refreshToken" yaml:"refreshToken"`
}

