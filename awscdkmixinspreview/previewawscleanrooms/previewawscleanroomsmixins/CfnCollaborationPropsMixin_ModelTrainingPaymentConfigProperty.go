package previewawscleanroomsmixins


// An object representing the collaboration member's model training payment responsibilities set by the collaboration creator.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   modelTrainingPaymentConfigProperty := &ModelTrainingPaymentConfigProperty{
//   	IsResponsible: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-collaboration-modeltrainingpaymentconfig.html
//
type CfnCollaborationPropsMixin_ModelTrainingPaymentConfigProperty struct {
	// Indicates whether the collaboration creator has configured the collaboration member to pay for model training costs ( `TRUE` ) or has not configured the collaboration member to pay for model training costs ( `FALSE` ).
	//
	// Exactly one member can be configured to pay for model training costs. An error is returned if the collaboration creator sets a `TRUE` value for more than one member in the collaboration.
	//
	// If the collaboration creator hasn't specified anyone as the member paying for model training costs, then the member who can query is the default payer. An error is returned if the collaboration creator sets a `FALSE` value for the member who can query.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-collaboration-modeltrainingpaymentconfig.html#cfn-cleanrooms-collaboration-modeltrainingpaymentconfig-isresponsible
	//
	IsResponsible interface{} `field:"optional" json:"isResponsible" yaml:"isResponsible"`
}

