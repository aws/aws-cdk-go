package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpApiDomainConfigurationProperty := &httpApiDomainConfigurationProperty{
//   	certificateArn: jsii.String("certificateArn"),
//   	domainName: jsii.String("domainName"),
//
//   	// the properties below are optional
//   	basePath: jsii.String("basePath"),
//   	endpointConfiguration: jsii.String("endpointConfiguration"),
//   	mutualTlsAuthentication: &mutualTlsAuthenticationProperty{
//   		truststoreUri: jsii.String("truststoreUri"),
//   		truststoreVersion: jsii.Boolean(false),
//   	},
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
type CfnHttpApi_HttpApiDomainConfigurationProperty struct {
	// `CfnHttpApi.HttpApiDomainConfigurationProperty.CertificateArn`.
	CertificateArn *string `field:"required" json:"certificateArn" yaml:"certificateArn"`
	// `CfnHttpApi.HttpApiDomainConfigurationProperty.DomainName`.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// `CfnHttpApi.HttpApiDomainConfigurationProperty.BasePath`.
	BasePath *string `field:"optional" json:"basePath" yaml:"basePath"`
	// `CfnHttpApi.HttpApiDomainConfigurationProperty.EndpointConfiguration`.
	EndpointConfiguration *string `field:"optional" json:"endpointConfiguration" yaml:"endpointConfiguration"`
	// `CfnHttpApi.HttpApiDomainConfigurationProperty.MutualTlsAuthentication`.
	MutualTlsAuthentication interface{} `field:"optional" json:"mutualTlsAuthentication" yaml:"mutualTlsAuthentication"`
	// `CfnHttpApi.HttpApiDomainConfigurationProperty.Route53`.
	Route53 interface{} `field:"optional" json:"route53" yaml:"route53"`
	// `CfnHttpApi.HttpApiDomainConfigurationProperty.SecurityPolicy`.
	SecurityPolicy *string `field:"optional" json:"securityPolicy" yaml:"securityPolicy"`
}

