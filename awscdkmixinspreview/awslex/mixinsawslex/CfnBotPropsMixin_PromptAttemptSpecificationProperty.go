package mixinsawslex


// Specifies the settings on a prompt attempt.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   promptAttemptSpecificationProperty := &PromptAttemptSpecificationProperty{
//   	AllowedInputTypes: &AllowedInputTypesProperty{
//   		AllowAudioInput: jsii.Boolean(false),
//   		AllowDtmfInput: jsii.Boolean(false),
//   	},
//   	AllowInterrupt: jsii.Boolean(false),
//   	AudioAndDtmfInputSpecification: &AudioAndDTMFInputSpecificationProperty{
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
//   		StartTimeoutMs: jsii.Number(123),
//   	},
//   	TextInputSpecification: &TextInputSpecificationProperty{
//   		StartTimeoutMs: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-promptattemptspecification.html
//
type CfnBotPropsMixin_PromptAttemptSpecificationProperty struct {
	// Indicates the allowed input types of the prompt attempt.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-promptattemptspecification.html#cfn-lex-bot-promptattemptspecification-allowedinputtypes
	//
	AllowedInputTypes interface{} `field:"optional" json:"allowedInputTypes" yaml:"allowedInputTypes"`
	// Indicates whether the user can interrupt a speech prompt attempt from the bot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-promptattemptspecification.html#cfn-lex-bot-promptattemptspecification-allowinterrupt
	//
	AllowInterrupt interface{} `field:"optional" json:"allowInterrupt" yaml:"allowInterrupt"`
	// Specifies the settings on audio and DTMF input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-promptattemptspecification.html#cfn-lex-bot-promptattemptspecification-audioanddtmfinputspecification
	//
	AudioAndDtmfInputSpecification interface{} `field:"optional" json:"audioAndDtmfInputSpecification" yaml:"audioAndDtmfInputSpecification"`
	// Specifies the settings on text input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-promptattemptspecification.html#cfn-lex-bot-promptattemptspecification-textinputspecification
	//
	TextInputSpecification interface{} `field:"optional" json:"textInputSpecification" yaml:"textInputSpecification"`
}

