package awsiot


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   registrationConfigProperty := &registrationConfigProperty{
//   	roleArn: jsii.String("roleArn"),
//   	templateBody: jsii.String("templateBody"),
//   	templateName: jsii.String("templateName"),
//   }
//
type CfnCACertificate_RegistrationConfigProperty struct {
	// `CfnCACertificate.RegistrationConfigProperty.RoleArn`.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// `CfnCACertificate.RegistrationConfigProperty.TemplateBody`.
	TemplateBody *string `field:"optional" json:"templateBody" yaml:"templateBody"`
	// `CfnCACertificate.RegistrationConfigProperty.TemplateName`.
	TemplateName *string `field:"optional" json:"templateName" yaml:"templateName"`
}

