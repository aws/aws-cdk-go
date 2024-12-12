package awscognito


// Properties for defining a `CfnUserPoolDomain`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnUserPoolDomainProps := &CfnUserPoolDomainProps{
//   	Domain: jsii.String("domain"),
//   	UserPoolId: jsii.String("userPoolId"),
//
//   	// the properties below are optional
//   	CustomDomainConfig: &CustomDomainConfigTypeProperty{
//   		CertificateArn: jsii.String("certificateArn"),
//   	},
//   	ManagedLoginVersion: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooldomain.html
//
type CfnUserPoolDomainProps struct {
	// The domain name for the custom domain that hosts the sign-up and sign-in pages for your application.
	//
	// One example might be `auth.example.com` .
	//
	// This string can include only lowercase letters, numbers, and hyphens. Don't use a hyphen for the first or last character. Use periods to separate subdomain names.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooldomain.html#cfn-cognito-userpooldomain-domain
	//
	Domain *string `field:"required" json:"domain" yaml:"domain"`
	// The ID of the user pool that is associated with the custom domain whose certificate you're updating.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooldomain.html#cfn-cognito-userpooldomain-userpoolid
	//
	UserPoolId *string `field:"required" json:"userPoolId" yaml:"userPoolId"`
	// The configuration for a custom domain that hosts the sign-up and sign-in pages for your application.
	//
	// Use this object to specify an SSL certificate that is managed by ACM.
	//
	// When you create a custom domain, the passkey RP ID defaults to the custom domain. If you had a prefix domain active, this will cause passkey integration for your prefix domain to stop working due to a mismatch in RP ID. To keep the prefix domain passkey integration working, you can explicitly set RP ID to the prefix domain. Update the RP ID in a [SetUserPoolMfaConfig](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_SetUserPoolMfaConfig.html) request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooldomain.html#cfn-cognito-userpooldomain-customdomainconfig
	//
	CustomDomainConfig interface{} `field:"optional" json:"customDomainConfig" yaml:"customDomainConfig"`
	// A version number that indicates the state of managed login for your domain.
	//
	// Version `1` is hosted UI (classic). Version `2` is the newer managed login with the branding designer. For more information, see [Managed login](https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-managed-login.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooldomain.html#cfn-cognito-userpooldomain-managedloginversion
	//
	ManagedLoginVersion *float64 `field:"optional" json:"managedLoginVersion" yaml:"managedLoginVersion"`
}

