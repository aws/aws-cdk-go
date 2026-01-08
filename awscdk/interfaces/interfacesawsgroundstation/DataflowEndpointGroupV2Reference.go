package interfacesawsgroundstation


// A reference to a DataflowEndpointGroupV2 resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataflowEndpointGroupV2Reference := &DataflowEndpointGroupV2Reference{
//   	DataflowEndpointGroupV2Arn: jsii.String("dataflowEndpointGroupV2Arn"),
//   	DataflowEndpointGroupV2Id: jsii.String("dataflowEndpointGroupV2Id"),
//   }
//
type DataflowEndpointGroupV2Reference struct {
	// The ARN of the DataflowEndpointGroupV2 resource.
	DataflowEndpointGroupV2Arn *string `field:"required" json:"dataflowEndpointGroupV2Arn" yaml:"dataflowEndpointGroupV2Arn"`
	// The Id of the DataflowEndpointGroupV2 resource.
	DataflowEndpointGroupV2Id *string `field:"required" json:"dataflowEndpointGroupV2Id" yaml:"dataflowEndpointGroupV2Id"`
}

