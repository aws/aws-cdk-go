package awslex


// Specifies the audio and DTMF input specification.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   audioAndDTMFInputSpecificationProperty := &audioAndDTMFInputSpecificationProperty{
//   	startTimeoutMs: jsii.Number(123),
//
//   	// the properties below are optional
//   	audioSpecification: &audioSpecificationProperty{
//   		endTimeoutMs: jsii.Number(123),
//   		maxLengthMs: jsii.Number(123),
//   	},
//   	dtmfSpecification: &dTMFSpecificationProperty{
//   		deletionCharacter: jsii.String("deletionCharacter"),
//   		endCharacter: jsii.String("endCharacter"),
//   		endTimeoutMs: jsii.Number(123),
//   		maxLength: jsii.Number(123),
//   	},
//   }
//
type CfnBot_AudioAndDTMFInputSpecificationProperty struct {
	// Time for which a bot waits before assuming that the customer isn't going to speak or press a key.
	//
	// This timeout is shared between Audio and DTMF inputs.
	StartTimeoutMs *float64 `field:"required" json:"startTimeoutMs" yaml:"startTimeoutMs"`
	// Specifies the settings on audio input.
	AudioSpecification interface{} `field:"optional" json:"audioSpecification" yaml:"audioSpecification"`
	// Specifies the settings on DTMF input.
	DtmfSpecification interface{} `field:"optional" json:"dtmfSpecification" yaml:"dtmfSpecification"`
}

