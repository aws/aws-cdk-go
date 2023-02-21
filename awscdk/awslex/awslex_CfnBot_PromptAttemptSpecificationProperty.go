package awslex


// Specifies the settings on a prompt attempt.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   promptAttemptSpecificationProperty := &promptAttemptSpecificationProperty{
//   	allowedInputTypes: &allowedInputTypesProperty{
//   		allowAudioInput: jsii.Boolean(false),
//   		allowDtmfInput: jsii.Boolean(false),
//   	},
//
//   	// the properties below are optional
//   	allowInterrupt: jsii.Boolean(false),
//   	audioAndDtmfInputSpecification: &audioAndDTMFInputSpecificationProperty{
//   		startTimeoutMs: jsii.Number(123),
//
//   		// the properties below are optional
//   		audioSpecification: &audioSpecificationProperty{
//   			endTimeoutMs: jsii.Number(123),
//   			maxLengthMs: jsii.Number(123),
//   		},
//   		dtmfSpecification: &dTMFSpecificationProperty{
//   			deletionCharacter: jsii.String("deletionCharacter"),
//   			endCharacter: jsii.String("endCharacter"),
//   			endTimeoutMs: jsii.Number(123),
//   			maxLength: jsii.Number(123),
//   		},
//   	},
//   	textInputSpecification: &textInputSpecificationProperty{
//   		startTimeoutMs: jsii.Number(123),
//   	},
//   }
//
type CfnBot_PromptAttemptSpecificationProperty struct {
	// Indicates the allowed input types of the prompt attempt.
	AllowedInputTypes interface{} `field:"required" json:"allowedInputTypes" yaml:"allowedInputTypes"`
	// Indicates whether the user can interrupt a speech prompt attempt from the bot.
	AllowInterrupt interface{} `field:"optional" json:"allowInterrupt" yaml:"allowInterrupt"`
	// Specifies the settings on audio and DTMF input.
	AudioAndDtmfInputSpecification interface{} `field:"optional" json:"audioAndDtmfInputSpecification" yaml:"audioAndDtmfInputSpecification"`
	// Specifies the settings on text input.
	TextInputSpecification interface{} `field:"optional" json:"textInputSpecification" yaml:"textInputSpecification"`
}

