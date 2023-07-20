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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooldomain.html
//
type CfnUserPoolDomainProps struct {
	// The domain name for the domain that hosts the sign-up and sign-in pages for your application.
	//
	// For example: `auth.example.com` . If you're using a prefix domain, this field denotes the first part of the domain before `.auth.[region].amazoncognito.com` .
	//
	// This string can include only lowercase letters, numbers, and hyphens. Don't use a hyphen for the first or last character. Use periods to separate subdomain names.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooldomain.html#cfn-cognito-userpooldomain-domain
	//
	Domain *string `field:"required" json:"domain" yaml:"domain"`
	// The user pool ID for the user pool where you want to associate a user pool domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooldomain.html#cfn-cognito-userpooldomain-userpoolid
	//
	UserPoolId *string `field:"required" json:"userPoolId" yaml:"userPoolId"`
	// The configuration for a custom domain that hosts the sign-up and sign-in pages for your application.
	//
	// Use this object to specify an SSL certificate that is managed by ACM.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooldomain.html#cfn-cognito-userpooldomain-customdomainconfig
	//
	CustomDomainConfig interface{} `field:"optional" json:"customDomainConfig" yaml:"customDomainConfig"`
}

