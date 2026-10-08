package awsglue


// A structure that defines a view representation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
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
type CfnTablePropsMixin_ViewRepresentationProperty struct {
	// The dialect of the view representation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-viewrepresentation.html#cfn-glue-table-viewrepresentation-dialect
	//
	Dialect *string `field:"optional" json:"dialect" yaml:"dialect"`
	// The version of the dialect.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-viewrepresentation.html#cfn-glue-table-viewrepresentation-dialectversion
	//
	DialectVersion *string `field:"optional" json:"dialectVersion" yaml:"dialectVersion"`
	// The connection used for validation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-viewrepresentation.html#cfn-glue-table-viewrepresentation-validationconnection
	//
	ValidationConnection *string `field:"optional" json:"validationConnection" yaml:"validationConnection"`
	// The expanded text of the view.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-viewrepresentation.html#cfn-glue-table-viewrepresentation-viewexpandedtext
	//
	ViewExpandedText *string `field:"optional" json:"viewExpandedText" yaml:"viewExpandedText"`
	// The original text of the view.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-viewrepresentation.html#cfn-glue-table-viewrepresentation-vieworiginaltext
	//
	ViewOriginalText *string `field:"optional" json:"viewOriginalText" yaml:"viewOriginalText"`
}

