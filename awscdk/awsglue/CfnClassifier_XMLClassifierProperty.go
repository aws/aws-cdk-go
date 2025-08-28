package awsglue


// A classifier for `XML` content.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   xMLClassifierProperty := &XMLClassifierProperty{
//   	Classification: jsii.String("classification"),
//   	RowTag: jsii.String("rowTag"),
//
//   	// the properties below are optional
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-classifier-xmlclassifier.html
//
type CfnClassifier_XMLClassifierProperty struct {
	// An identifier of the data format that the classifier matches.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-classifier-xmlclassifier.html#cfn-glue-classifier-xmlclassifier-classification
	//
	Classification *string `field:"required" json:"classification" yaml:"classification"`
	// The XML tag designating the element that contains each record in an XML document being parsed.
	//
	// This can't identify a self-closing element (closed by `/>` ). An empty row element that contains only attributes can be parsed as long as it ends with a closing tag (for example, `<row item_a="A" item_b="B"></row>` is okay, but `<row item_a="A" item_b="B" />` is not).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-classifier-xmlclassifier.html#cfn-glue-classifier-xmlclassifier-rowtag
	//
	RowTag *string `field:"required" json:"rowTag" yaml:"rowTag"`
	// The name of the classifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-classifier-xmlclassifier.html#cfn-glue-classifier-xmlclassifier-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

