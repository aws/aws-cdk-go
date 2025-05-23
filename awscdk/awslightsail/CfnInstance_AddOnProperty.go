package awslightsail


// `AddOn` is a property of the [AWS::Lightsail::Instance](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-instance.html) resource. It describes the add-ons for an instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   addOnProperty := &AddOnProperty{
//   	AddOnType: jsii.String("addOnType"),
//
//   	// the properties below are optional
//   	AutoSnapshotAddOnRequest: &AutoSnapshotAddOnProperty{
//   		SnapshotTimeOfDay: jsii.String("snapshotTimeOfDay"),
//   	},
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-instance-addon.html
//
type CfnInstance_AddOnProperty struct {
	// The add-on type (for example, `AutoSnapshot` ).
	//
	// > `AutoSnapshot` is the only add-on that can be enabled for an instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-instance-addon.html#cfn-lightsail-instance-addon-addontype
	//
	AddOnType *string `field:"required" json:"addOnType" yaml:"addOnType"`
	// The parameters for the automatic snapshot add-on, such as the daily time when an automatic snapshot will be created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-instance-addon.html#cfn-lightsail-instance-addon-autosnapshotaddonrequest
	//
	AutoSnapshotAddOnRequest interface{} `field:"optional" json:"autoSnapshotAddOnRequest" yaml:"autoSnapshotAddOnRequest"`
	// The status of the add-on.
	//
	// Valid Values: `Enabled` | `Disabled`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-instance-addon.html#cfn-lightsail-instance-addon-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

