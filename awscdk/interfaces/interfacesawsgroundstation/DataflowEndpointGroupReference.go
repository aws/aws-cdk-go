package interfacesawsgroundstation


// A reference to a DataflowEndpointGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataflowEndpointGroupReference := &DataflowEndpointGroupReference{
//   	DataflowEndpointGroupArn: jsii.String("dataflowEndpointGroupArn"),
//   	DataflowEndpointGroupId: jsii.String("dataflowEndpointGroupId"),
//   }
//
type DataflowEndpointGroupReference struct {
	// The ARN of the DataflowEndpointGroup resource.
	DataflowEndpointGroupArn *string `field:"required" json:"dataflowEndpointGroupArn" yaml:"dataflowEndpointGroupArn"`
	// The Id of the DataflowEndpointGroup resource.
	DataflowEndpointGroupId *string `field:"required" json:"dataflowEndpointGroupId" yaml:"dataflowEndpointGroupId"`
}

