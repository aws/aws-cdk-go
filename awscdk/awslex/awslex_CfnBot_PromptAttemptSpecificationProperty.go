package awslex


// Specifies the settings on a prompt attempt.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   promptAttemptSpecificationProperty := &PromptAttemptSpecificationProperty{
//   	AllowedInputTypes: &AllowedInputTypesProperty{
//   		AllowAudioInput: jsii.Boolean(false),
//   		AllowDtmfInput: jsii.Boolean(false),
//   	},
//
//   	// the properties below are optional
//   	AllowInterrupt: jsii.Boolean(false),
//   	AudioAndDtmfInputSpecification: &AudioAndDTMFInputSpecificationProperty{
//   		StartTimeoutMs: jsii.Number(123),
//
//   		// the properties below are optional
//   		AudioSpecification: &AudioSpecificationProperty{
//   			EndTimeoutMs: jsii.Number(123),
//   			MaxLengthMs: jsii.Number(123),
//   		},
//   		DtmfSpecification: &DTMFSpecificationProperty{
//   			DeletionCharacter: jsii.String("deletionCharacter"),
//   			EndCharacter: jsii.String("endCharacter"),
//   			EndTimeoutMs: jsii.Number(123),
//   			MaxLength: jsii.Number(123),
//   		},
//   	},
//   	TextInputSpecification: &TextInputSpecificationProperty{
//   		StartTimeoutMs: jsii.Number(123),
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

