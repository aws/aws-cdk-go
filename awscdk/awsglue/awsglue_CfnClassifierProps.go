package awsglue


// Properties for defining a `CfnClassifier`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnClassifierProps := &cfnClassifierProps{
//   	csvClassifier: &csvClassifierProperty{
//   		allowSingleColumn: jsii.Boolean(false),
//   		containsHeader: jsii.String("containsHeader"),
//   		delimiter: jsii.String("delimiter"),
//   		disableValueTrimming: jsii.Boolean(false),
//   		header: []*string{
//   			jsii.String("header"),
//   		},
//   		name: jsii.String("name"),
//   		quoteSymbol: jsii.String("quoteSymbol"),
//   	},
//   	grokClassifier: &grokClassifierProperty{
//   		classification: jsii.String("classification"),
//   		grokPattern: jsii.String("grokPattern"),
//
//   		// the properties below are optional
//   		customPatterns: jsii.String("customPatterns"),
//   		name: jsii.String("name"),
//   	},
//   	jsonClassifier: &jsonClassifierProperty{
//   		jsonPath: jsii.String("jsonPath"),
//
//   		// the properties below are optional
//   		name: jsii.String("name"),
//   	},
//   	xmlClassifier: &xMLClassifierProperty{
//   		classification: jsii.String("classification"),
//   		rowTag: jsii.String("rowTag"),
//
//   		// the properties below are optional
//   		name: jsii.String("name"),
//   	},
//   }
//
type CfnClassifierProps struct {
	// A classifier for comma-separated values (CSV).
	CsvClassifier interface{} `field:"optional" json:"csvClassifier" yaml:"csvClassifier"`
	// A classifier that uses `grok` .
	GrokClassifier interface{} `field:"optional" json:"grokClassifier" yaml:"grokClassifier"`
	// A classifier for JSON content.
	JsonClassifier interface{} `field:"optional" json:"jsonClassifier" yaml:"jsonClassifier"`
	// A classifier for XML content.
	XmlClassifier interface{} `field:"optional" json:"xmlClassifier" yaml:"xmlClassifier"`
}

