package interfacesawslambda


// A reference to a EventInvokeConfig resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventInvokeConfigReference := &EventInvokeConfigReference{
//   	FunctionName: jsii.String("functionName"),
//   	Qualifier: jsii.String("qualifier"),
//   }
//
type EventInvokeConfigReference struct {
	// The FunctionName of the EventInvokeConfig resource.
	FunctionName *string `field:"required" json:"functionName" yaml:"functionName"`
	// The Qualifier of the EventInvokeConfig resource.
	Qualifier *string `field:"required" json:"qualifier" yaml:"qualifier"`
}

