package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   statusProperty := &StatusProperty{
//   	Message: jsii.String("message"),
//   	PrimaryStatus: jsii.String("primaryStatus"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-experimenttrialcomponent-status.html
//
type CfnExperimentTrialComponentPropsMixin_StatusProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-experimenttrialcomponent-status.html#cfn-sagemaker-experimenttrialcomponent-status-message
	//
	Message *string `field:"optional" json:"message" yaml:"message"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-experimenttrialcomponent-status.html#cfn-sagemaker-experimenttrialcomponent-status-primarystatus
	//
	PrimaryStatus *string `field:"optional" json:"primaryStatus" yaml:"primaryStatus"`
}

