package awsglue


// Properties for defining a `CfnClassifier`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnClassifierProps := &CfnClassifierProps{
//   	CsvClassifier: &CsvClassifierProperty{
//   		AllowSingleColumn: jsii.Boolean(false),
//   		ContainsHeader: jsii.String("containsHeader"),
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
//   		GrokPattern: jsii.String("grokPattern"),
//
//   		// the properties below are optional
//   		CustomPatterns: jsii.String("customPatterns"),
//   		Name: jsii.String("name"),
//   	},
//   	JsonClassifier: &JsonClassifierProperty{
//   		JsonPath: jsii.String("jsonPath"),
//
//   		// the properties below are optional
//   		Name: jsii.String("name"),
//   	},
//   	XmlClassifier: &XMLClassifierProperty{
//   		Classification: jsii.String("classification"),
//   		RowTag: jsii.String("rowTag"),
//
//   		// the properties below are optional
//   		Name: jsii.String("name"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-classifier.html
//
type CfnClassifierProps struct {
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

