package awsiot


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   awsJobAbortConfigProperty := &AwsJobAbortConfigProperty{
//   	AbortCriteriaList: []interface{}{
//   		&AwsJobAbortCriteriaProperty{
//   			Action: jsii.String("action"),
//   			FailureType: jsii.String("failureType"),
//   			MinNumberOfExecutedThings: jsii.Number(123),
//   			ThresholdPercentage: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-otaupdate-awsjobabortconfig.html
//
type CfnOTAUpdate_AwsJobAbortConfigProperty struct {
	// The list of criteria that determine when and how to abort the job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-otaupdate-awsjobabortconfig.html#cfn-iot-otaupdate-awsjobabortconfig-abortcriterialist
	//
	AbortCriteriaList interface{} `field:"required" json:"abortCriteriaList" yaml:"abortCriteriaList"`
}

