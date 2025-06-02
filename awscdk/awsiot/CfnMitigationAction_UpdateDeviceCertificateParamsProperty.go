package awsiot


// Parameters to define a mitigation action that changes the state of the device certificate to inactive.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   updateDeviceCertificateParamsProperty := &UpdateDeviceCertificateParamsProperty{
//   	Action: jsii.String("action"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-mitigationaction-updatedevicecertificateparams.html
//
type CfnMitigationAction_UpdateDeviceCertificateParamsProperty struct {
	// The action that you want to apply to the device certificate.
	//
	// The only supported value is `DEACTIVATE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-mitigationaction-updatedevicecertificateparams.html#cfn-iot-mitigationaction-updatedevicecertificateparams-action
	//
	Action *string `field:"required" json:"action" yaml:"action"`
}

