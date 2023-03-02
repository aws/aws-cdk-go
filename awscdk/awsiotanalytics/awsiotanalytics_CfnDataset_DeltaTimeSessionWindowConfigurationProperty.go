package awsiotanalytics


// A structure that contains the configuration information of a delta time session window.
//
// [`DeltaTime`](https://docs.aws.amazon.com/iotanalytics/latest/APIReference/API_DeltaTime.html) specifies a time interval. You can use `DeltaTime` to create dataset contents with data that has arrived in the data store since the last execution. For an example of `DeltaTime` , see [Creating a SQL dataset with a delta window (CLI)](https://docs.aws.amazon.com/iotanalytics/latest/userguide/automate-create-dataset.html#automate-example6) in the *AWS IoT Analytics User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deltaTimeSessionWindowConfigurationProperty := &deltaTimeSessionWindowConfigurationProperty{
//   	timeoutInMinutes: jsii.Number(123),
//   }
//
type CfnDataset_DeltaTimeSessionWindowConfigurationProperty struct {
	// A time interval.
	//
	// You can use `timeoutInMinutes` so that AWS IoT Analytics can batch up late data notifications that have been generated since the last execution. AWS IoT Analytics sends one batch of notifications to Amazon CloudWatch Events at one time.
	//
	// For more information about how to write a timestamp expression, see [Date and Time Functions and Operators](https://docs.aws.amazon.com/https://prestodb.io/docs/0.172/functions/datetime.html) , in the *Presto 0.172 Documentation* .
	TimeoutInMinutes *float64 `field:"required" json:"timeoutInMinutes" yaml:"timeoutInMinutes"`
}

