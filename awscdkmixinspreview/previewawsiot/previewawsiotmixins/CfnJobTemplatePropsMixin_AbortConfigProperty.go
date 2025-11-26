package previewawsiotmixins


// The criteria that determine when and how a job abort takes place.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   abortConfigProperty := &AbortConfigProperty{
//   	CriteriaList: []interface{}{
//   		&AbortCriteriaProperty{
//   			Action: jsii.String("action"),
//   			FailureType: jsii.String("failureType"),
//   			MinNumberOfExecutedThings: jsii.Number(123),
//   			ThresholdPercentage: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-jobtemplate-abortconfig.html
//
type CfnJobTemplatePropsMixin_AbortConfigProperty struct {
	// The list of criteria that determine when and how to abort the job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-jobtemplate-abortconfig.html#cfn-iot-jobtemplate-abortconfig-criterialist
	//
	CriteriaList interface{} `field:"optional" json:"criteriaList" yaml:"criteriaList"`
}

