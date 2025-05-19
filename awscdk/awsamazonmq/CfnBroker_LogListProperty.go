package awsamazonmq


// The list of information about logs to be enabled for the specified broker.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   logListProperty := &LogListProperty{
//   	Audit: jsii.Boolean(false),
//   	General: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amazonmq-broker-loglist.html
//
type CfnBroker_LogListProperty struct {
	// Enables audit logging.
	//
	// Every user management action made using JMX or the ActiveMQ Web Console is logged. Does not apply to RabbitMQ brokers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amazonmq-broker-loglist.html#cfn-amazonmq-broker-loglist-audit
	//
	Audit interface{} `field:"optional" json:"audit" yaml:"audit"`
	// Enables general logging.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amazonmq-broker-loglist.html#cfn-amazonmq-broker-loglist-general
	//
	General interface{} `field:"optional" json:"general" yaml:"general"`
}

