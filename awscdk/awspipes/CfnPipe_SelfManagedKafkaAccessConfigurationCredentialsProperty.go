package awspipes


// The AWS Secrets Manager secret that stores your stream credentials.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   selfManagedKafkaAccessConfigurationCredentialsProperty := &SelfManagedKafkaAccessConfigurationCredentialsProperty{
//   	BasicAuth: jsii.String("basicAuth"),
//   	ClientCertificateTlsAuth: jsii.String("clientCertificateTlsAuth"),
//   	SaslScram256Auth: jsii.String("saslScram256Auth"),
//   	SaslScram512Auth: jsii.String("saslScram512Auth"),
//   }
//
type CfnPipe_SelfManagedKafkaAccessConfigurationCredentialsProperty struct {
	// The ARN of the Secrets Manager secret.
	BasicAuth *string `field:"optional" json:"basicAuth" yaml:"basicAuth"`
	// The ARN of the Secrets Manager secret.
	ClientCertificateTlsAuth *string `field:"optional" json:"clientCertificateTlsAuth" yaml:"clientCertificateTlsAuth"`
	// The ARN of the Secrets Manager secret.
	SaslScram256Auth *string `field:"optional" json:"saslScram256Auth" yaml:"saslScram256Auth"`
	// The ARN of the Secrets Manager secret.
	SaslScram512Auth *string `field:"optional" json:"saslScram512Auth" yaml:"saslScram512Auth"`
}

