package awslogs


// Properties for defining a `CfnSubscriptionFilter`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSubscriptionFilterProps := &CfnSubscriptionFilterProps{
//   	DestinationArn: jsii.String("destinationArn"),
//   	FilterPattern: jsii.String("filterPattern"),
//   	LogGroupName: jsii.String("logGroupName"),
//
//   	// the properties below are optional
//   	Distribution: jsii.String("distribution"),
//   	FilterName: jsii.String("filterName"),
//   	RoleArn: jsii.String("roleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-subscriptionfilter.html
//
type CfnSubscriptionFilterProps struct {
	// The Amazon Resource Name (ARN) of the destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-subscriptionfilter.html#cfn-logs-subscriptionfilter-destinationarn
	//
	DestinationArn *string `field:"required" json:"destinationArn" yaml:"destinationArn"`
	// The filtering expressions that restrict what gets delivered to the destination AWS resource.
	//
	// For more information about the filter pattern syntax, see [Filter and Pattern Syntax](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/FilterAndPatternSyntax.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-subscriptionfilter.html#cfn-logs-subscriptionfilter-filterpattern
	//
	FilterPattern *string `field:"required" json:"filterPattern" yaml:"filterPattern"`
	// The log group to associate with the subscription filter.
	//
	// All log events that are uploaded to this log group are filtered and delivered to the specified AWS resource if the filter pattern matches the log events.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-subscriptionfilter.html#cfn-logs-subscriptionfilter-loggroupname
	//
	LogGroupName *string `field:"required" json:"logGroupName" yaml:"logGroupName"`
	// The method used to distribute log data to the destination, which can be either random or grouped by log stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-subscriptionfilter.html#cfn-logs-subscriptionfilter-distribution
	//
	Distribution *string `field:"optional" json:"distribution" yaml:"distribution"`
	// The name of the subscription filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-subscriptionfilter.html#cfn-logs-subscriptionfilter-filtername
	//
	FilterName *string `field:"optional" json:"filterName" yaml:"filterName"`
	// The ARN of an IAM role that grants CloudWatch Logs permissions to deliver ingested log events to the destination stream.
	//
	// You don't need to provide the ARN when you are working with a logical destination for cross-account delivery.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-subscriptionfilter.html#cfn-logs-subscriptionfilter-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

