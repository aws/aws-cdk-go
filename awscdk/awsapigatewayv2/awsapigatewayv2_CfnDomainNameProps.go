package awsapigatewayv2


// Properties for defining a `CfnDomainName`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tags interface{}
//
//   cfnDomainNameProps := &cfnDomainNameProps{
//   	domainName: jsii.String("domainName"),
//
//   	// the properties below are optional
//   	domainNameConfigurations: []interface{}{
//   		&domainNameConfigurationProperty{
//   			certificateArn: jsii.String("certificateArn"),
//   			certificateName: jsii.String("certificateName"),
//   			endpointType: jsii.String("endpointType"),
//   			ownershipVerificationCertificateArn: jsii.String("ownershipVerificationCertificateArn"),
//   			securityPolicy: jsii.String("securityPolicy"),
//   		},
//   	},
//   	mutualTlsAuthentication: &mutualTlsAuthenticationProperty{
//   		truststoreUri: jsii.String("truststoreUri"),
//   		truststoreVersion: jsii.String("truststoreVersion"),
//   	},
//   	tags: tags,
//   }
//
type CfnDomainNameProps struct {
	// The custom domain name for your API in Amazon API Gateway.
	//
	// Uppercase letters are not supported.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The domain name configurations.
	DomainNameConfigurations interface{} `field:"optional" json:"domainNameConfigurations" yaml:"domainNameConfigurations"`
	// The mutual TLS authentication configuration for a custom domain name.
	MutualTlsAuthentication interface{} `field:"optional" json:"mutualTlsAuthentication" yaml:"mutualTlsAuthentication"`
	// The collection of tags associated with a domain name.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

