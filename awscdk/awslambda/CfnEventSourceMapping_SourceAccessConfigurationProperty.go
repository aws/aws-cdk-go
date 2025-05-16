package awslambda


// An array of the authentication protocol, VPC components, or virtual host to secure and define your event source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceAccessConfigurationProperty := &SourceAccessConfigurationProperty{
//   	Type: jsii.String("type"),
//   	Uri: jsii.String("uri"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventsourcemapping-sourceaccessconfiguration.html
//
type CfnEventSourceMapping_SourceAccessConfigurationProperty struct {
	// The type of authentication protocol, VPC components, or virtual host for your event source. For example: `"Type":"SASL_SCRAM_512_AUTH"` .
	//
	// - `BASIC_AUTH` – (Amazon MQ) The AWS Secrets Manager secret that stores your broker credentials.
	// - `BASIC_AUTH` – (Self-managed Apache Kafka) The Secrets Manager ARN of your secret key used for SASL/PLAIN authentication of your Apache Kafka brokers.
	// - `VPC_SUBNET` – (Self-managed Apache Kafka) The subnets associated with your VPC. Lambda connects to these subnets to fetch data from your self-managed Apache Kafka cluster.
	// - `VPC_SECURITY_GROUP` – (Self-managed Apache Kafka) The VPC security group used to manage access to your self-managed Apache Kafka brokers.
	// - `SASL_SCRAM_256_AUTH` – (Self-managed Apache Kafka) The Secrets Manager ARN of your secret key used for SASL SCRAM-256 authentication of your self-managed Apache Kafka brokers.
	// - `SASL_SCRAM_512_AUTH` – (Amazon MSK, Self-managed Apache Kafka) The Secrets Manager ARN of your secret key used for SASL SCRAM-512 authentication of your self-managed Apache Kafka brokers.
	// - `VIRTUAL_HOST` –- (RabbitMQ) The name of the virtual host in your RabbitMQ broker. Lambda uses this RabbitMQ host as the event source. This property cannot be specified in an UpdateEventSourceMapping API call.
	// - `CLIENT_CERTIFICATE_TLS_AUTH` – (Amazon MSK, self-managed Apache Kafka) The Secrets Manager ARN of your secret key containing the certificate chain (X.509 PEM), private key (PKCS#8 PEM), and private key password (optional) used for mutual TLS authentication of your MSK/Apache Kafka brokers.
	// - `SERVER_ROOT_CA_CERTIFICATE` – (Self-managed Apache Kafka) The Secrets Manager ARN of your secret key containing the root CA certificate (X.509 PEM) used for TLS encryption of your Apache Kafka brokers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventsourcemapping-sourceaccessconfiguration.html#cfn-lambda-eventsourcemapping-sourceaccessconfiguration-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The value for your chosen configuration in `Type` .
	//
	// For example: `"URI": "arn:aws:secretsmanager:us-east-1:01234567890:secret:MyBrokerSecretName"` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventsourcemapping-sourceaccessconfiguration.html#cfn-lambda-eventsourcemapping-sourceaccessconfiguration-uri
	//
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

