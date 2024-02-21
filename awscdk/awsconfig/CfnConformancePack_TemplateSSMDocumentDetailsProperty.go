package awsconfig


// This API allows you to create a conformance pack template with an AWS Systems Manager document (SSM document).
//
// To deploy a conformance pack using an SSM document, first create an SSM document with conformance pack content, and then provide the `DocumentName` in the [PutConformancePack API](https://docs.aws.amazon.com/config/latest/APIReference/API_PutConformancePack.html) . You can also provide the `DocumentVersion` .
//
// The `TemplateSSMDocumentDetails` object contains the name of the SSM document and the version of the SSM document.
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
	// The name or Amazon Resource Name (ARN) of the SSM document to use to create a conformance pack.
	//
	// If you use the document name, AWS Config checks only your account and AWS Region for the SSM document.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-conformancepack-templatessmdocumentdetails.html#cfn-config-conformancepack-templatessmdocumentdetails-documentname
	//
	DocumentName *string `field:"optional" json:"documentName" yaml:"documentName"`
	// The version of the SSM document to use to create a conformance pack.
	//
	// By default, AWS Config uses the latest version.
	//
	// > This field is optional.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-conformancepack-templatessmdocumentdetails.html#cfn-config-conformancepack-templatessmdocumentdetails-documentversion
	//
	DocumentVersion *string `field:"optional" json:"documentVersion" yaml:"documentVersion"`
}

