package awsiot


// Parameters to define a mitigation action that changes the state of the CA certificate to inactive.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   updateCACertificateParamsProperty := &UpdateCACertificateParamsProperty{
//   	Action: jsii.String("action"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-mitigationaction-updatecacertificateparams.html
//
type CfnMitigationAction_UpdateCACertificateParamsProperty struct {
	// The action that you want to apply to the CA certificate.
	//
	// The only supported value is `DEACTIVATE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-mitigationaction-updatecacertificateparams.html#cfn-iot-mitigationaction-updatecacertificateparams-action
	//
	Action *string `field:"required" json:"action" yaml:"action"`
}

