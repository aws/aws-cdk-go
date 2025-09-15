package awskendra


// A reference to a Faq resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   faqReference := &FaqReference{
//   	FaqArn: jsii.String("faqArn"),
//   	FaqId: jsii.String("faqId"),
//   	IndexId: jsii.String("indexId"),
//   }
//
type FaqReference struct {
	// The ARN of the Faq resource.
	FaqArn *string `field:"required" json:"faqArn" yaml:"faqArn"`
	// The Id of the Faq resource.
	FaqId *string `field:"required" json:"faqId" yaml:"faqId"`
	// The IndexId of the Faq resource.
	IndexId *string `field:"required" json:"indexId" yaml:"indexId"`
}

