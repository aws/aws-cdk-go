package awsbedrockalpha


// Properties for creating a text template configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_alpha "github.com/aws/aws-cdk-go/awsbedrockalpha"
//
//   textTemplateConfigurationProps := &TextTemplateConfigurationProps{
//   	Text: jsii.String("text"),
//
//   	// the properties below are optional
//   	InputVariables: []*string{
//   		jsii.String("inputVariables"),
//   	},
//   }
//
// Experimental.
type TextTemplateConfigurationProps struct {
	// The text content of the template.
	// Experimental.
	Text *string `field:"required" json:"text" yaml:"text"`
	// The input variables for the template.
	// Default: - No input variables.
	//
	// Experimental.
	InputVariables *[]*string `field:"optional" json:"inputVariables" yaml:"inputVariables"`
}

