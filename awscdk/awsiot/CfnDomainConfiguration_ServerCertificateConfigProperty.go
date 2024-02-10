package awsiot


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serverCertificateConfigProperty := &ServerCertificateConfigProperty{
//   	EnableOcspCheck: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-domainconfiguration-servercertificateconfig.html
//
type CfnDomainConfiguration_ServerCertificateConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-domainconfiguration-servercertificateconfig.html#cfn-iot-domainconfiguration-servercertificateconfig-enableocspcheck
	//
	EnableOcspCheck interface{} `field:"optional" json:"enableOcspCheck" yaml:"enableOcspCheck"`
}

