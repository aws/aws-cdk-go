package mixinsawsglue


// A classifier for `JSON` content.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   jsonClassifierProperty := &JsonClassifierProperty{
//   	JsonPath: jsii.String("jsonPath"),
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-classifier-jsonclassifier.html
//
type CfnClassifierPropsMixin_JsonClassifierProperty struct {
	// A `JsonPath` string defining the JSON data for the classifier to classify.
	//
	// AWS Glue supports a subset of `JsonPath` , as described in [Writing JsonPath Custom Classifiers](https://docs.aws.amazon.com/glue/latest/dg/custom-classifier.html#custom-classifier-json) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-classifier-jsonclassifier.html#cfn-glue-classifier-jsonclassifier-jsonpath
	//
	JsonPath *string `field:"optional" json:"jsonPath" yaml:"jsonPath"`
	// The name of the classifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-classifier-jsonclassifier.html#cfn-glue-classifier-jsonclassifier-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

