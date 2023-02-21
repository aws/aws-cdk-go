package awsdatapipeline


// A key-value pair that describes a property of a `PipelineObject` .
//
// The value is specified as either a string value ( `StringValue` ) or a reference to another object ( `RefValue` ) but not as both. To view fields for a data pipeline object, see [Pipeline Object Reference](https://docs.aws.amazon.com/datapipeline/latest/DeveloperGuide/dp-pipeline-objects.html) in the *AWS Data Pipeline Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fieldProperty := &fieldProperty{
//   	key: jsii.String("key"),
//
//   	// the properties below are optional
//   	refValue: jsii.String("refValue"),
//   	stringValue: jsii.String("stringValue"),
//   }
//
type CfnPipeline_FieldProperty struct {
	// Specifies the name of a field for a particular object.
	//
	// To view valid values for a particular field, see [Pipeline Object Reference](https://docs.aws.amazon.com/datapipeline/latest/DeveloperGuide/dp-pipeline-objects.html) in the *AWS Data Pipeline Developer Guide* .
	Key *string `field:"required" json:"key" yaml:"key"`
	// A field value that you specify as an identifier of another object in the same pipeline definition.
	//
	// > You can specify the field value as either a string value ( `StringValue` ) or a reference to another object ( `RefValue` ), but not both.
	//
	// Required if the key that you are using requires it.
	RefValue *string `field:"optional" json:"refValue" yaml:"refValue"`
	// A field value that you specify as a string.
	//
	// To view valid values for a particular field, see [Pipeline Object Reference](https://docs.aws.amazon.com/datapipeline/latest/DeveloperGuide/dp-pipeline-objects.html) in the *AWS Data Pipeline Developer Guide* .
	//
	// > You can specify the field value as either a string value ( `StringValue` ) or a reference to another object ( `RefValue` ), but not both.
	//
	// Required if the key that you are using requires it.
	StringValue *string `field:"optional" json:"stringValue" yaml:"stringValue"`
}

