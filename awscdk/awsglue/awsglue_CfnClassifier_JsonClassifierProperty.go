package awsglue


// A classifier for `JSON` content.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   jsonClassifierProperty := &jsonClassifierProperty{
//   	jsonPath: jsii.String("jsonPath"),
//
//   	// the properties below are optional
//   	name: jsii.String("name"),
//   }
//
type CfnClassifier_JsonClassifierProperty struct {
	// A `JsonPath` string defining the JSON data for the classifier to classify.
	//
	// AWS Glue supports a subset of `JsonPath` , as described in [Writing JsonPath Custom Classifiers](https://docs.aws.amazon.com/glue/latest/dg/custom-classifier.html#custom-classifier-json) .
	JsonPath *string `field:"required" json:"jsonPath" yaml:"jsonPath"`
	// The name of the classifier.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

