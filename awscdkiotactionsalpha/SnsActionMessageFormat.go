package awscdkiotactionsalpha


// SNS topic action message format options.
//
// Example:
//   import sns "github.com/aws/aws-cdk-go/awscdk"
//
//
//   topic := sns.NewTopic(this, jsii.String("MyTopic"))
//
//   topicRule := iot.NewTopicRule(this, jsii.String("TopicRule"), &TopicRuleProps{
//   	Sql: iot.IotSql_FromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id, year, month, day FROM 'device/+/data'")),
//   	Actions: []iAction{
//   		actions.NewSnsTopicAction(topic, &SnsTopicActionProps{
//   			MessageFormat: actions.SnsActionMessageFormat_JSON,
//   		}),
//   	},
//   })
//
// Experimental.
type SnsActionMessageFormat string

const (
	// RAW message format.
	// Experimental.
	SnsActionMessageFormat_RAW SnsActionMessageFormat = "RAW"
	// JSON message format.
	// Experimental.
	SnsActionMessageFormat_JSON SnsActionMessageFormat = "JSON"
)

