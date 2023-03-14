package awsiotanalytics


// Used to limit data to that which has arrived since the last execution of the action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deltaTimeProperty := &deltaTimeProperty{
//   	offsetSeconds: jsii.Number(123),
//   	timeExpression: jsii.String("timeExpression"),
//   }
//
type CfnDataset_DeltaTimeProperty struct {
	// The number of seconds of estimated in-flight lag time of message data.
	//
	// When you create dataset contents using message data from a specified timeframe, some message data might still be in flight when processing begins, and so do not arrive in time to be processed. Use this field to make allowances for the in flight time of your message data, so that data not processed from a previous timeframe is included with the next timeframe. Otherwise, missed message data would be excluded from processing during the next timeframe too, because its timestamp places it within the previous timeframe.
	OffsetSeconds *float64 `field:"required" json:"offsetSeconds" yaml:"offsetSeconds"`
	// An expression by which the time of the message data might be determined.
	//
	// This can be the name of a timestamp field or a SQL expression that is used to derive the time the message data was generated.
	TimeExpression *string `field:"required" json:"timeExpression" yaml:"timeExpression"`
}

