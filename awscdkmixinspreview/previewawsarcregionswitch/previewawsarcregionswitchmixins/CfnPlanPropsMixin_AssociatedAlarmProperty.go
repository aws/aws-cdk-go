package previewawsarcregionswitchmixins


// An Amazon CloudWatch alarm associated with a Region switch plan.
//
// These alarms can be used to trigger automatic execution of the plan.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   associatedAlarmProperty := &AssociatedAlarmProperty{
//   	AlarmType: jsii.String("alarmType"),
//   	CrossAccountRole: jsii.String("crossAccountRole"),
//   	ExternalId: jsii.String("externalId"),
//   	ResourceIdentifier: jsii.String("resourceIdentifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-associatedalarm.html
//
type CfnPlanPropsMixin_AssociatedAlarmProperty struct {
	// The alarm type for an associated alarm.
	//
	// An associated CloudWatch alarm can be an application health alarm or a trigger alarm.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-associatedalarm.html#cfn-arcregionswitch-plan-associatedalarm-alarmtype
	//
	AlarmType *string `field:"optional" json:"alarmType" yaml:"alarmType"`
	// The cross account role for the configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-associatedalarm.html#cfn-arcregionswitch-plan-associatedalarm-crossaccountrole
	//
	CrossAccountRole *string `field:"optional" json:"crossAccountRole" yaml:"crossAccountRole"`
	// The external ID (secret key) for the configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-associatedalarm.html#cfn-arcregionswitch-plan-associatedalarm-externalid
	//
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// The resource identifier for alarms that you associate with a plan.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-associatedalarm.html#cfn-arcregionswitch-plan-associatedalarm-resourceidentifier
	//
	ResourceIdentifier *string `field:"optional" json:"resourceIdentifier" yaml:"resourceIdentifier"`
}

