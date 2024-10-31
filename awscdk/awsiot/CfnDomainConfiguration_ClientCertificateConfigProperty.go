package awsiot


// An object that speciﬁes the client certificate conﬁguration for a domain.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clientCertificateConfigProperty := &ClientCertificateConfigProperty{
//   	ClientCertificateCallbackArn: jsii.String("clientCertificateCallbackArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-domainconfiguration-clientcertificateconfig.html
//
type CfnDomainConfiguration_ClientCertificateConfigProperty struct {
	// The ARN of the Lambda function that IoT invokes after mutual TLS authentication during the connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-domainconfiguration-clientcertificateconfig.html#cfn-iot-domainconfiguration-clientcertificateconfig-clientcertificatecallbackarn
	//
	ClientCertificateCallbackArn *string `field:"optional" json:"clientCertificateCallbackArn" yaml:"clientCertificateCallbackArn"`
}

