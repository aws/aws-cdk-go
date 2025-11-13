package interfacesawsamazonmq


// A reference to a Broker resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   brokerReference := &BrokerReference{
//   	BrokerArn: jsii.String("brokerArn"),
//   	BrokerId: jsii.String("brokerId"),
//   }
//
type BrokerReference struct {
	// The ARN of the Broker resource.
	BrokerArn *string `field:"required" json:"brokerArn" yaml:"brokerArn"`
	// The Id of the Broker resource.
	BrokerId *string `field:"required" json:"brokerId" yaml:"brokerId"`
}

