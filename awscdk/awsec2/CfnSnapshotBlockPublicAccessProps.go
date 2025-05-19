package awsec2


// Properties for defining a `CfnSnapshotBlockPublicAccess`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSnapshotBlockPublicAccessProps := &CfnSnapshotBlockPublicAccessProps{
//   	State: jsii.String("state"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-snapshotblockpublicaccess.html
//
type CfnSnapshotBlockPublicAccessProps struct {
	// The mode in which to enable block public access for snapshots for the Region.
	//
	// Specify one of the following values:
	//
	// - `block-all-sharing` - Prevents all public sharing of snapshots in the Region. Users in the account will no longer be able to request new public sharing. Additionally, snapshots that are already publicly shared are treated as private and they are no longer publicly available.
	//
	// > If you enable block public access for snapshots in `block-all-sharing` mode, it does not change the permissions for snapshots that are already publicly shared. Instead, it prevents these snapshots from be publicly visible and publicly accessible. Therefore, the attributes for these snapshots still indicate that they are publicly shared, even though they are not publicly available.
	// - `block-new-sharing` - Prevents only new public sharing of snapshots in the Region. Users in the account will no longer be able to request new public sharing. However, snapshots that are already publicly shared, remain publicly available.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-snapshotblockpublicaccess.html#cfn-ec2-snapshotblockpublicaccess-state
	//
	State *string `field:"required" json:"state" yaml:"state"`
}

