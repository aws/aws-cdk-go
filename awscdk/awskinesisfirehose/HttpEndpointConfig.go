package awskinesisfirehose

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

// Describes the configuration of the Http endpoint to which Kinesis Firehose delivers data.
//
// Example:
//   var endpointConfig HttpEndpointConfig
//
//   httpDestination := firehose.NewHttpEndpoint(&HttpEndpointProps{
//   	EndpointConfig: EndpointConfig,
//   })
//
type HttpEndpointConfig struct {
	// The URL of the Http endpoint selected as the destination.
	Url *string `field:"required" json:"url" yaml:"url"`
	// The access key used to authenticate with the Http endpoint.
	//
	// Used only when `secret` is not set. If both `accessKey` and `secret` are provided,
	// `secret` (AWS Secrets Manager) takes precedence and this value is ignored. The access
	// key is rendered into the CloudFormation template.
	// Default: - none; authentication uses `secret` if provided, otherwise no access key.
	//
	AccessKey awscdk.SecretValue `field:"optional" json:"accessKey" yaml:"accessKey"`
	// The name of the Http endpoint selected as the destination.
	// Default: - None.
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A Secrets Manager secret used to authenticate with the Http endpoint.
	//
	// When set, Firehose retrieves the credential from AWS Secrets Manager, and this takes
	// precedence over `accessKey`.
	// Default: - none; `accessKey` is used if provided.
	//
	Secret awssecretsmanager.ISecret `field:"optional" json:"secret" yaml:"secret"`
}

