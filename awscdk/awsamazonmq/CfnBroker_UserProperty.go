package awsamazonmq


// The list of broker users (persons or applications) who can access queues and topics.
//
// For Amazon MQ for RabbitMQ brokers, one and only one administrative user is accepted and created when a broker is first provisioned. All subsequent broker users are created via the RabbitMQ web console or by using the RabbitMQ management API.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   userProperty := &UserProperty{
//   	Password: jsii.String("password"),
//   	Username: jsii.String("username"),
//
//   	// the properties below are optional
//   	ConsoleAccess: jsii.Boolean(false),
//   	Groups: []*string{
//   		jsii.String("groups"),
//   	},
//   	JolokiaApiAccess: jsii.String("jolokiaApiAccess"),
//   	ReplicationUser: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amazonmq-broker-user.html
//
type CfnBroker_UserProperty struct {
	// The password of the user.
	//
	// This value must be at least 12 characters long, must contain at least 4 unique characters, and must not contain commas, colons, or equal signs (,:=).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amazonmq-broker-user.html#cfn-amazonmq-broker-user-password
	//
	Password *string `field:"required" json:"password" yaml:"password"`
	// The username of the broker user.
	//
	// For Amazon MQ for ActiveMQ brokers, this value can contain only alphanumeric characters, dashes, periods, underscores, and tildes (- . _ ~). For Amazon MQ for RabbitMQ brokers, this value can contain only alphanumeric characters, dashes, periods, underscores (- . _). This value must not contain a tilde (~) character. Amazon MQ prohibts using guest as a valid usename. This value must be 2-100 characters long.
	//
	// > Do not add personally identifiable information (PII) or other confidential or sensitive information in broker usernames. Broker usernames are accessible to other AWS services, including CloudWatch Logs . Broker usernames are not intended to be used for private or sensitive data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amazonmq-broker-user.html#cfn-amazonmq-broker-user-username
	//
	Username *string `field:"required" json:"username" yaml:"username"`
	// Enables access to the ActiveMQ web console for the ActiveMQ user.
	//
	// Does not apply to RabbitMQ brokers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amazonmq-broker-user.html#cfn-amazonmq-broker-user-consoleaccess
	//
	ConsoleAccess interface{} `field:"optional" json:"consoleAccess" yaml:"consoleAccess"`
	// The list of groups (20 maximum) to which the ActiveMQ user belongs.
	//
	// This value can contain only alphanumeric characters, dashes, periods, underscores, and tildes (- . _ ~). This value must be 2-100 characters long. Does not apply to RabbitMQ brokers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amazonmq-broker-user.html#cfn-amazonmq-broker-user-groups
	//
	Groups *[]*string `field:"optional" json:"groups" yaml:"groups"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amazonmq-broker-user.html#cfn-amazonmq-broker-user-jolokiaapiaccess
	//
	JolokiaApiAccess *string `field:"optional" json:"jolokiaApiAccess" yaml:"jolokiaApiAccess"`
	// Defines if this user is intended for CRDR replication purposes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amazonmq-broker-user.html#cfn-amazonmq-broker-user-replicationuser
	//
	ReplicationUser interface{} `field:"optional" json:"replicationUser" yaml:"replicationUser"`
}

