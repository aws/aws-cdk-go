package awslex


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
	// `CfnBot.AllowedInputTypesProperty.AllowAudioInput`.
	AllowAudioInput interface{} `field:"required" json:"allowAudioInput" yaml:"allowAudioInput"`
	// `CfnBot.AllowedInputTypesProperty.AllowDTMFInput`.
	AllowDtmfInput interface{} `field:"required" json:"allowDtmfInput" yaml:"allowDtmfInput"`
}

