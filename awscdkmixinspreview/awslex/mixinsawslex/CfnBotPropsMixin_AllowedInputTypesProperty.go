package mixinsawslex


// Specifies the allowed input types.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   allowedInputTypesProperty := &AllowedInputTypesProperty{
//   	AllowAudioInput: jsii.Boolean(false),
//   	AllowDtmfInput: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-allowedinputtypes.html
//
type CfnBotPropsMixin_AllowedInputTypesProperty struct {
	// Indicates whether audio input is allowed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-allowedinputtypes.html#cfn-lex-bot-allowedinputtypes-allowaudioinput
	//
	AllowAudioInput interface{} `field:"optional" json:"allowAudioInput" yaml:"allowAudioInput"`
	// Indicates whether DTMF input is allowed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-allowedinputtypes.html#cfn-lex-bot-allowedinputtypes-allowdtmfinput
	//
	AllowDtmfInput interface{} `field:"optional" json:"allowDtmfInput" yaml:"allowDtmfInput"`
}

