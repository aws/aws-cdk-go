package awsapigatewayv2


// Properties for defining a `CfnDomainName`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDomainNameProps := &CfnDomainNameProps{
//   	DomainName: jsii.String("domainName"),
//
//   	// the properties below are optional
//   	DomainNameConfigurations: []interface{}{
//   		&DomainNameConfigurationProperty{
//   			CertificateArn: jsii.String("certificateArn"),
//   			CertificateName: jsii.String("certificateName"),
//   			EndpointType: jsii.String("endpointType"),
//   			OwnershipVerificationCertificateArn: jsii.String("ownershipVerificationCertificateArn"),
//   			SecurityPolicy: jsii.String("securityPolicy"),
//   		},
//   	},
//   	MutualTlsAuthentication: &MutualTlsAuthenticationProperty{
//   		TruststoreUri: jsii.String("truststoreUri"),
//   		TruststoreVersion: jsii.String("truststoreVersion"),
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-domainname.html
//
type CfnDomainNameProps struct {
	// The custom domain name for your API in Amazon API Gateway.
	//
	// Uppercase letters are not supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-domainname.html#cfn-apigatewayv2-domainname-domainname
	//
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The domain name configurations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-domainname.html#cfn-apigatewayv2-domainname-domainnameconfigurations
	//
	DomainNameConfigurations interface{} `field:"optional" json:"domainNameConfigurations" yaml:"domainNameConfigurations"`
	// The mutual TLS authentication configuration for a custom domain name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-domainname.html#cfn-apigatewayv2-domainname-mutualtlsauthentication
	//
	MutualTlsAuthentication interface{} `field:"optional" json:"mutualTlsAuthentication" yaml:"mutualTlsAuthentication"`
	// The collection of tags associated with a domain name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-domainname.html#cfn-apigatewayv2-domainname-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

