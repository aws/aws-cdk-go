package awscdk


// Provenance and confidence for context.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnContextTrust := &CfnContextTrust{
//   	Conf: cdk.CfnContextTrustConfidence_HIGH,
//   	Src: cdk.CfnContextTrustSource_AUTHORED,
//
//   	// the properties below are optional
//   	Cite: jsii.String("cite"),
//   	Note: jsii.String("note"),
//   }
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-attribute-metadata.html#aws-attribute-metadata-context-schema
//
type CfnContextTrust struct {
	// Confidence in the context.
	Conf CfnContextTrustConfidence `field:"required" json:"conf" yaml:"conf"`
	// How the context was produced.
	Src CfnContextTrustSource `field:"required" json:"src" yaml:"src"`
	// Reference to the source of the context.
	// Default: - no citation.
	//
	Cite *string `field:"optional" json:"cite" yaml:"cite"`
	// Additional provenance details.
	// Default: - no note.
	//
	Note *string `field:"optional" json:"note" yaml:"note"`
}

