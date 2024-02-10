package awsiot

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDomainConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDomainConfigurationProps := &CfnDomainConfigurationProps{
//   	AuthorizerConfig: &AuthorizerConfigProperty{
//   		AllowAuthorizerOverride: jsii.Boolean(false),
//   		DefaultAuthorizerName: jsii.String("defaultAuthorizerName"),
//   	},
//   	DomainConfigurationName: jsii.String("domainConfigurationName"),
//   	DomainConfigurationStatus: jsii.String("domainConfigurationStatus"),
//   	DomainName: jsii.String("domainName"),
//   	ServerCertificateArns: []*string{
//   		jsii.String("serverCertificateArns"),
//   	},
//   	ServerCertificateConfig: &ServerCertificateConfigProperty{
//   		EnableOcspCheck: jsii.Boolean(false),
//   	},
//   	ServiceType: jsii.String("serviceType"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TlsConfig: &TlsConfigProperty{
//   		SecurityPolicy: jsii.String("securityPolicy"),
//   	},
//   	ValidationCertificateArn: jsii.String("validationCertificateArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-domainconfiguration.html
//
type CfnDomainConfigurationProps struct {
	// An object that specifies the authorization service for a domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-domainconfiguration.html#cfn-iot-domainconfiguration-authorizerconfig
	//
	AuthorizerConfig interface{} `field:"optional" json:"authorizerConfig" yaml:"authorizerConfig"`
	// The name of the domain configuration.
	//
	// This value must be unique to a region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-domainconfiguration.html#cfn-iot-domainconfiguration-domainconfigurationname
	//
	DomainConfigurationName *string `field:"optional" json:"domainConfigurationName" yaml:"domainConfigurationName"`
	// The status to which the domain configuration should be updated.
	//
	// Valid values: `ENABLED` | `DISABLED`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-domainconfiguration.html#cfn-iot-domainconfiguration-domainconfigurationstatus
	//
	DomainConfigurationStatus *string `field:"optional" json:"domainConfigurationStatus" yaml:"domainConfigurationStatus"`
	// The name of the domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-domainconfiguration.html#cfn-iot-domainconfiguration-domainname
	//
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// The ARNs of the certificates that AWS IoT passes to the device during the TLS handshake.
	//
	// Currently you can specify only one certificate ARN. This value is not required for AWS -managed domains.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-domainconfiguration.html#cfn-iot-domainconfiguration-servercertificatearns
	//
	ServerCertificateArns *[]*string `field:"optional" json:"serverCertificateArns" yaml:"serverCertificateArns"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-domainconfiguration.html#cfn-iot-domainconfiguration-servercertificateconfig
	//
	ServerCertificateConfig interface{} `field:"optional" json:"serverCertificateConfig" yaml:"serverCertificateConfig"`
	// The type of service delivered by the endpoint.
	//
	// > AWS IoT Core currently supports only the `DATA` service type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-domainconfiguration.html#cfn-iot-domainconfiguration-servicetype
	//
	ServiceType *string `field:"optional" json:"serviceType" yaml:"serviceType"`
	// Metadata which can be used to manage the domain configuration.
	//
	// > For URI Request parameters use format: ...key1=value1&key2=value2...
	// >
	// > For the CLI command-line parameter use format: &&tags "key1=value1&key2=value2..."
	// >
	// > For the cli-input-json file use format: "tags": "key1=value1&key2=value2..."
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-domainconfiguration.html#cfn-iot-domainconfiguration-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// An object that specifies the TLS configuration for a domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-domainconfiguration.html#cfn-iot-domainconfiguration-tlsconfig
	//
	TlsConfig interface{} `field:"optional" json:"tlsConfig" yaml:"tlsConfig"`
	// The certificate used to validate the server certificate and prove domain name ownership.
	//
	// This certificate must be signed by a public certificate authority. This value is not required for AWS -managed domains.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-domainconfiguration.html#cfn-iot-domainconfiguration-validationcertificatearn
	//
	ValidationCertificateArn *string `field:"optional" json:"validationCertificateArn" yaml:"validationCertificateArn"`
}

