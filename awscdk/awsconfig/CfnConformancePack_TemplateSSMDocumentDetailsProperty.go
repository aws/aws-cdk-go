package awsconfig


// The TemplateSSMDocumentDetails object contains the name of the SSM document and the version of the SSM document.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   templateSSMDocumentDetailsProperty := &TemplateSSMDocumentDetailsProperty{
//   	DocumentName: jsii.String("documentName"),
//   	DocumentVersion: jsii.String("documentVersion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-conformancepack-templatessmdocumentdetails.html
//
type CfnConformancePack_TemplateSSMDocumentDetailsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-conformancepack-templatessmdocumentdetails.html#cfn-config-conformancepack-templatessmdocumentdetails-documentname
	//
	DocumentName *string `field:"optional" json:"documentName" yaml:"documentName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-conformancepack-templatessmdocumentdetails.html#cfn-config-conformancepack-templatessmdocumentdetails-documentversion
	//
	DocumentVersion *string `field:"optional" json:"documentVersion" yaml:"documentVersion"`
}

