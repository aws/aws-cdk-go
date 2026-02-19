package awslogs


// Properties for defining a `CfnScheduledQuery`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnScheduledQueryProps := &CfnScheduledQueryProps{
//   	ExecutionRoleArn: jsii.String("executionRoleArn"),
//   	Name: jsii.String("name"),
//   	QueryLanguage: jsii.String("queryLanguage"),
//   	QueryString: jsii.String("queryString"),
//   	ScheduleExpression: jsii.String("scheduleExpression"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	DestinationConfiguration: &DestinationConfigurationProperty{
//   		S3Configuration: &S3ConfigurationProperty{
//   			DestinationIdentifier: jsii.String("destinationIdentifier"),
//   			RoleArn: jsii.String("roleArn"),
//   		},
//   	},
//   	LogGroupIdentifiers: []*string{
//   		jsii.String("logGroupIdentifiers"),
//   	},
//   	ScheduleEndTime: jsii.Number(123),
//   	ScheduleStartTime: jsii.Number(123),
//   	StartTimeOffset: jsii.Number(123),
//   	State: jsii.String("state"),
//   	Tags: []TagsItemsProperty{
//   		&TagsItemsProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Timezone: jsii.String("timezone"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-scheduledquery.html
//
type CfnScheduledQueryProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-scheduledquery.html#cfn-logs-scheduledquery-executionrolearn
	//
	ExecutionRoleArn *string `field:"required" json:"executionRoleArn" yaml:"executionRoleArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-scheduledquery.html#cfn-logs-scheduledquery-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-scheduledquery.html#cfn-logs-scheduledquery-querylanguage
	//
	QueryLanguage *string `field:"required" json:"queryLanguage" yaml:"queryLanguage"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-scheduledquery.html#cfn-logs-scheduledquery-querystring
	//
	QueryString *string `field:"required" json:"queryString" yaml:"queryString"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-scheduledquery.html#cfn-logs-scheduledquery-scheduleexpression
	//
	ScheduleExpression *string `field:"required" json:"scheduleExpression" yaml:"scheduleExpression"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-scheduledquery.html#cfn-logs-scheduledquery-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-scheduledquery.html#cfn-logs-scheduledquery-destinationconfiguration
	//
	DestinationConfiguration interface{} `field:"optional" json:"destinationConfiguration" yaml:"destinationConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-scheduledquery.html#cfn-logs-scheduledquery-loggroupidentifiers
	//
	LogGroupIdentifiers *[]*string `field:"optional" json:"logGroupIdentifiers" yaml:"logGroupIdentifiers"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-scheduledquery.html#cfn-logs-scheduledquery-scheduleendtime
	//
	ScheduleEndTime *float64 `field:"optional" json:"scheduleEndTime" yaml:"scheduleEndTime"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-scheduledquery.html#cfn-logs-scheduledquery-schedulestarttime
	//
	ScheduleStartTime *float64 `field:"optional" json:"scheduleStartTime" yaml:"scheduleStartTime"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-scheduledquery.html#cfn-logs-scheduledquery-starttimeoffset
	//
	StartTimeOffset *float64 `field:"optional" json:"startTimeOffset" yaml:"startTimeOffset"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-scheduledquery.html#cfn-logs-scheduledquery-state
	//
	State *string `field:"optional" json:"state" yaml:"state"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-scheduledquery.html#cfn-logs-scheduledquery-tags
	//
	Tags *[]*CfnScheduledQuery_TagsItemsProperty `field:"optional" json:"tags" yaml:"tags"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-scheduledquery.html#cfn-logs-scheduledquery-timezone
	//
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
}

