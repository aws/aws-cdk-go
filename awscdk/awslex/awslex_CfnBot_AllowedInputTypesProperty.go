package awslex


// Specifies the allowed input types.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   allowedInputTypesProperty := &allowedInputTypesProperty{
//   	allowAudioInput: jsii.Boolean(false),
//   	allowDtmfInput: jsii.Boolean(false),
//   }
//
type CfnBot_AllowedInputTypesProperty struct {
	// Indicates whether audio input is allowed.
	AllowAudioInput interface{} `field:"required" json:"allowAudioInput" yaml:"allowAudioInput"`
	// Indicates whether DTMF input is allowed.
	AllowDtmfInput interface{} `field:"required" json:"allowDtmfInput" yaml:"allowDtmfInput"`
}

