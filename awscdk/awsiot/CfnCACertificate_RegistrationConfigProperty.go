package awsiot


// The registration configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   registrationConfigProperty := &RegistrationConfigProperty{
//   	RoleArn: jsii.String("roleArn"),
//   	TemplateBody: jsii.String("templateBody"),
//   	TemplateName: jsii.String("templateName"),
//   }
//
type CfnCACertificate_RegistrationConfigProperty struct {
	// The ARN of the role.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The template body.
	TemplateBody *string `field:"optional" json:"templateBody" yaml:"templateBody"`
	// The name of the provisioning template.
	TemplateName *string `field:"optional" json:"templateName" yaml:"templateName"`
}

