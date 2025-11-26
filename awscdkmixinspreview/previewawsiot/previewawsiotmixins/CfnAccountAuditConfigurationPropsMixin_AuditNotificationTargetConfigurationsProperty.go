package previewawsiotmixins


// The configuration of the audit notification target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   auditNotificationTargetConfigurationsProperty := &AuditNotificationTargetConfigurationsProperty{
//   	Sns: &AuditNotificationTargetProperty{
//   		Enabled: jsii.Boolean(false),
//   		RoleArn: jsii.String("roleArn"),
//   		TargetArn: jsii.String("targetArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-accountauditconfiguration-auditnotificationtargetconfigurations.html
//
type CfnAccountAuditConfigurationPropsMixin_AuditNotificationTargetConfigurationsProperty struct {
	// The `Sns` notification target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-accountauditconfiguration-auditnotificationtargetconfigurations.html#cfn-iot-accountauditconfiguration-auditnotificationtargetconfigurations-sns
	//
	Sns interface{} `field:"optional" json:"sns" yaml:"sns"`
}

