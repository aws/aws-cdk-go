package previewawsiotmixins


// The registration configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   registrationConfigProperty := &RegistrationConfigProperty{
//   	RoleArn: jsii.String("roleArn"),
//   	TemplateBody: jsii.String("templateBody"),
//   	TemplateName: jsii.String("templateName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-cacertificate-registrationconfig.html
//
type CfnCACertificatePropsMixin_RegistrationConfigProperty struct {
	// The ARN of the role.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-cacertificate-registrationconfig.html#cfn-iot-cacertificate-registrationconfig-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The template body.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-cacertificate-registrationconfig.html#cfn-iot-cacertificate-registrationconfig-templatebody
	//
	TemplateBody *string `field:"optional" json:"templateBody" yaml:"templateBody"`
	// The name of the provisioning template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-cacertificate-registrationconfig.html#cfn-iot-cacertificate-registrationconfig-templatename
	//
	TemplateName *string `field:"optional" json:"templateName" yaml:"templateName"`
}

