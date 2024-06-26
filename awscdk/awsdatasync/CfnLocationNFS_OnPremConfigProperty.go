package awsdatasync


// The AWS DataSync agents that are connecting to a Network File System (NFS) location.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   onPremConfigProperty := &OnPremConfigProperty{
//   	AgentArns: []*string{
//   		jsii.String("agentArns"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locationnfs-onpremconfig.html
//
type CfnLocationNFS_OnPremConfigProperty struct {
	// The Amazon Resource Names (ARNs) of the agents connecting to a transfer location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locationnfs-onpremconfig.html#cfn-datasync-locationnfs-onpremconfig-agentarns
	//
	AgentArns *[]*string `field:"required" json:"agentArns" yaml:"agentArns"`
}

