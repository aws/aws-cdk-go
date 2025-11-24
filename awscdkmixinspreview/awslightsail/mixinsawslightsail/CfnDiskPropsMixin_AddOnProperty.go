package mixinsawslightsail


// `AddOn` is a property of the [AWS::Lightsail::Disk](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-disk.html) resource. It describes the add-ons for a disk.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   addOnProperty := &AddOnProperty{
//   	AddOnType: jsii.String("addOnType"),
//   	AutoSnapshotAddOnRequest: &AutoSnapshotAddOnProperty{
//   		SnapshotTimeOfDay: jsii.String("snapshotTimeOfDay"),
//   	},
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-disk-addon.html
//
type CfnDiskPropsMixin_AddOnProperty struct {
	// The add-on type (for example, `AutoSnapshot` ).
	//
	// > `AutoSnapshot` is the only add-on that can be enabled for a disk.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-disk-addon.html#cfn-lightsail-disk-addon-addontype
	//
	AddOnType *string `field:"optional" json:"addOnType" yaml:"addOnType"`
	// The parameters for the automatic snapshot add-on, such as the daily time when an automatic snapshot will be created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-disk-addon.html#cfn-lightsail-disk-addon-autosnapshotaddonrequest
	//
	AutoSnapshotAddOnRequest interface{} `field:"optional" json:"autoSnapshotAddOnRequest" yaml:"autoSnapshotAddOnRequest"`
	// The status of the add-on.
	//
	// Valid Values: `Enabled` | `Disabled`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-disk-addon.html#cfn-lightsail-disk-addon-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

