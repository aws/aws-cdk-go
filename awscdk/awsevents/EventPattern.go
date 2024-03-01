package awsevents


// Events in Amazon CloudWatch Events are represented as JSON objects. For more information about JSON objects, see RFC 7159.
//
// **Important**: this class can only be used with a `Rule` class. In particular,
// do not use it with `CfnRule` class: your pattern will not be rendered
// correctly. In a `CfnRule` class, write the pattern as you normally would when
// directly writing CloudFormation.
//
// Rules use event patterns to select events and route them to targets. A
// pattern either matches an event or it doesn't. Event patterns are represented
// as JSON objects with a structure that is similar to that of events.
//
// It is important to remember the following about event pattern matching:
//
// - For a pattern to match an event, the event must contain all the field names
//   listed in the pattern. The field names must appear in the event with the
//   same nesting structure.
//
// - Other fields of the event not mentioned in the pattern are ignored;
//   effectively, there is a ``"*": "*"`` wildcard for fields not mentioned.
//
// - The matching is exact (character-by-character), without case-folding or any
//   other string normalization.
//
// - The values being matched follow JSON rules: Strings enclosed in quotes,
//   numbers, and the unquoted keywords true, false, and null.
//
// - Number matching is at the string representation level. For example, 300,
//   300.0, and 3.0e2 are not considered equal.
//
// For custom events, some optional properties are required. For more information, see
// [Minimum information needed for a valid custom event](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-events-structure.html#eb-custom-event).
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   fn := lambda.NewFunction(this, jsii.String("MyFunc"), &FunctionProps{
//   	Runtime: lambda.Runtime_NODEJS_LATEST(),
//   	Handler: jsii.String("index.handler"),
//   	Code: lambda.Code_FromInline(jsii.String("exports.handler = handler.toString()")),
//   })
//
//   rule := events.NewRule(this, jsii.String("rule"), &RuleProps{
//   	EventPattern: &EventPattern{
//   		Source: []*string{
//   			jsii.String("aws.ec2"),
//   		},
//   	},
//   })
//
//   queue := sqs.NewQueue(this, jsii.String("Queue"))
//
//   rule.AddTarget(targets.NewLambdaFunction(fn, &LambdaFunctionProps{
//   	DeadLetterQueue: queue,
//   	 // Optional: add a dead letter queue
//   	MaxEventAge: awscdk.Duration_Hours(jsii.Number(2)),
//   	 // Optional: set the maxEventAge retry policy
//   	RetryAttempts: jsii.Number(2),
//   }))
//
// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/CloudWatchEventsandEventPatterns.html
//
type EventPattern struct {
	// The 12-digit number identifying an AWS account.
	// Default: - No filtering on account.
	//
	Account *[]*string `field:"optional" json:"account" yaml:"account"`
	// A JSON object, whose content is at the discretion of the service originating the event.
	// Default: - No filtering on detail.
	//
	Detail *map[string]interface{} `field:"optional" json:"detail" yaml:"detail"`
	// Identifies, in combination with the source field, the fields and values that appear in the detail field.
	//
	// Represents the "detail-type" event field.
	// Default: - No filtering on detail type.
	//
	DetailType *[]*string `field:"optional" json:"detailType" yaml:"detailType"`
	// A unique value is generated for every event.
	//
	// This can be helpful in
	// tracing events as they move through rules to targets, and are processed.
	// Default: - No filtering on id.
	//
	Id *[]*string `field:"optional" json:"id" yaml:"id"`
	// Identifies the AWS region where the event originated.
	// Default: - No filtering on region.
	//
	Region *[]*string `field:"optional" json:"region" yaml:"region"`
	// This JSON array contains ARNs that identify resources that are involved in the event.
	//
	// Inclusion of these ARNs is at the discretion of the
	// service.
	//
	// For example, Amazon EC2 instance state-changes include Amazon EC2
	// instance ARNs, Auto Scaling events include ARNs for both instances and
	// Auto Scaling groups, but API calls with AWS CloudTrail do not include
	// resource ARNs.
	// Default: - No filtering on resource.
	//
	Resources *[]*string `field:"optional" json:"resources" yaml:"resources"`
	// Identifies the service that sourced the event.
	//
	// All events sourced from
	// within AWS begin with "aws." Customer-generated events can have any value
	// here, as long as it doesn't begin with "aws." We recommend the use of
	// Java package-name style reverse domain-name strings.
	//
	// To find the correct value for source for an AWS service, see the table in
	// AWS Service Namespaces. For example, the source value for Amazon
	// CloudFront is aws.cloudfront.
	// See: http://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#genref-aws-service-namespaces
	//
	// Default: - No filtering on source.
	//
	Source *[]*string `field:"optional" json:"source" yaml:"source"`
	// The event timestamp, which can be specified by the service originating the event.
	//
	// If the event spans a time interval, the service might choose
	// to report the start time, so this value can be noticeably before the time
	// the event is actually received.
	// Default: - No filtering on time.
	//
	Time *[]*string `field:"optional" json:"time" yaml:"time"`
	// By default, this is set to 0 (zero) in all events.
	// Default: - No filtering on version.
	//
	Version *[]*string `field:"optional" json:"version" yaml:"version"`
}

