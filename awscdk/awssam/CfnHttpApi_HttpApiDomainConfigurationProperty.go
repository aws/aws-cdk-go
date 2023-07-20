package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpApiDomainConfigurationProperty := &HttpApiDomainConfigurationProperty{
//   	CertificateArn: jsii.String("certificateArn"),
//   	DomainName: jsii.String("domainName"),
//
//   	// the properties below are optional
//   	BasePath: jsii.String("basePath"),
//   	EndpointConfiguration: jsii.String("endpointConfiguration"),
//   	MutualTlsAuthentication: &MutualTlsAuthenticationProperty{
//   		TruststoreUri: jsii.String("truststoreUri"),
//   		TruststoreVersion: jsii.Boolean(false),
//   	},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-httpapi-httpapidomainconfiguration.html
//
type CfnHttpApi_HttpApiDomainConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-httpapi-httpapidomainconfiguration.html#cfn-serverless-httpapi-httpapidomainconfiguration-certificatearn
	//
	CertificateArn *string `field:"required" json:"certificateArn" yaml:"certificateArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-httpapi-httpapidomainconfiguration.html#cfn-serverless-httpapi-httpapidomainconfiguration-domainname
	//
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-httpapi-httpapidomainconfiguration.html#cfn-serverless-httpapi-httpapidomainconfiguration-basepath
	//
	BasePath *string `field:"optional" json:"basePath" yaml:"basePath"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-httpapi-httpapidomainconfiguration.html#cfn-serverless-httpapi-httpapidomainconfiguration-endpointconfiguration
	//
	EndpointConfiguration *string `field:"optional" json:"endpointConfiguration" yaml:"endpointConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-httpapi-httpapidomainconfiguration.html#cfn-serverless-httpapi-httpapidomainconfiguration-mutualtlsauthentication
	//
	MutualTlsAuthentication interface{} `field:"optional" json:"mutualTlsAuthentication" yaml:"mutualTlsAuthentication"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-httpapi-httpapidomainconfiguration.html#cfn-serverless-httpapi-httpapidomainconfiguration-route53
	//
	Route53 interface{} `field:"optional" json:"route53" yaml:"route53"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-httpapi-httpapidomainconfiguration.html#cfn-serverless-httpapi-httpapidomainconfiguration-securitypolicy
	//
	SecurityPolicy *string `field:"optional" json:"securityPolicy" yaml:"securityPolicy"`
}

