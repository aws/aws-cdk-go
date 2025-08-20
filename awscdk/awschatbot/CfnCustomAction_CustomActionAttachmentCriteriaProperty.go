package awschatbot


// > AWS Chatbot is now  . [Learn more](https://docs.aws.amazon.com//chatbot/latest/adminguide/service-rename.html) >  > `Type` attribute values remain unchanged.
//
// A criteria for when a button should be shown based on values in the notification.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customActionAttachmentCriteriaProperty := &CustomActionAttachmentCriteriaProperty{
//   	Operator: jsii.String("operator"),
//   	VariableName: jsii.String("variableName"),
//
//   	// the properties below are optional
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chatbot-customaction-customactionattachmentcriteria.html
//
type CfnCustomAction_CustomActionAttachmentCriteriaProperty struct {
	// The operation to perform on the named variable.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chatbot-customaction-customactionattachmentcriteria.html#cfn-chatbot-customaction-customactionattachmentcriteria-operator
	//
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// The name of the variable to operate on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chatbot-customaction-customactionattachmentcriteria.html#cfn-chatbot-customaction-customactionattachmentcriteria-variablename
	//
	VariableName *string `field:"required" json:"variableName" yaml:"variableName"`
	// A value that is compared with the actual value of the variable based on the behavior of the operator.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chatbot-customaction-customactionattachmentcriteria.html#cfn-chatbot-customaction-customactionattachmentcriteria-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

