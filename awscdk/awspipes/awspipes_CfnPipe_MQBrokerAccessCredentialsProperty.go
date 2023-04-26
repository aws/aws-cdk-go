package awspipes


// The AWS Secrets Manager secret that stores your broker credentials.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mQBrokerAccessCredentialsProperty := &MQBrokerAccessCredentialsProperty{
//   	BasicAuth: jsii.String("basicAuth"),
//   }
//
type CfnPipe_MQBrokerAccessCredentialsProperty struct {
	// The ARN of the Secrets Manager secret.
	BasicAuth *string `field:"required" json:"basicAuth" yaml:"basicAuth"`
}

