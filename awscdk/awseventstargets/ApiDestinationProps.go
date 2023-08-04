package awseventstargets

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
)

// Customize the EventBridge Api Destinations Target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var queue queue
//   var role role
//   var ruleTargetInput ruleTargetInput
//
//   apiDestinationProps := &ApiDestinationProps{
//   	DeadLetterQueue: queue,
//   	Event: ruleTargetInput,
//   	EventRole: role,
//   	HeaderParameters: map[string]*string{
//   		"headerParametersKey": jsii.String("headerParameters"),
//   	},
//   	MaxEventAge: cdk.Duration_Minutes(jsii.Number(30)),
//   	PathParameterValues: []*string{
//   		jsii.String("pathParameterValues"),
//   	},
//   	QueryStringParameters: map[string]*string{
//   		"queryStringParametersKey": jsii.String("queryStringParameters"),
//   	},
//   	RetryAttempts: jsii.Number(123),
//   }
//
type ApiDestinationProps struct {
	// The SQS queue to be used as deadLetterQueue. Check out the [considerations for using a dead-letter queue](https://docs.aws.amazon.com/eventbridge/latest/userguide/rule-dlq.html#dlq-considerations).
	//
	// The events not successfully delivered are automatically retried for a specified period of time,
	// depending on the retry policy of the target.
	// If an event is not delivered before all retry attempts are exhausted, it will be sent to the dead letter queue.
	// Default: - no dead-letter queue.
	//
	DeadLetterQueue awssqs.IQueue `field:"optional" json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// The maximum age of a request that Lambda sends to a function for processing.
	//
	// Minimum value of 60.
	// Maximum value of 86400.
	// Default: Duration.hours(24)
	//
	MaxEventAge awscdk.Duration `field:"optional" json:"maxEventAge" yaml:"maxEventAge"`
	// The maximum number of times to retry when the function returns an error.
	//
	// Minimum value of 0.
	// Maximum value of 185.
	// Default: 185.
	//
	RetryAttempts *float64 `field:"optional" json:"retryAttempts" yaml:"retryAttempts"`
	// The event to send.
	// Default: - the entire EventBridge event.
	//
	Event awsevents.RuleTargetInput `field:"optional" json:"event" yaml:"event"`
	// The role to assume before invoking the target.
	// Default: - a new role will be created.
	//
	EventRole awsiam.IRole `field:"optional" json:"eventRole" yaml:"eventRole"`
	// Additional headers sent to the API Destination.
	//
	// These are merged with headers specified on the Connection, with
	// the headers on the Connection taking precedence.
	//
	// You can only specify secret values on the Connection.
	// Default: - none.
	//
	HeaderParameters *map[string]*string `field:"optional" json:"headerParameters" yaml:"headerParameters"`
	// Path parameters to insert in place of path wildcards (`*`).
	//
	// If the API destination has a wilcard in the path, these path parts
	// will be inserted in that place.
	// Default: - none.
	//
	PathParameterValues *[]*string `field:"optional" json:"pathParameterValues" yaml:"pathParameterValues"`
	// Additional query string parameters sent to the API Destination.
	//
	// These are merged with headers specified on the Connection, with
	// the headers on the Connection taking precedence.
	//
	// You can only specify secret values on the Connection.
	// Default: - none.
	//
	QueryStringParameters *map[string]*string `field:"optional" json:"queryStringParameters" yaml:"queryStringParameters"`
}

