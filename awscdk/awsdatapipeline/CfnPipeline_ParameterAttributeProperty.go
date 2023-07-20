package awsdatapipeline


// `Attribute` is a property of `ParameterObject` that defines the attributes of a parameter object as key-value pairs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   parameterAttributeProperty := &ParameterAttributeProperty{
//   	Key: jsii.String("key"),
//   	StringValue: jsii.String("stringValue"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-parameterattribute.html
//
type CfnPipeline_ParameterAttributeProperty struct {
	// The field identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-parameterattribute.html#cfn-datapipeline-pipeline-parameterattribute-key
	//
	Key *string `field:"required" json:"key" yaml:"key"`
	// The field value, expressed as a String.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-parameterattribute.html#cfn-datapipeline-pipeline-parameterattribute-stringvalue
	//
	StringValue *string `field:"required" json:"stringValue" yaml:"stringValue"`
}

