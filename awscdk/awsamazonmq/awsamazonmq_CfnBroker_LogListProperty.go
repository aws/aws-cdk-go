package awsamazonmq


// The list of information about logs to be enabled for the specified broker.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   logListProperty := &logListProperty{
//   	audit: jsii.Boolean(false),
//   	general: jsii.Boolean(false),
//   }
//
type CfnBroker_LogListProperty struct {
	// Enables audit logging.
	//
	// Every user management action made using JMX or the ActiveMQ Web Console is logged. Does not apply to RabbitMQ brokers.
	Audit interface{} `field:"optional" json:"audit" yaml:"audit"`
	// Enables general logging.
	General interface{} `field:"optional" json:"general" yaml:"general"`
}

