package awssagemaker


// The Json format.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   jsonProperty := &JsonProperty{
//   	Line: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-dataqualityjobdefinition-json.html
//
type CfnDataQualityJobDefinition_JsonProperty struct {
	// A boolean flag indicating if it is JSON line format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-dataqualityjobdefinition-json.html#cfn-sagemaker-dataqualityjobdefinition-json-line
	//
	Line interface{} `field:"optional" json:"line" yaml:"line"`
}

