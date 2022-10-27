package awslex


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
	// `CfnBot.AudioAndDTMFInputSpecificationProperty.StartTimeoutMs`.
	StartTimeoutMs *float64 `field:"required" json:"startTimeoutMs" yaml:"startTimeoutMs"`
	// `CfnBot.AudioAndDTMFInputSpecificationProperty.AudioSpecification`.
	AudioSpecification interface{} `field:"optional" json:"audioSpecification" yaml:"audioSpecification"`
	// `CfnBot.AudioAndDTMFInputSpecificationProperty.DTMFSpecification`.
	DtmfSpecification interface{} `field:"optional" json:"dtmfSpecification" yaml:"dtmfSpecification"`
}

