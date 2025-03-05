package awsiotcoredeviceadvisor


// Information of a test device.
//
// A thing ARN, certificate ARN or device role ARN is required.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deviceUnderTestProperty := &DeviceUnderTestProperty{
//   	CertificateArn: jsii.String("certificateArn"),
//   	ThingArn: jsii.String("thingArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotcoredeviceadvisor-suitedefinition-deviceundertest.html
//
type CfnSuiteDefinition_DeviceUnderTestProperty struct {
	// Lists device's certificate ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotcoredeviceadvisor-suitedefinition-deviceundertest.html#cfn-iotcoredeviceadvisor-suitedefinition-deviceundertest-certificatearn
	//
	CertificateArn *string `field:"optional" json:"certificateArn" yaml:"certificateArn"`
	// Lists device's thing ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotcoredeviceadvisor-suitedefinition-deviceundertest.html#cfn-iotcoredeviceadvisor-suitedefinition-deviceundertest-thingarn
	//
	ThingArn *string `field:"optional" json:"thingArn" yaml:"thingArn"`
}

