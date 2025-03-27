package awsquicksight


// The source template of the template.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   templateSourceTemplateProperty := &TemplateSourceTemplateProperty{
//   	Arn: jsii.String("arn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-templatesourcetemplate.html
//
type CfnTemplate_TemplateSourceTemplateProperty struct {
	// The Amazon Resource Name (ARN) of the resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-templatesourcetemplate.html#cfn-quicksight-template-templatesourcetemplate-arn
	//
	Arn *string `field:"required" json:"arn" yaml:"arn"`
}

