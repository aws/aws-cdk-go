package awspipes


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   selfManagedKafkaAccessConfigurationCredentialsProperty := &selfManagedKafkaAccessConfigurationCredentialsProperty{
//   	basicAuth: jsii.String("basicAuth"),
//   	clientCertificateTlsAuth: jsii.String("clientCertificateTlsAuth"),
//   	saslScram256Auth: jsii.String("saslScram256Auth"),
//   	saslScram512Auth: jsii.String("saslScram512Auth"),
//   }
//
type CfnPipe_SelfManagedKafkaAccessConfigurationCredentialsProperty struct {
	// `CfnPipe.SelfManagedKafkaAccessConfigurationCredentialsProperty.BasicAuth`.
	BasicAuth *string `field:"optional" json:"basicAuth" yaml:"basicAuth"`
	// `CfnPipe.SelfManagedKafkaAccessConfigurationCredentialsProperty.ClientCertificateTlsAuth`.
	ClientCertificateTlsAuth *string `field:"optional" json:"clientCertificateTlsAuth" yaml:"clientCertificateTlsAuth"`
	// `CfnPipe.SelfManagedKafkaAccessConfigurationCredentialsProperty.SaslScram256Auth`.
	SaslScram256Auth *string `field:"optional" json:"saslScram256Auth" yaml:"saslScram256Auth"`
	// `CfnPipe.SelfManagedKafkaAccessConfigurationCredentialsProperty.SaslScram512Auth`.
	SaslScram512Auth *string `field:"optional" json:"saslScram512Auth" yaml:"saslScram512Auth"`
}

