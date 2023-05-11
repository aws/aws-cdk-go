package awspinpoint


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   templateProperty := &TemplateProperty{
//   	Name: jsii.String("name"),
//   	Version: jsii.String("version"),
//   }
//
type CfnCampaign_TemplateProperty struct {
	// `CfnCampaign.TemplateProperty.Name`.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// `CfnCampaign.TemplateProperty.Version`.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

