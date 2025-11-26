package previewawsiotmixins


// The `MachineLearningDetectionConfig` property type controls confidence of the machine learning model.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   machineLearningDetectionConfigProperty := &MachineLearningDetectionConfigProperty{
//   	ConfidenceLevel: jsii.String("confidenceLevel"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-securityprofile-machinelearningdetectionconfig.html
//
type CfnSecurityProfilePropsMixin_MachineLearningDetectionConfigProperty struct {
	// The model confidence level.
	//
	// There are three levels of confidence, `"high"` , `"medium"` , and `"low"` .
	//
	// The higher the confidence level, the lower the sensitivity, and the lower the alarm frequency will be.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-securityprofile-machinelearningdetectionconfig.html#cfn-iot-securityprofile-machinelearningdetectionconfig-confidencelevel
	//
	ConfidenceLevel *string `field:"optional" json:"confidenceLevel" yaml:"confidenceLevel"`
}

