package awsquicksight


// The source template of the template.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   templateSourceTemplateProperty := &templateSourceTemplateProperty{
//   	arn: jsii.String("arn"),
//   }
//
type CfnTemplate_TemplateSourceTemplateProperty struct {
	// The Amazon Resource Name (ARN) of the resource.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
}

