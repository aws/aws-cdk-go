package awspinpoint


// Specifies the name and version of the message template to use for the message.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   templateProperty := &templateProperty{
//   	name: jsii.String("name"),
//   	version: jsii.String("version"),
//   }
//
type CfnCampaign_TemplateProperty struct {
	// The name of the message template to use for the message.
	//
	// If specified, this value must match the name of an existing message template.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The unique identifier for the version of the message template to use for the message.
	//
	// If specified, this value must match the identifier for an existing template version. To retrieve a list of versions and version identifiers for a template, use the Template Versions resource.
	//
	// If you don't specify a value for this property, Amazon Pinpoint uses the *active version* of the template. The *active version* is typically the version of a template that's been most recently reviewed and approved for use, depending on your workflow. It isn't necessarily the latest version of a template.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

