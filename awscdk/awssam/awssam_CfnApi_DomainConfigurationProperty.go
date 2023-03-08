package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   domainConfigurationProperty := &domainConfigurationProperty{
//   	certificateArn: jsii.String("certificateArn"),
//   	domainName: jsii.String("domainName"),
//
//   	// the properties below are optional
//   	basePath: []*string{
//   		jsii.String("basePath"),
//   	},
//   	endpointConfiguration: jsii.String("endpointConfiguration"),
//   	mutualTlsAuthentication: &mutualTlsAuthenticationProperty{
//   		truststoreUri: jsii.String("truststoreUri"),
//   		truststoreVersion: jsii.String("truststoreVersion"),
//   	},
//   	ownershipVerificationCertificateArn: jsii.String("ownershipVerificationCertificateArn"),
//   	route53: &route53ConfigurationProperty{
//   		distributedDomainName: jsii.String("distributedDomainName"),
//   		evaluateTargetHealth: jsii.Boolean(false),
//   		hostedZoneId: jsii.String("hostedZoneId"),
//   		hostedZoneName: jsii.String("hostedZoneName"),
//   		ipV6: jsii.Boolean(false),
//   	},
//   	securityPolicy: jsii.String("securityPolicy"),
//   }
//
type CfnApi_DomainConfigurationProperty struct {
	// `CfnApi.DomainConfigurationProperty.CertificateArn`.
	CertificateArn *string `field:"required" json:"certificateArn" yaml:"certificateArn"`
	// `CfnApi.DomainConfigurationProperty.DomainName`.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// `CfnApi.DomainConfigurationProperty.BasePath`.
	BasePath *[]*string `field:"optional" json:"basePath" yaml:"basePath"`
	// `CfnApi.DomainConfigurationProperty.EndpointConfiguration`.
	EndpointConfiguration *string `field:"optional" json:"endpointConfiguration" yaml:"endpointConfiguration"`
	// `CfnApi.DomainConfigurationProperty.MutualTlsAuthentication`.
	MutualTlsAuthentication interface{} `field:"optional" json:"mutualTlsAuthentication" yaml:"mutualTlsAuthentication"`
	// `CfnApi.DomainConfigurationProperty.OwnershipVerificationCertificateArn`.
	OwnershipVerificationCertificateArn *string `field:"optional" json:"ownershipVerificationCertificateArn" yaml:"ownershipVerificationCertificateArn"`
	// `CfnApi.DomainConfigurationProperty.Route53`.
	Route53 interface{} `field:"optional" json:"route53" yaml:"route53"`
	// `CfnApi.DomainConfigurationProperty.SecurityPolicy`.
	SecurityPolicy *string `field:"optional" json:"securityPolicy" yaml:"securityPolicy"`
}

