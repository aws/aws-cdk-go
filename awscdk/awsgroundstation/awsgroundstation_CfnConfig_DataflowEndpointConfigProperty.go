package awsgroundstation


// Provides information to AWS Ground Station about which IP endpoints to use during a contact.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataflowEndpointConfigProperty := &dataflowEndpointConfigProperty{
//   	dataflowEndpointName: jsii.String("dataflowEndpointName"),
//   	dataflowEndpointRegion: jsii.String("dataflowEndpointRegion"),
//   }
//
type CfnConfig_DataflowEndpointConfigProperty struct {
	// The name of the dataflow endpoint to use during contacts.
	DataflowEndpointName *string `field:"optional" json:"dataflowEndpointName" yaml:"dataflowEndpointName"`
	// The region of the dataflow endpoint to use during contacts.
	//
	// When omitted, Ground Station will use the region of the contact.
	DataflowEndpointRegion *string `field:"optional" json:"dataflowEndpointRegion" yaml:"dataflowEndpointRegion"`
}

