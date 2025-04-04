package awssagemaker


// An Amazon SageMaker Model Card.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   modelCardProperty := &ModelCardProperty{
//   	ModelCardContent: jsii.String("modelCardContent"),
//   	ModelCardStatus: jsii.String("modelCardStatus"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-modelcard.html
//
type CfnModelPackage_ModelCardProperty struct {
	// The content of the model card.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-modelcard.html#cfn-sagemaker-modelpackage-modelcard-modelcardcontent
	//
	ModelCardContent *string `field:"required" json:"modelCardContent" yaml:"modelCardContent"`
	// The approval status of the model card within your organization.
	//
	// Different organizations might have different criteria for model card review and approval.
	//
	// - `Draft` : The model card is a work in progress.
	// - `PendingReview` : The model card is pending review.
	// - `Approved` : The model card is approved.
	// - `Archived` : The model card is archived. No more updates should be made to the model card, but it can still be exported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-modelcard.html#cfn-sagemaker-modelpackage-modelcard-modelcardstatus
	//
	ModelCardStatus *string `field:"required" json:"modelCardStatus" yaml:"modelCardStatus"`
}

