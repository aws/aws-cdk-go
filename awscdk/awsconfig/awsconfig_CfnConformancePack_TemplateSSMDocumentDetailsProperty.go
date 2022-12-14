package awsconfig


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   templateSSMDocumentDetailsProperty := &templateSSMDocumentDetailsProperty{
//   	documentName: jsii.String("documentName"),
//   	documentVersion: jsii.String("documentVersion"),
//   }
//
type CfnConformancePack_TemplateSSMDocumentDetailsProperty struct {
	// `CfnConformancePack.TemplateSSMDocumentDetailsProperty.DocumentName`.
	DocumentName *string `field:"optional" json:"documentName" yaml:"documentName"`
	// `CfnConformancePack.TemplateSSMDocumentDetailsProperty.DocumentVersion`.
	DocumentVersion *string `field:"optional" json:"documentVersion" yaml:"documentVersion"`
}

