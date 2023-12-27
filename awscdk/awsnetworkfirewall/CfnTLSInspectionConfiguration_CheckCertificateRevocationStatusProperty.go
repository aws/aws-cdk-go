package awsnetworkfirewall


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   checkCertificateRevocationStatusProperty := &CheckCertificateRevocationStatusProperty{
//   	RevokedStatusAction: jsii.String("revokedStatusAction"),
//   	UnknownStatusAction: jsii.String("unknownStatusAction"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-tlsinspectionconfiguration-checkcertificaterevocationstatus.html
//
type CfnTLSInspectionConfiguration_CheckCertificateRevocationStatusProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-tlsinspectionconfiguration-checkcertificaterevocationstatus.html#cfn-networkfirewall-tlsinspectionconfiguration-checkcertificaterevocationstatus-revokedstatusaction
	//
	RevokedStatusAction *string `field:"optional" json:"revokedStatusAction" yaml:"revokedStatusAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-tlsinspectionconfiguration-checkcertificaterevocationstatus.html#cfn-networkfirewall-tlsinspectionconfiguration-checkcertificaterevocationstatus-unknownstatusaction
	//
	UnknownStatusAction *string `field:"optional" json:"unknownStatusAction" yaml:"unknownStatusAction"`
}

