package awsiot


// The encryption configuration details that include the status information of the AWS Key Management Service ( AWS  ) key and the AWS  access role.
//
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
	// The health status of KMS key and AWS  access role.
	//
	// If either KMS key or AWS  access role is `UNHEALTHY` , the return value will be `UNHEALTHY` . To use a customer managed KMS key, the value of `configurationStatus` must be `HEALTHY` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-encryptionconfiguration-configurationdetails.html#cfn-iot-encryptionconfiguration-configurationdetails-configurationstatus
	//
	ConfigurationStatus *string `field:"optional" json:"configurationStatus" yaml:"configurationStatus"`
	// The error code that indicates either the KMS key or the AWS  access role is `UNHEALTHY` .
	//
	// Valid values: `KMS_KEY_VALIDATION_ERROR` and `ROLE_VALIDATION_ERROR` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-encryptionconfiguration-configurationdetails.html#cfn-iot-encryptionconfiguration-configurationdetails-errorcode
	//
	ErrorCode *string `field:"optional" json:"errorCode" yaml:"errorCode"`
	// The detailed error message that corresponds to the `errorCode` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-encryptionconfiguration-configurationdetails.html#cfn-iot-encryptionconfiguration-configurationdetails-errormessage
	//
	ErrorMessage *string `field:"optional" json:"errorMessage" yaml:"errorMessage"`
}

