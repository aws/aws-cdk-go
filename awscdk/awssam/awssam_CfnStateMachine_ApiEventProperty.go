package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   apiEventProperty := &apiEventProperty{
//   	method: jsii.String("method"),
//   	path: jsii.String("path"),
//
//   	// the properties below are optional
//   	restApiId: jsii.String("restApiId"),
//   }
//
type CfnStateMachine_ApiEventProperty struct {
	// `CfnStateMachine.ApiEventProperty.Method`.
	Method *string `field:"required" json:"method" yaml:"method"`
	// `CfnStateMachine.ApiEventProperty.Path`.
	Path *string `field:"required" json:"path" yaml:"path"`
	// `CfnStateMachine.ApiEventProperty.RestApiId`.
	RestApiId *string `field:"optional" json:"restApiId" yaml:"restApiId"`
}

