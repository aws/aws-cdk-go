package interfacesawsresiliencehubv2


// A reference to a ServiceFunction resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serviceFunctionReference := &ServiceFunctionReference{
//   	ServiceArn: jsii.String("serviceArn"),
//   	ServiceFunctionId: jsii.String("serviceFunctionId"),
//   }
//
type ServiceFunctionReference struct {
	// The ServiceArn of the ServiceFunction resource.
	ServiceArn *string `field:"required" json:"serviceArn" yaml:"serviceArn"`
	// The ServiceFunctionId of the ServiceFunction resource.
	ServiceFunctionId *string `field:"required" json:"serviceFunctionId" yaml:"serviceFunctionId"`
}

