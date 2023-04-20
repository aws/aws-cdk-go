package awsstepfunctionstasks


// Properties for InvokeFunction.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var payload interface{}
//
//   invokeFunctionProps := &InvokeFunctionProps{
//   	Payload: map[string]interface{}{
//   		"payloadKey": payload,
//   	},
//   }
//
// Deprecated: use `LambdaInvoke`.
type InvokeFunctionProps struct {
	// The JSON that you want to provide to your Lambda function as input.
	//
	// This parameter is named as payload to keep consistent with RunLambdaTask class.
	// Deprecated: use `LambdaInvoke`.
	Payload *map[string]interface{} `field:"optional" json:"payload" yaml:"payload"`
}

