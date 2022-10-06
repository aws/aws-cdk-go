package awsiot

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnCACertificate`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCACertificateProps := &cfnCACertificateProps{
//   	caCertificatePem: jsii.String("caCertificatePem"),
//   	status: jsii.String("status"),
//
//   	// the properties below are optional
//   	autoRegistrationStatus: jsii.String("autoRegistrationStatus"),
//   	certificateMode: jsii.String("certificateMode"),
//   	registrationConfig: &registrationConfigProperty{
//   		roleArn: jsii.String("roleArn"),
//   		templateBody: jsii.String("templateBody"),
//   		templateName: jsii.String("templateName"),
//   	},
//   	removeAutoRegistration: jsii.Boolean(false),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	verificationCertificatePem: jsii.String("verificationCertificatePem"),
//   }
//
type CfnCACertificateProps struct {
	// `AWS::IoT::CACertificate.CACertificatePem`.
	CaCertificatePem *string `field:"required" json:"caCertificatePem" yaml:"caCertificatePem"`
	// `AWS::IoT::CACertificate.Status`.
	Status *string `field:"required" json:"status" yaml:"status"`
	// `AWS::IoT::CACertificate.AutoRegistrationStatus`.
	AutoRegistrationStatus *string `field:"optional" json:"autoRegistrationStatus" yaml:"autoRegistrationStatus"`
	// `AWS::IoT::CACertificate.CertificateMode`.
	CertificateMode *string `field:"optional" json:"certificateMode" yaml:"certificateMode"`
	// `AWS::IoT::CACertificate.RegistrationConfig`.
	RegistrationConfig interface{} `field:"optional" json:"registrationConfig" yaml:"registrationConfig"`
	// `AWS::IoT::CACertificate.RemoveAutoRegistration`.
	RemoveAutoRegistration interface{} `field:"optional" json:"removeAutoRegistration" yaml:"removeAutoRegistration"`
	// `AWS::IoT::CACertificate.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// `AWS::IoT::CACertificate.VerificationCertificatePem`.
	VerificationCertificatePem *string `field:"optional" json:"verificationCertificatePem" yaml:"verificationCertificatePem"`
}

