package awsamazonmq


// A list of information about the configuration.
//
// > Does not apply to RabbitMQ brokers.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   configurationIdProperty := &ConfigurationIdProperty{
//   	Id: jsii.String("id"),
//   	Revision: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amazonmq-broker-configurationid.html
//
type CfnBroker_ConfigurationIdProperty struct {
	// The unique ID that Amazon MQ generates for the configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amazonmq-broker-configurationid.html#cfn-amazonmq-broker-configurationid-id
	//
	Id *string `field:"required" json:"id" yaml:"id"`
	// The revision number of the configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amazonmq-broker-configurationid.html#cfn-amazonmq-broker-configurationid-revision
	//
	Revision *float64 `field:"required" json:"revision" yaml:"revision"`
}

