package previewawsiotmixins


// The configuration that determines how many retries are allowed for each failure type for a job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   jobExecutionsRetryConfigProperty := &JobExecutionsRetryConfigProperty{
//   	RetryCriteriaList: []interface{}{
//   		&RetryCriteriaProperty{
//   			FailureType: jsii.String("failureType"),
//   			NumberOfRetries: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-jobtemplate-jobexecutionsretryconfig.html
//
type CfnJobTemplatePropsMixin_JobExecutionsRetryConfigProperty struct {
	// The list of criteria that determines how many retries are allowed for each failure type for a job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-jobtemplate-jobexecutionsretryconfig.html#cfn-iot-jobtemplate-jobexecutionsretryconfig-retrycriterialist
	//
	RetryCriteriaList interface{} `field:"optional" json:"retryCriteriaList" yaml:"retryCriteriaList"`
}

