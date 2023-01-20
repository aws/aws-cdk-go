package awslogs


// Properties for defining a `CfnSubscriptionFilter`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSubscriptionFilterProps := &cfnSubscriptionFilterProps{
//   	destinationArn: jsii.String("destinationArn"),
//   	filterPattern: jsii.String("filterPattern"),
//   	logGroupName: jsii.String("logGroupName"),
//
//   	// the properties below are optional
//   	distribution: jsii.String("distribution"),
//   	filterName: jsii.String("filterName"),
//   	roleArn: jsii.String("roleArn"),
//   }
//
type CfnSubscriptionFilterProps struct {
	// The Amazon Resource Name (ARN) of the destination.
	DestinationArn *string `field:"required" json:"destinationArn" yaml:"destinationArn"`
	// The filtering expressions that restrict what gets delivered to the destination AWS resource.
	//
	// For more information about the filter pattern syntax, see [Filter and Pattern Syntax](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/FilterAndPatternSyntax.html) .
	FilterPattern *string `field:"required" json:"filterPattern" yaml:"filterPattern"`
	// The log group to associate with the subscription filter.
	//
	// All log events that are uploaded to this log group are filtered and delivered to the specified AWS resource if the filter pattern matches the log events.
	LogGroupName *string `field:"required" json:"logGroupName" yaml:"logGroupName"`
	// `AWS::Logs::SubscriptionFilter.Distribution`.
	Distribution *string `field:"optional" json:"distribution" yaml:"distribution"`
	// `AWS::Logs::SubscriptionFilter.FilterName`.
	FilterName *string `field:"optional" json:"filterName" yaml:"filterName"`
	// The ARN of an IAM role that grants CloudWatch Logs permissions to deliver ingested log events to the destination stream.
	//
	// You don't need to provide the ARN when you are working with a logical destination for cross-account delivery.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

