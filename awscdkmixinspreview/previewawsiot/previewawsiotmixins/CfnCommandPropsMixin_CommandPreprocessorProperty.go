package previewawsiotmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   commandPreprocessorProperty := &CommandPreprocessorProperty{
//   	AwsJsonSubstitution: &AwsJsonSubstitutionCommandPreprocessorConfigProperty{
//   		OutputFormat: jsii.String("outputFormat"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-command-commandpreprocessor.html
//
type CfnCommandPropsMixin_CommandPreprocessorProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-command-commandpreprocessor.html#cfn-iot-command-commandpreprocessor-awsjsonsubstitution
	//
	AwsJsonSubstitution interface{} `field:"optional" json:"awsJsonSubstitution" yaml:"awsJsonSubstitution"`
}

