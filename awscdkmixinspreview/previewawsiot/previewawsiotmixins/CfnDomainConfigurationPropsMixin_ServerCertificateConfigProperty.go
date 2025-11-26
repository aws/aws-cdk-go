package previewawsiotmixins


// The server certificate configuration.
//
// For more information, see [Configurable endpoints](https://docs.aws.amazon.com//iot/latest/developerguide/iot-custom-endpoints-configurable.html) from the AWS IoT Core Developer Guide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   serverCertificateConfigProperty := &ServerCertificateConfigProperty{
//   	EnableOcspCheck: jsii.Boolean(false),
//   	OcspAuthorizedResponderArn: jsii.String("ocspAuthorizedResponderArn"),
//   	OcspLambdaArn: jsii.String("ocspLambdaArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-domainconfiguration-servercertificateconfig.html
//
type CfnDomainConfigurationPropsMixin_ServerCertificateConfigProperty struct {
	// A Boolean value that indicates whether Online Certificate Status Protocol (OCSP) server certificate check is enabled or not.
	//
	// For more information, see [Configurable endpoints](https://docs.aws.amazon.com//iot/latest/developerguide/iot-custom-endpoints-configurable.html) from the AWS IoT Core Developer Guide.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-domainconfiguration-servercertificateconfig.html#cfn-iot-domainconfiguration-servercertificateconfig-enableocspcheck
	//
	EnableOcspCheck interface{} `field:"optional" json:"enableOcspCheck" yaml:"enableOcspCheck"`
	// The Amazon Resource Name (ARN) for an X.509 certificate stored in ACM. If provided, AWS IoT Core will use this certificate to validate the signature of the received OCSP response. The OCSP responder must sign responses using either this authorized responder certificate or the issuing certificate, depending on whether the ARN is provided or not. The certificate must be in the same account and region as the domain configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-domainconfiguration-servercertificateconfig.html#cfn-iot-domainconfiguration-servercertificateconfig-ocspauthorizedresponderarn
	//
	OcspAuthorizedResponderArn *string `field:"optional" json:"ocspAuthorizedResponderArn" yaml:"ocspAuthorizedResponderArn"`
	// The Amazon Resource Name (ARN) for a Lambda function that acts as a Request for Comments (RFC) 6960-compliant Online Certificate Status Protocol (OCSP) responder, supporting basic OCSP responses.
	//
	// The Lambda function accepts a base64-encoding of the OCSP request in the Distinguished Encoding Rules (DER) format. The Lambda function's response is also a base64-encoded OCSP response in the DER format. The response size must not exceed 4 kilobytes (KiB). The Lambda function must be in the same account and region as the domain configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-domainconfiguration-servercertificateconfig.html#cfn-iot-domainconfiguration-servercertificateconfig-ocsplambdaarn
	//
	OcspLambdaArn *string `field:"optional" json:"ocspLambdaArn" yaml:"ocspLambdaArn"`
}

