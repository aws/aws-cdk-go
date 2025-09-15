package awspcaconnectorad


// A reference to a Template resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   templateReference := &TemplateReference{
//   	TemplateArn: jsii.String("templateArn"),
//   }
//
type TemplateReference struct {
	// The TemplateArn of the Template resource.
	TemplateArn *string `field:"required" json:"templateArn" yaml:"templateArn"`
}

