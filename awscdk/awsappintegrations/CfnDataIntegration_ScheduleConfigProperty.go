package awsappintegrations


// The name of the data and how often it should be pulled from the source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scheduleConfigProperty := &ScheduleConfigProperty{
//   	ScheduleExpression: jsii.String("scheduleExpression"),
//
//   	// the properties below are optional
//   	FirstExecutionFrom: jsii.String("firstExecutionFrom"),
//   	Object: jsii.String("object"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appintegrations-dataintegration-scheduleconfig.html
//
type CfnDataIntegration_ScheduleConfigProperty struct {
	// How often the data should be pulled from data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appintegrations-dataintegration-scheduleconfig.html#cfn-appintegrations-dataintegration-scheduleconfig-scheduleexpression
	//
	ScheduleExpression *string `field:"required" json:"scheduleExpression" yaml:"scheduleExpression"`
	// The start date for objects to import in the first flow run as an Unix/epoch timestamp in milliseconds or in ISO-8601 format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appintegrations-dataintegration-scheduleconfig.html#cfn-appintegrations-dataintegration-scheduleconfig-firstexecutionfrom
	//
	FirstExecutionFrom *string `field:"optional" json:"firstExecutionFrom" yaml:"firstExecutionFrom"`
	// The name of the object to pull from the data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appintegrations-dataintegration-scheduleconfig.html#cfn-appintegrations-dataintegration-scheduleconfig-object
	//
	Object *string `field:"optional" json:"object" yaml:"object"`
}

