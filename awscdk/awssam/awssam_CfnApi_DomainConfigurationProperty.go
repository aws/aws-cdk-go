package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   domainConfigurationProperty := &DomainConfigurationProperty{
//   	CertificateArn: jsii.String("certificateArn"),
//   	DomainName: jsii.String("domainName"),
//
//   	// the properties below are optional
//   	BasePath: []*string{
//   		jsii.String("basePath"),
//   	},
//   	EndpointConfiguration: jsii.String("endpointConfiguration"),
//   	MutualTlsAuthentication: &MutualTlsAuthenticationProperty{
//   		TruststoreUri: jsii.String("truststoreUri"),
//   		TruststoreVersion: jsii.String("truststoreVersion"),
//   	},
//   	OwnershipVerificationCertificateArn: jsii.String("ownershipVerificationCertificateArn"),
//   	Route53: &Route53ConfigurationProperty{
//   		DistributedDomainName: jsii.String("distributedDomainName"),
//   		EvaluateTargetHealth: jsii.Boolean(false),
//   		HostedZoneId: jsii.String("hostedZoneId"),
//   		HostedZoneName: jsii.String("hostedZoneName"),
//   		IpV6: jsii.Boolean(false),
//   	},
//   	SecurityPolicy: jsii.String("securityPolicy"),
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

