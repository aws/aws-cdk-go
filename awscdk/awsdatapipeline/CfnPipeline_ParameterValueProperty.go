package awsdatapipeline


// A value or list of parameter values.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   parameterValueProperty := &ParameterValueProperty{
//   	Id: jsii.String("id"),
//   	StringValue: jsii.String("stringValue"),
//   }
//
type CfnPipeline_ParameterValueProperty struct {
	// The ID of the parameter value.
	Id *string `field:"required" json:"id" yaml:"id"`
	// The field value, expressed as a String.
	StringValue *string `field:"required" json:"stringValue" yaml:"stringValue"`
}

