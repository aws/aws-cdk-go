package awselasticache


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   readEndPointProperty := &ReadEndPointProperty{
//   	Addresses: jsii.String("addresses"),
//   	AddressesList: []*string{
//   		jsii.String("addressesList"),
//   	},
//   	Ports: jsii.String("ports"),
//   	PortsList: []*string{
//   		jsii.String("portsList"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-replicationgroup-readendpoint.html
//
type CfnReplicationGroup_ReadEndPointProperty struct {
	// A string containing a comma-separated list of endpoints for the primary and read-only replicas, formatted as [address1, address2, ...]. The order of the addresses maps to the order of the ports from the ReadEndPoint.Ports attribute.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-replicationgroup-readendpoint.html#cfn-elasticache-replicationgroup-readendpoint-addresses
	//
	Addresses *string `field:"optional" json:"addresses" yaml:"addresses"`
	// A list of endpoints for the read-only replicas.
	//
	// The order of the addresses maps to the order of the ports from the ReadEndPoint.Ports attribute.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-replicationgroup-readendpoint.html#cfn-elasticache-replicationgroup-readendpoint-addresseslist
	//
	AddressesList *[]*string `field:"optional" json:"addressesList" yaml:"addressesList"`
	// A string containing a comma-separated list of ports for the read-only replicas, formatted as [port1, port2, ...]. The order of the ports maps to the order of the addresses from the ReadEndPoint.Addresses attribute.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-replicationgroup-readendpoint.html#cfn-elasticache-replicationgroup-readendpoint-ports
	//
	Ports *string `field:"optional" json:"ports" yaml:"ports"`
	// A list of ports for the read-only replicas.
	//
	// The order of the ports maps to the order of the addresses from the ReadEndPoint.Addresses attribute.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-replicationgroup-readendpoint.html#cfn-elasticache-replicationgroup-readendpoint-portslist
	//
	PortsList *[]*string `field:"optional" json:"portsList" yaml:"portsList"`
}

