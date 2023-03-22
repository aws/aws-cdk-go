package awscognito


// Properties for defining a `CfnUserPoolDomain`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnUserPoolDomainProps := &cfnUserPoolDomainProps{
//   	domain: jsii.String("domain"),
//   	userPoolId: jsii.String("userPoolId"),
//
//   	// the properties below are optional
//   	customDomainConfig: &customDomainConfigTypeProperty{
//   		certificateArn: jsii.String("certificateArn"),
//   	},
//   }
//
type CfnUserPoolDomainProps struct {
	// The domain name for the domain that hosts the sign-up and sign-in pages for your application.
	//
	// For example: `auth.example.com` . If you're using a prefix domain, this field denotes the first part of the domain before `.auth.[region].amazoncognito.com` .
	//
	// This string can include only lowercase letters, numbers, and hyphens. Don't use a hyphen for the first or last character. Use periods to separate subdomain names.
	Domain *string `field:"required" json:"domain" yaml:"domain"`
	// The user pool ID for the user pool where you want to associate a user pool domain.
	UserPoolId *string `field:"required" json:"userPoolId" yaml:"userPoolId"`
	// The configuration for a custom domain that hosts the sign-up and sign-in pages for your application.
	//
	// Use this object to specify an SSL certificate that is managed by ACM.
	CustomDomainConfig interface{} `field:"optional" json:"customDomainConfig" yaml:"customDomainConfig"`
}

