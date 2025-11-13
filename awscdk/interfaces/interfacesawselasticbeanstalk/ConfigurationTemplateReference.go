package interfacesawselasticbeanstalk


// A reference to a ConfigurationTemplate resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   configurationTemplateReference := &ConfigurationTemplateReference{
//   	ApplicationName: jsii.String("applicationName"),
//   	TemplateName: jsii.String("templateName"),
//   }
//
type ConfigurationTemplateReference struct {
	// The ApplicationName of the ConfigurationTemplate resource.
	ApplicationName *string `field:"required" json:"applicationName" yaml:"applicationName"`
	// The TemplateName of the ConfigurationTemplate resource.
	TemplateName *string `field:"required" json:"templateName" yaml:"templateName"`
}

