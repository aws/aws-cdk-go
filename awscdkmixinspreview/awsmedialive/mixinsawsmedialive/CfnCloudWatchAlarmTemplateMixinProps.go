package mixinsawsmedialive


// Properties for CfnCloudWatchAlarmTemplatePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCloudWatchAlarmTemplateMixinProps := &CfnCloudWatchAlarmTemplateMixinProps{
//   	ComparisonOperator: jsii.String("comparisonOperator"),
//   	DatapointsToAlarm: jsii.Number(123),
//   	Description: jsii.String("description"),
//   	EvaluationPeriods: jsii.Number(123),
//   	GroupIdentifier: jsii.String("groupIdentifier"),
//   	MetricName: jsii.String("metricName"),
//   	Name: jsii.String("name"),
//   	Period: jsii.Number(123),
//   	Statistic: jsii.String("statistic"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	TargetResourceType: jsii.String("targetResourceType"),
//   	Threshold: jsii.Number(123),
//   	TreatMissingData: jsii.String("treatMissingData"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-cloudwatchalarmtemplate.html
//
type CfnCloudWatchAlarmTemplateMixinProps struct {
	// The comparison operator used to compare the specified statistic and the threshold.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-cloudwatchalarmtemplate.html#cfn-medialive-cloudwatchalarmtemplate-comparisonoperator
	//
	ComparisonOperator *string `field:"optional" json:"comparisonOperator" yaml:"comparisonOperator"`
	// The number of datapoints within the evaluation period that must be breaching to trigger the alarm.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-cloudwatchalarmtemplate.html#cfn-medialive-cloudwatchalarmtemplate-datapointstoalarm
	//
	// Default: - 0.
	//
	DatapointsToAlarm *float64 `field:"optional" json:"datapointsToAlarm" yaml:"datapointsToAlarm"`
	// A resource's optional description.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-cloudwatchalarmtemplate.html#cfn-medialive-cloudwatchalarmtemplate-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The number of periods over which data is compared to the specified threshold.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-cloudwatchalarmtemplate.html#cfn-medialive-cloudwatchalarmtemplate-evaluationperiods
	//
	// Default: - 0.
	//
	EvaluationPeriods *float64 `field:"optional" json:"evaluationPeriods" yaml:"evaluationPeriods"`
	// A cloudwatch alarm template group's identifier.
	//
	// Can be either be its id or current name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-cloudwatchalarmtemplate.html#cfn-medialive-cloudwatchalarmtemplate-groupidentifier
	//
	GroupIdentifier *string `field:"optional" json:"groupIdentifier" yaml:"groupIdentifier"`
	// The name of the metric associated with the alarm.
	//
	// Must be compatible with targetResourceType.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-cloudwatchalarmtemplate.html#cfn-medialive-cloudwatchalarmtemplate-metricname
	//
	MetricName *string `field:"optional" json:"metricName" yaml:"metricName"`
	// A resource's name.
	//
	// Names must be unique within the scope of a resource type in a specific region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-cloudwatchalarmtemplate.html#cfn-medialive-cloudwatchalarmtemplate-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The period, in seconds, over which the specified statistic is applied.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-cloudwatchalarmtemplate.html#cfn-medialive-cloudwatchalarmtemplate-period
	//
	// Default: - 0.
	//
	Period *float64 `field:"optional" json:"period" yaml:"period"`
	// The statistic to apply to the alarm's metric data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-cloudwatchalarmtemplate.html#cfn-medialive-cloudwatchalarmtemplate-statistic
	//
	Statistic *string `field:"optional" json:"statistic" yaml:"statistic"`
	// Represents the tags associated with a resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-cloudwatchalarmtemplate.html#cfn-medialive-cloudwatchalarmtemplate-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The resource type this template should dynamically generate CloudWatch metric alarms for.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-cloudwatchalarmtemplate.html#cfn-medialive-cloudwatchalarmtemplate-targetresourcetype
	//
	TargetResourceType *string `field:"optional" json:"targetResourceType" yaml:"targetResourceType"`
	// The threshold value to compare with the specified statistic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-cloudwatchalarmtemplate.html#cfn-medialive-cloudwatchalarmtemplate-threshold
	//
	// Default: - 0.
	//
	Threshold *float64 `field:"optional" json:"threshold" yaml:"threshold"`
	// Specifies how missing data points are treated when evaluating the alarm's condition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-cloudwatchalarmtemplate.html#cfn-medialive-cloudwatchalarmtemplate-treatmissingdata
	//
	TreatMissingData *string `field:"optional" json:"treatMissingData" yaml:"treatMissingData"`
}

