package awsinspectorv2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   periodicScanConfigurationProperty := &PeriodicScanConfigurationProperty{
//   	Frequency: jsii.String("frequency"),
//   	FrequencyExpression: jsii.String("frequencyExpression"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-codesecurityscanconfiguration-periodicscanconfiguration.html
//
type CfnCodeSecurityScanConfiguration_PeriodicScanConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-codesecurityscanconfiguration-periodicscanconfiguration.html#cfn-inspectorv2-codesecurityscanconfiguration-periodicscanconfiguration-frequency
	//
	Frequency *string `field:"optional" json:"frequency" yaml:"frequency"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-codesecurityscanconfiguration-periodicscanconfiguration.html#cfn-inspectorv2-codesecurityscanconfiguration-periodicscanconfiguration-frequencyexpression
	//
	FrequencyExpression *string `field:"optional" json:"frequencyExpression" yaml:"frequencyExpression"`
}

