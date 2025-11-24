package mixinsawsiot


// An object that contains information about a server certificate.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   serverCertificateSummaryProperty := &ServerCertificateSummaryProperty{
//   	ServerCertificateArn: jsii.String("serverCertificateArn"),
//   	ServerCertificateStatus: jsii.String("serverCertificateStatus"),
//   	ServerCertificateStatusDetail: jsii.String("serverCertificateStatusDetail"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-domainconfiguration-servercertificatesummary.html
//
type CfnDomainConfigurationPropsMixin_ServerCertificateSummaryProperty struct {
	// The ARN of the server certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-domainconfiguration-servercertificatesummary.html#cfn-iot-domainconfiguration-servercertificatesummary-servercertificatearn
	//
	ServerCertificateArn *string `field:"optional" json:"serverCertificateArn" yaml:"serverCertificateArn"`
	// The status of the server certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-domainconfiguration-servercertificatesummary.html#cfn-iot-domainconfiguration-servercertificatesummary-servercertificatestatus
	//
	ServerCertificateStatus *string `field:"optional" json:"serverCertificateStatus" yaml:"serverCertificateStatus"`
	// Details that explain the status of the server certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-domainconfiguration-servercertificatesummary.html#cfn-iot-domainconfiguration-servercertificatesummary-servercertificatestatusdetail
	//
	ServerCertificateStatusDetail *string `field:"optional" json:"serverCertificateStatusDetail" yaml:"serverCertificateStatusDetail"`
}

