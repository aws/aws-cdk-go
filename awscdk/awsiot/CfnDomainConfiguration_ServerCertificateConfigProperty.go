package awsiot


// The server certificate configuration.
//
// For more information, see [Configurable endpoints](https://docs.aws.amazon.com//iot/latest/developerguide/iot-custom-endpoints-configurable.html) from the AWS IoT Core Developer Guide.
//
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
	// A Boolean value that indicates whether Online Certificate Status Protocol (OCSP) server certificate check is enabled or not.
	//
	// For more information, see [Configurable endpoints](https://docs.aws.amazon.com//iot/latest/developerguide/iot-custom-endpoints-configurable.html) from the AWS IoT Core Developer Guide.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-domainconfiguration-servercertificateconfig.html#cfn-iot-domainconfiguration-servercertificateconfig-enableocspcheck
	//
	EnableOcspCheck interface{} `field:"optional" json:"enableOcspCheck" yaml:"enableOcspCheck"`
}

