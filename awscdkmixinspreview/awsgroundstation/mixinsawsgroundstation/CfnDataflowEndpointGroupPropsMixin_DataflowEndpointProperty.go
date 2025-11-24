package mixinsawsgroundstation


// Contains information such as socket address and name that defines an endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataflowEndpointProperty := &DataflowEndpointProperty{
//   	Address: &SocketAddressProperty{
//   		Name: jsii.String("name"),
//   		Port: jsii.Number(123),
//   	},
//   	Mtu: jsii.Number(123),
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroup-dataflowendpoint.html
//
type CfnDataflowEndpointGroupPropsMixin_DataflowEndpointProperty struct {
	// The address and port of an endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroup-dataflowendpoint.html#cfn-groundstation-dataflowendpointgroup-dataflowendpoint-address
	//
	Address interface{} `field:"optional" json:"address" yaml:"address"`
	// Maximum transmission unit (MTU) size in bytes of a dataflow endpoint.
	//
	// Valid values are between 1400 and 1500. A default value of 1500 is used if not set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroup-dataflowendpoint.html#cfn-groundstation-dataflowendpointgroup-dataflowendpoint-mtu
	//
	Mtu *float64 `field:"optional" json:"mtu" yaml:"mtu"`
	// The endpoint name.
	//
	// When listing available contacts for a satellite, Ground Station searches for a dataflow endpoint whose name matches the value specified by the dataflow endpoint config of the selected mission profile. If no matching dataflow endpoints are found then Ground Station will not display any available contacts for the satellite.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-dataflowendpointgroup-dataflowendpoint.html#cfn-groundstation-dataflowendpointgroup-dataflowendpoint-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

