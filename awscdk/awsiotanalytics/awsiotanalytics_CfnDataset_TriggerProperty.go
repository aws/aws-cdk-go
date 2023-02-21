package awsiotanalytics


// The "DatasetTrigger" that specifies when the data set is automatically updated.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   triggerProperty := &triggerProperty{
//   	schedule: &scheduleProperty{
//   		scheduleExpression: jsii.String("scheduleExpression"),
//   	},
//   	triggeringDataset: &triggeringDatasetProperty{
//   		datasetName: jsii.String("datasetName"),
//   	},
//   }
//
type CfnDataset_TriggerProperty struct {
	// The "Schedule" when the trigger is initiated.
	Schedule interface{} `field:"optional" json:"schedule" yaml:"schedule"`
	// Information about the data set whose content generation triggers the new data set content generation.
	TriggeringDataset interface{} `field:"optional" json:"triggeringDataset" yaml:"triggeringDataset"`
}

