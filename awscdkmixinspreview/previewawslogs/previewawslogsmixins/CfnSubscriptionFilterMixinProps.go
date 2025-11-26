package previewawslogsmixins


// Properties for CfnSubscriptionFilterPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnSubscriptionFilterMixinProps := &CfnSubscriptionFilterMixinProps{
//   	ApplyOnTransformedLogs: jsii.Boolean(false),
//   	DestinationArn: jsii.String("destinationArn"),
//   	Distribution: jsii.String("distribution"),
//   	EmitSystemFields: []*string{
//   		jsii.String("emitSystemFields"),
//   	},
//   	FieldSelectionCriteria: jsii.String("fieldSelectionCriteria"),
//   	FilterName: jsii.String("filterName"),
//   	FilterPattern: jsii.String("filterPattern"),
//   	LogGroupName: jsii.String("logGroupName"),
//   	RoleArn: jsii.String("roleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-subscriptionfilter.html
//
type CfnSubscriptionFilterMixinProps struct {
	// This parameter is valid only for log groups that have an active log transformer.
	//
	// For more information about log transformers, see [PutTransformer](https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutTransformer.html) .
	//
	// If this value is `true` , the subscription filter is applied on the transformed version of the log events instead of the original ingested log events.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-subscriptionfilter.html#cfn-logs-subscriptionfilter-applyontransformedlogs
	//
	ApplyOnTransformedLogs interface{} `field:"optional" json:"applyOnTransformedLogs" yaml:"applyOnTransformedLogs"`
	// The Amazon Resource Name (ARN) of the destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-subscriptionfilter.html#cfn-logs-subscriptionfilter-destinationarn
	//
	DestinationArn *string `field:"optional" json:"destinationArn" yaml:"destinationArn"`
	// The method used to distribute log data to the destination, which can be either random or grouped by log stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-subscriptionfilter.html#cfn-logs-subscriptionfilter-distribution
	//
	Distribution *string `field:"optional" json:"distribution" yaml:"distribution"`
	// The list of system fields that are included in the log events sent to the subscription destination.
	//
	// Returns the `emitSystemFields` value if it was specified when the subscription filter was created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-subscriptionfilter.html#cfn-logs-subscriptionfilter-emitsystemfields
	//
	EmitSystemFields *[]*string `field:"optional" json:"emitSystemFields" yaml:"emitSystemFields"`
	// The filter expression that specifies which log events are processed by this subscription filter based on system fields.
	//
	// Returns the `fieldSelectionCriteria` value if it was specified when the subscription filter was created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-subscriptionfilter.html#cfn-logs-subscriptionfilter-fieldselectioncriteria
	//
	FieldSelectionCriteria *string `field:"optional" json:"fieldSelectionCriteria" yaml:"fieldSelectionCriteria"`
	// The name of the subscription filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-subscriptionfilter.html#cfn-logs-subscriptionfilter-filtername
	//
	FilterName *string `field:"optional" json:"filterName" yaml:"filterName"`
	// The filtering expressions that restrict what gets delivered to the destination AWS resource.
	//
	// For more information about the filter pattern syntax, see [Filter and Pattern Syntax](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/FilterAndPatternSyntax.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-subscriptionfilter.html#cfn-logs-subscriptionfilter-filterpattern
	//
	FilterPattern *string `field:"optional" json:"filterPattern" yaml:"filterPattern"`
	// The log group to associate with the subscription filter.
	//
	// All log events that are uploaded to this log group are filtered and delivered to the specified AWS resource if the filter pattern matches the log events.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-subscriptionfilter.html#cfn-logs-subscriptionfilter-loggroupname
	//
	LogGroupName *string `field:"optional" json:"logGroupName" yaml:"logGroupName"`
	// The ARN of an IAM role that grants CloudWatch Logs permissions to deliver ingested log events to the destination stream.
	//
	// You don't need to provide the ARN when you are working with a logical destination for cross-account delivery.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-subscriptionfilter.html#cfn-logs-subscriptionfilter-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

