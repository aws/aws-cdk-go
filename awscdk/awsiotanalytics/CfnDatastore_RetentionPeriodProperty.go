package awsiotanalytics


// How long, in days, message data is kept.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   retentionPeriodProperty := &RetentionPeriodProperty{
//   	NumberOfDays: jsii.Number(123),
//   	Unlimited: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-datastore-retentionperiod.html
//
type CfnDatastore_RetentionPeriodProperty struct {
	// The number of days that message data is kept.
	//
	// The `unlimited` parameter must be false.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-datastore-retentionperiod.html#cfn-iotanalytics-datastore-retentionperiod-numberofdays
	//
	NumberOfDays *float64 `field:"optional" json:"numberOfDays" yaml:"numberOfDays"`
	// If true, message data is kept indefinitely.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-datastore-retentionperiod.html#cfn-iotanalytics-datastore-retentionperiod-unlimited
	//
	Unlimited interface{} `field:"optional" json:"unlimited" yaml:"unlimited"`
}

