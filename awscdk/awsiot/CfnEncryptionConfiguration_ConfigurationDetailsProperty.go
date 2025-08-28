package awsiot


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   configurationDetailsProperty := &ConfigurationDetailsProperty{
//   	ConfigurationStatus: jsii.String("configurationStatus"),
//   	ErrorCode: jsii.String("errorCode"),
//   	ErrorMessage: jsii.String("errorMessage"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-encryptionconfiguration-configurationdetails.html
//
type CfnEncryptionConfiguration_ConfigurationDetailsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-encryptionconfiguration-configurationdetails.html#cfn-iot-encryptionconfiguration-configurationdetails-configurationstatus
	//
	ConfigurationStatus *string `field:"optional" json:"configurationStatus" yaml:"configurationStatus"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-encryptionconfiguration-configurationdetails.html#cfn-iot-encryptionconfiguration-configurationdetails-errorcode
	//
	ErrorCode *string `field:"optional" json:"errorCode" yaml:"errorCode"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-encryptionconfiguration-configurationdetails.html#cfn-iot-encryptionconfiguration-configurationdetails-errormessage
	//
	ErrorMessage *string `field:"optional" json:"errorMessage" yaml:"errorMessage"`
}

