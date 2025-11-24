package mixinsawsdatasync


// The AWS DataSync agents that can connect to your Network File System (NFS) file server.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   onPremConfigProperty := &OnPremConfigProperty{
//   	AgentArns: []*string{
//   		jsii.String("agentArns"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locationnfs-onpremconfig.html
//
type CfnLocationNFSPropsMixin_OnPremConfigProperty struct {
	// The Amazon Resource Names (ARNs) of the DataSync agents that can connect to your NFS file server.
	//
	// You can specify more than one agent. For more information, see [Using multiple DataSync agents](https://docs.aws.amazon.com/datasync/latest/userguide/do-i-need-datasync-agent.html#multiple-agents) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locationnfs-onpremconfig.html#cfn-datasync-locationnfs-onpremconfig-agentarns
	//
	AgentArns *[]*string `field:"optional" json:"agentArns" yaml:"agentArns"`
}

