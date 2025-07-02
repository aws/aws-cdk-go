package awssso


// A structure that describes the sign-in options for an application portal.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   signInOptionsProperty := &SignInOptionsProperty{
//   	Origin: jsii.String("origin"),
//
//   	// the properties below are optional
//   	ApplicationUrl: jsii.String("applicationUrl"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sso-application-signinoptions.html
//
type CfnApplication_SignInOptionsProperty struct {
	// This determines how IAM Identity Center navigates the user to the target application.
	//
	// It can be one of the following values:
	//
	// - `APPLICATION` : IAM Identity Center redirects the customer to the configured `ApplicationUrl` .
	// - `IDENTITY_CENTER` : IAM Identity Center uses SAML identity-provider initiated authentication to sign the customer directly into a SAML-based application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sso-application-signinoptions.html#cfn-sso-application-signinoptions-origin
	//
	Origin *string `field:"required" json:"origin" yaml:"origin"`
	// The URL that accepts authentication requests for an application.
	//
	// This is a required parameter if the `Origin` parameter is `APPLICATION` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sso-application-signinoptions.html#cfn-sso-application-signinoptions-applicationurl
	//
	ApplicationUrl *string `field:"optional" json:"applicationUrl" yaml:"applicationUrl"`
}

