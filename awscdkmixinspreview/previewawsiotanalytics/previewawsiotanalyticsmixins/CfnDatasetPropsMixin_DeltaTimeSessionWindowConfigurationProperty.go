package previewawsiotanalyticsmixins


// A structure that contains the configuration information of a delta time session window.
//
// [`DeltaTime`](https://docs.aws.amazon.com/iotanalytics/latest/APIReference/API_DeltaTime.html) specifies a time interval. You can use `DeltaTime` to create dataset contents with data that has arrived in the data store since the last execution. For an example of `DeltaTime` , see [Creating a SQL dataset with a delta window (CLI)](https://docs.aws.amazon.com/iotanalytics/latest/userguide/automate-create-dataset.html#automate-example6) in the *ITA User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   deltaTimeSessionWindowConfigurationProperty := &DeltaTimeSessionWindowConfigurationProperty{
//   	TimeoutInMinutes: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-deltatimesessionwindowconfiguration.html
//
type CfnDatasetPropsMixin_DeltaTimeSessionWindowConfigurationProperty struct {
	// A time interval.
	//
	// You can use `timeoutInMinutes` so that ITA can batch up late data notifications that have been generated since the last execution. ITA sends one batch of notifications to Amazon CloudWatch Events at one time.
	//
	// For more information about how to write a timestamp expression, see [Date and Time Functions and Operators](https://docs.aws.amazon.com/https://prestodb.io/docs/current/functions/datetime.html) , in the *Presto 0.172 Documentation* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-deltatimesessionwindowconfiguration.html#cfn-iotanalytics-dataset-deltatimesessionwindowconfiguration-timeoutinminutes
	//
	TimeoutInMinutes *float64 `field:"optional" json:"timeoutInMinutes" yaml:"timeoutInMinutes"`
}

