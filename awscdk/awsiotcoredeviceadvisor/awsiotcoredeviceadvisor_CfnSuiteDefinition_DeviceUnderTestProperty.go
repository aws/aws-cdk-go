package awsiotcoredeviceadvisor


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
	// `CfnSuiteDefinition.DeviceUnderTestProperty.CertificateArn`.
	CertificateArn *string `field:"optional" json:"certificateArn" yaml:"certificateArn"`
	// `CfnSuiteDefinition.DeviceUnderTestProperty.ThingArn`.
	ThingArn *string `field:"optional" json:"thingArn" yaml:"thingArn"`
}

