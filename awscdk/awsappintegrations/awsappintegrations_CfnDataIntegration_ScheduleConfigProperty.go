package awsappintegrations


// The name of the data and how often it should be pulled from the source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scheduleConfigProperty := &scheduleConfigProperty{
//   	firstExecutionFrom: jsii.String("firstExecutionFrom"),
//   	object: jsii.String("object"),
//   	scheduleExpression: jsii.String("scheduleExpression"),
//   }
//
type CfnDataIntegration_ScheduleConfigProperty struct {
	// The start date for objects to import in the first flow run as an Unix/epoch timestamp in milliseconds or in ISO-8601 format.
	FirstExecutionFrom *string `field:"required" json:"firstExecutionFrom" yaml:"firstExecutionFrom"`
	// The name of the object to pull from the data source.
	Object *string `field:"required" json:"object" yaml:"object"`
	// How often the data should be pulled from data source.
	ScheduleExpression *string `field:"required" json:"scheduleExpression" yaml:"scheduleExpression"`
}

