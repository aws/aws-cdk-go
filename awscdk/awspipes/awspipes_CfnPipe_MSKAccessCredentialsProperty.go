package awspipes


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mSKAccessCredentialsProperty := &mSKAccessCredentialsProperty{
//   	clientCertificateTlsAuth: jsii.String("clientCertificateTlsAuth"),
//   	saslScram512Auth: jsii.String("saslScram512Auth"),
//   }
//
type CfnPipe_MSKAccessCredentialsProperty struct {
	// `CfnPipe.MSKAccessCredentialsProperty.ClientCertificateTlsAuth`.
	ClientCertificateTlsAuth *string `field:"optional" json:"clientCertificateTlsAuth" yaml:"clientCertificateTlsAuth"`
	// `CfnPipe.MSKAccessCredentialsProperty.SaslScram512Auth`.
	SaslScram512Auth *string `field:"optional" json:"saslScram512Auth" yaml:"saslScram512Auth"`
}

