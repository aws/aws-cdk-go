package awspipes


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mQBrokerAccessCredentialsProperty := &mQBrokerAccessCredentialsProperty{
//   	basicAuth: jsii.String("basicAuth"),
//   }
//
type CfnPipe_MQBrokerAccessCredentialsProperty struct {
	// `CfnPipe.MQBrokerAccessCredentialsProperty.BasicAuth`.
	BasicAuth *string `field:"required" json:"basicAuth" yaml:"basicAuth"`
}

