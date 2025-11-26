package previewawsgroundstationmixins


// Defines decoding settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   decodeConfigProperty := &DecodeConfigProperty{
//   	UnvalidatedJson: jsii.String("unvalidatedJson"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-decodeconfig.html
//
type CfnConfigPropsMixin_DecodeConfigProperty struct {
	// The decoding settings are in JSON format and define a set of steps to perform to decode the data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-decodeconfig.html#cfn-groundstation-config-decodeconfig-unvalidatedjson
	//
	UnvalidatedJson *string `field:"optional" json:"unvalidatedJson" yaml:"unvalidatedJson"`
}

