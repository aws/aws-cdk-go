package awsiot

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnCACertificate`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCACertificateProps := &CfnCACertificateProps{
//   	CaCertificatePem: jsii.String("caCertificatePem"),
//   	Status: jsii.String("status"),
//
//   	// the properties below are optional
//   	AutoRegistrationStatus: jsii.String("autoRegistrationStatus"),
//   	CertificateMode: jsii.String("certificateMode"),
//   	RegistrationConfig: &RegistrationConfigProperty{
//   		RoleArn: jsii.String("roleArn"),
//   		TemplateBody: jsii.String("templateBody"),
//   		TemplateName: jsii.String("templateName"),
//   	},
//   	RemoveAutoRegistration: jsii.Boolean(false),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VerificationCertificatePem: jsii.String("verificationCertificatePem"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-cacertificate.html
//
type CfnCACertificateProps struct {
	// The certificate data in PEM format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-cacertificate.html#cfn-iot-cacertificate-cacertificatepem
	//
	CaCertificatePem *string `field:"required" json:"caCertificatePem" yaml:"caCertificatePem"`
	// The status of the CA certificate.
	//
	// Valid values are "ACTIVE" and "INACTIVE".
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-cacertificate.html#cfn-iot-cacertificate-status
	//
	Status *string `field:"required" json:"status" yaml:"status"`
	// Whether the CA certificate is configured for auto registration of device certificates.
	//
	// Valid values are "ENABLE" and "DISABLE".
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-cacertificate.html#cfn-iot-cacertificate-autoregistrationstatus
	//
	AutoRegistrationStatus *string `field:"optional" json:"autoRegistrationStatus" yaml:"autoRegistrationStatus"`
	// The mode of the CA.
	//
	// All the device certificates that are registered using this CA will be registered in the same mode as the CA. For more information about certificate mode for device certificates, see [certificate mode](https://docs.aws.amazon.com//iot/latest/apireference/API_CertificateDescription.html#iot-Type-CertificateDescription-certificateMode) .
	//
	// Valid values are "DEFAULT" and "SNI_ONLY".
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-cacertificate.html#cfn-iot-cacertificate-certificatemode
	//
	CertificateMode *string `field:"optional" json:"certificateMode" yaml:"certificateMode"`
	// Information about the registration configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-cacertificate.html#cfn-iot-cacertificate-registrationconfig
	//
	RegistrationConfig interface{} `field:"optional" json:"registrationConfig" yaml:"registrationConfig"`
	// If true, removes auto registration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-cacertificate.html#cfn-iot-cacertificate-removeautoregistration
	//
	RemoveAutoRegistration interface{} `field:"optional" json:"removeAutoRegistration" yaml:"removeAutoRegistration"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-cacertificate.html#cfn-iot-cacertificate-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The private key verification certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-cacertificate.html#cfn-iot-cacertificate-verificationcertificatepem
	//
	VerificationCertificatePem *string `field:"optional" json:"verificationCertificatePem" yaml:"verificationCertificatePem"`
}

