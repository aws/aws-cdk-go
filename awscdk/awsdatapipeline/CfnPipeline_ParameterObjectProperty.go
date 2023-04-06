package awsdatapipeline


// Contains information about a parameter object.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   parameterObjectProperty := &ParameterObjectProperty{
//   	Attributes: []interface{}{
//   		&ParameterAttributeProperty{
//   			Key: jsii.String("key"),
//   			StringValue: jsii.String("stringValue"),
//   		},
//   	},
//   	Id: jsii.String("id"),
//   }
//
type CfnPipeline_ParameterObjectProperty struct {
	// The attributes of the parameter object.
	Attributes interface{} `field:"required" json:"attributes" yaml:"attributes"`
	// The ID of the parameter object.
	Id *string `field:"required" json:"id" yaml:"id"`
}

