package awsservicecatalog


// A reference to a ServiceAction resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serviceActionReference := &ServiceActionReference{
//   	ServiceActionId: jsii.String("serviceActionId"),
//   }
//
type ServiceActionReference struct {
	// The Id of the ServiceAction resource.
	ServiceActionId *string `field:"required" json:"serviceActionId" yaml:"serviceActionId"`
}

