package awsquicksight


// A reference to a Template resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   templateReference := &TemplateReference{
//   	AwsAccountId: jsii.String("awsAccountId"),
//   	TemplateArn: jsii.String("templateArn"),
//   	TemplateId: jsii.String("templateId"),
//   }
//
type TemplateReference struct {
	// The AwsAccountId of the Template resource.
	AwsAccountId *string `field:"required" json:"awsAccountId" yaml:"awsAccountId"`
	// The ARN of the Template resource.
	TemplateArn *string `field:"required" json:"templateArn" yaml:"templateArn"`
	// The TemplateId of the Template resource.
	TemplateId *string `field:"required" json:"templateId" yaml:"templateId"`
}

