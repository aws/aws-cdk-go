package awsgroundstation


// Contains information such as socket address and name that defines an endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataflowEndpointProperty := &dataflowEndpointProperty{
//   	address: &socketAddressProperty{
//   		name: jsii.String("name"),
//   		port: jsii.Number(123),
//   	},
//   	mtu: jsii.Number(123),
//   	name: jsii.String("name"),
//   }
//
type CfnDataflowEndpointGroup_DataflowEndpointProperty struct {
	// The address and port of an endpoint.
	Address interface{} `field:"optional" json:"address" yaml:"address"`
	// Maximum transmission unit (MTU) size in bytes of a dataflow endpoint.
	//
	// Valid values are between 1400 and 1500. A default value of 1500 is used if not set.
	Mtu *float64 `field:"optional" json:"mtu" yaml:"mtu"`
	// The endpoint name.
	//
	// When listing available contacts for a satellite, Ground Station searches for a dataflow endpoint whose name matches the value specified by the dataflow endpoint config of the selected mission profile. If no matching dataflow endpoints are found then Ground Station will not display any available contacts for the satellite.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

