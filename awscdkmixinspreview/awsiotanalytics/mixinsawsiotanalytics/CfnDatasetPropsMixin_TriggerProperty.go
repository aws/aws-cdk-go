package mixinsawsiotanalytics


// The "DatasetTrigger" that specifies when the data set is automatically updated.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   triggerProperty := &TriggerProperty{
//   	Schedule: &ScheduleProperty{
//   		ScheduleExpression: jsii.String("scheduleExpression"),
//   	},
//   	TriggeringDataset: &TriggeringDatasetProperty{
//   		DatasetName: jsii.String("datasetName"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-trigger.html
//
type CfnDatasetPropsMixin_TriggerProperty struct {
	// The "Schedule" when the trigger is initiated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-trigger.html#cfn-iotanalytics-dataset-trigger-schedule
	//
	Schedule interface{} `field:"optional" json:"schedule" yaml:"schedule"`
	// Information about the data set whose content generation triggers the new data set content generation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-trigger.html#cfn-iotanalytics-dataset-trigger-triggeringdataset
	//
	TriggeringDataset interface{} `field:"optional" json:"triggeringDataset" yaml:"triggeringDataset"`
}

