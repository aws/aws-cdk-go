package awsglue


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   viewRepresentationProperty := &ViewRepresentationProperty{
//   	Dialect: jsii.String("dialect"),
//   	DialectVersion: jsii.String("dialectVersion"),
//   	ValidationConnection: jsii.String("validationConnection"),
//   	ViewExpandedText: jsii.String("viewExpandedText"),
//   	ViewOriginalText: jsii.String("viewOriginalText"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-viewrepresentation.html
//
type CfnTable_ViewRepresentationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-viewrepresentation.html#cfn-glue-table-viewrepresentation-dialect
	//
	Dialect *string `field:"optional" json:"dialect" yaml:"dialect"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-viewrepresentation.html#cfn-glue-table-viewrepresentation-dialectversion
	//
	DialectVersion *string `field:"optional" json:"dialectVersion" yaml:"dialectVersion"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-viewrepresentation.html#cfn-glue-table-viewrepresentation-validationconnection
	//
	ValidationConnection *string `field:"optional" json:"validationConnection" yaml:"validationConnection"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-viewrepresentation.html#cfn-glue-table-viewrepresentation-viewexpandedtext
	//
	ViewExpandedText *string `field:"optional" json:"viewExpandedText" yaml:"viewExpandedText"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-viewrepresentation.html#cfn-glue-table-viewrepresentation-vieworiginaltext
	//
	ViewOriginalText *string `field:"optional" json:"viewOriginalText" yaml:"viewOriginalText"`
}

