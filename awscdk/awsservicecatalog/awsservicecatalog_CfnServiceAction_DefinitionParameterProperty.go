package awsservicecatalog


// The list of parameters in JSON format.
//
// For example: `[{\"Name\":\"InstanceId\",\"Type\":\"TARGET\"}] or [{\"Name\":\"InstanceId\",\"Type\":\"TEXT_VALUE\"}]` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   definitionParameterProperty := &definitionParameterProperty{
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//   }
//
type CfnServiceAction_DefinitionParameterProperty struct {
	// The parameter key.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value of the parameter.
	Value *string `field:"required" json:"value" yaml:"value"`
}

