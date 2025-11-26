package previewawsgreengrassv2mixins


// Contains a list of criteria that define when and how to cancel a configuration deployment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ioTJobAbortConfigProperty := &IoTJobAbortConfigProperty{
//   	CriteriaList: []interface{}{
//   		&IoTJobAbortCriteriaProperty{
//   			Action: jsii.String("action"),
//   			FailureType: jsii.String("failureType"),
//   			MinNumberOfExecutedThings: jsii.Number(123),
//   			ThresholdPercentage: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-deployment-iotjobabortconfig.html
//
type CfnDeploymentPropsMixin_IoTJobAbortConfigProperty struct {
	// The list of criteria that define when and how to cancel the configuration deployment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-deployment-iotjobabortconfig.html#cfn-greengrassv2-deployment-iotjobabortconfig-criterialist
	//
	CriteriaList interface{} `field:"optional" json:"criteriaList" yaml:"criteriaList"`
}

