package awslex


// Specifies the DTMF input specifications.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dTMFSpecificationProperty := &DTMFSpecificationProperty{
//   	DeletionCharacter: jsii.String("deletionCharacter"),
//   	EndCharacter: jsii.String("endCharacter"),
//   	EndTimeoutMs: jsii.Number(123),
//   	MaxLength: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-dtmfspecification.html
//
type CfnBot_DTMFSpecificationProperty struct {
	// The DTMF character that clears the accumulated DTMF digits and immediately ends the input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-dtmfspecification.html#cfn-lex-bot-dtmfspecification-deletioncharacter
	//
	DeletionCharacter *string `field:"required" json:"deletionCharacter" yaml:"deletionCharacter"`
	// The DTMF character that immediately ends input.
	//
	// If the user does not press this character, the input ends after the end timeout.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-dtmfspecification.html#cfn-lex-bot-dtmfspecification-endcharacter
	//
	EndCharacter *string `field:"required" json:"endCharacter" yaml:"endCharacter"`
	// How long the bot should wait after the last DTMF character input before assuming that the input has concluded.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-dtmfspecification.html#cfn-lex-bot-dtmfspecification-endtimeoutms
	//
	EndTimeoutMs *float64 `field:"required" json:"endTimeoutMs" yaml:"endTimeoutMs"`
	// The maximum number of DTMF digits allowed in an utterance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-dtmfspecification.html#cfn-lex-bot-dtmfspecification-maxlength
	//
	MaxLength *float64 `field:"required" json:"maxLength" yaml:"maxLength"`
}

