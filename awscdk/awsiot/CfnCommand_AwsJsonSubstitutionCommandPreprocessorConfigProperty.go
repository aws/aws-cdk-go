package awsiot


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   awsJsonSubstitutionCommandPreprocessorConfigProperty := &AwsJsonSubstitutionCommandPreprocessorConfigProperty{
//   	OutputFormat: jsii.String("outputFormat"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-command-awsjsonsubstitutioncommandpreprocessorconfig.html
//
type CfnCommand_AwsJsonSubstitutionCommandPreprocessorConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-command-awsjsonsubstitutioncommandpreprocessorconfig.html#cfn-iot-command-awsjsonsubstitutioncommandpreprocessorconfig-outputformat
	//
	OutputFormat *string `field:"required" json:"outputFormat" yaml:"outputFormat"`
}

