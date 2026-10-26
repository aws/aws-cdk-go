package awsstoragegateway


// Default values for the file and directory permissions of an NFS share.
//
// Removing this property from a template leaves the live values unchanged rather than clearing them, because the service has no reset for it: a create that omits the object leaves it absent, but the update path merges member by member keyed on the same internal sentinel that reports absence, so no request can restore that state. Set the members you want rather than relying on removal.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   nFSFileShareDefaultsProperty := &NFSFileShareDefaultsProperty{
//   	DirectoryMode: jsii.String("directoryMode"),
//   	FileMode: jsii.String("fileMode"),
//   	GroupId: jsii.Number(123),
//   	OwnerId: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-storagegateway-nfsfileshare-nfsfilesharedefaults.html
//
type CfnNFSFileSharePropsMixin_NFSFileShareDefaultsProperty struct {
	// The Unix directory mode in the form nnnn.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-storagegateway-nfsfileshare-nfsfilesharedefaults.html#cfn-storagegateway-nfsfileshare-nfsfilesharedefaults-directorymode
	//
	DirectoryMode *string `field:"optional" json:"directoryMode" yaml:"directoryMode"`
	// The Unix file mode in the form nnnn.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-storagegateway-nfsfileshare-nfsfilesharedefaults.html#cfn-storagegateway-nfsfileshare-nfsfilesharedefaults-filemode
	//
	FileMode *string `field:"optional" json:"fileMode" yaml:"fileMode"`
	// The default group ID for files in the file share.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-storagegateway-nfsfileshare-nfsfilesharedefaults.html#cfn-storagegateway-nfsfileshare-nfsfilesharedefaults-groupid
	//
	GroupId *float64 `field:"optional" json:"groupId" yaml:"groupId"`
	// The default owner ID for files in the file share.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-storagegateway-nfsfileshare-nfsfilesharedefaults.html#cfn-storagegateway-nfsfileshare-nfsfilesharedefaults-ownerid
	//
	OwnerId *float64 `field:"optional" json:"ownerId" yaml:"ownerId"`
}

