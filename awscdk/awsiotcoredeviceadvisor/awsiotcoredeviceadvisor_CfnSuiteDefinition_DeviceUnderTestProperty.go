package awsiotcoredeviceadvisor


// Information of a test device.
//
// A thing ARN or a certificate ARN is required.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deviceUnderTestProperty := &deviceUnderTestProperty{
//   	certificateArn: jsii.String("certificateArn"),
//   	thingArn: jsii.String("thingArn"),
//   }
//
type CfnSuiteDefinition_DeviceUnderTestProperty struct {
	// Lists devices certificate ARN.
	CertificateArn *string `field:"optional" json:"certificateArn" yaml:"certificateArn"`
	// Lists devices thing ARN.
	ThingArn *string `field:"optional" json:"thingArn" yaml:"thingArn"`
}

