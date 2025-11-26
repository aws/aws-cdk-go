package previewawsiotmixins


// The criteria that determines how many retries are allowed for each failure type for a job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   retryCriteriaProperty := &RetryCriteriaProperty{
//   	FailureType: jsii.String("failureType"),
//   	NumberOfRetries: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-jobtemplate-retrycriteria.html
//
type CfnJobTemplatePropsMixin_RetryCriteriaProperty struct {
	// The type of job execution failures that can initiate a job retry.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-jobtemplate-retrycriteria.html#cfn-iot-jobtemplate-retrycriteria-failuretype
	//
	FailureType *string `field:"optional" json:"failureType" yaml:"failureType"`
	// The number of retries allowed for a failure type for the job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-jobtemplate-retrycriteria.html#cfn-iot-jobtemplate-retrycriteria-numberofretries
	//
	NumberOfRetries *float64 `field:"optional" json:"numberOfRetries" yaml:"numberOfRetries"`
}

