package previewawsgluemixins


// Properties for CfnClassifierPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnClassifierMixinProps := &CfnClassifierMixinProps{
//   	CsvClassifier: &CsvClassifierProperty{
//   		AllowSingleColumn: jsii.Boolean(false),
//   		ContainsCustomDatatype: []*string{
//   			jsii.String("containsCustomDatatype"),
//   		},
//   		ContainsHeader: jsii.String("containsHeader"),
//   		CustomDatatypeConfigured: jsii.Boolean(false),
//   		Delimiter: jsii.String("delimiter"),
//   		DisableValueTrimming: jsii.Boolean(false),
//   		Header: []*string{
//   			jsii.String("header"),
//   		},
//   		Name: jsii.String("name"),
//   		QuoteSymbol: jsii.String("quoteSymbol"),
//   	},
//   	GrokClassifier: &GrokClassifierProperty{
//   		Classification: jsii.String("classification"),
//   		CustomPatterns: jsii.String("customPatterns"),
//   		GrokPattern: jsii.String("grokPattern"),
//   		Name: jsii.String("name"),
//   	},
//   	JsonClassifier: &JsonClassifierProperty{
//   		JsonPath: jsii.String("jsonPath"),
//   		Name: jsii.String("name"),
//   	},
//   	XmlClassifier: &XMLClassifierProperty{
//   		Classification: jsii.String("classification"),
//   		Name: jsii.String("name"),
//   		RowTag: jsii.String("rowTag"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-classifier.html
//
type CfnClassifierMixinProps struct {
	// A classifier for comma-separated values (CSV).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-classifier.html#cfn-glue-classifier-csvclassifier
	//
	CsvClassifier interface{} `field:"optional" json:"csvClassifier" yaml:"csvClassifier"`
	// A classifier that uses `grok` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-classifier.html#cfn-glue-classifier-grokclassifier
	//
	GrokClassifier interface{} `field:"optional" json:"grokClassifier" yaml:"grokClassifier"`
	// A classifier for JSON content.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-classifier.html#cfn-glue-classifier-jsonclassifier
	//
	JsonClassifier interface{} `field:"optional" json:"jsonClassifier" yaml:"jsonClassifier"`
	// A classifier for XML content.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-classifier.html#cfn-glue-classifier-xmlclassifier
	//
	XmlClassifier interface{} `field:"optional" json:"xmlClassifier" yaml:"xmlClassifier"`
}

