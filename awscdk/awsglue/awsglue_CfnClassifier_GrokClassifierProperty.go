package awsglue


// A classifier that uses `grok` patterns.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   grokClassifierProperty := &grokClassifierProperty{
//   	classification: jsii.String("classification"),
//   	grokPattern: jsii.String("grokPattern"),
//
//   	// the properties below are optional
//   	customPatterns: jsii.String("customPatterns"),
//   	name: jsii.String("name"),
//   }
//
type CfnClassifier_GrokClassifierProperty struct {
	// An identifier of the data format that the classifier matches, such as Twitter, JSON, Omniture logs, and so on.
	Classification *string `field:"required" json:"classification" yaml:"classification"`
	// The grok pattern applied to a data store by this classifier.
	//
	// For more information, see built-in patterns in [Writing Custom Classifiers](https://docs.aws.amazon.com/glue/latest/dg/custom-classifier.html) .
	GrokPattern *string `field:"required" json:"grokPattern" yaml:"grokPattern"`
	// Optional custom grok patterns defined by this classifier.
	//
	// For more information, see custom patterns in [Writing Custom Classifiers](https://docs.aws.amazon.com/glue/latest/dg/custom-classifier.html) .
	CustomPatterns *string `field:"optional" json:"customPatterns" yaml:"customPatterns"`
	// The name of the classifier.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

