package awsdatasync

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnLocationFSxONTAP`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLocationFSxONTAPProps := &cfnLocationFSxONTAPProps{
//   	protocol: &protocolProperty{
//   		nfs: &nFSProperty{
//   			mountOptions: &nfsMountOptionsProperty{
//   				version: jsii.String("version"),
//   			},
//   		},
//   		smb: &sMBProperty{
//   			mountOptions: &smbMountOptionsProperty{
//   				version: jsii.String("version"),
//   			},
//   			password: jsii.String("password"),
//   			user: jsii.String("user"),
//
//   			// the properties below are optional
//   			domain: jsii.String("domain"),
//   		},
//   	},
//   	securityGroupArns: []*string{
//   		jsii.String("securityGroupArns"),
//   	},
//   	storageVirtualMachineArn: jsii.String("storageVirtualMachineArn"),
//
//   	// the properties below are optional
//   	subdirectory: jsii.String("subdirectory"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnLocationFSxONTAPProps struct {
	// `AWS::DataSync::LocationFSxONTAP.Protocol`.
	Protocol interface{} `field:"required" json:"protocol" yaml:"protocol"`
	// `AWS::DataSync::LocationFSxONTAP.SecurityGroupArns`.
	SecurityGroupArns *[]*string `field:"required" json:"securityGroupArns" yaml:"securityGroupArns"`
	// `AWS::DataSync::LocationFSxONTAP.StorageVirtualMachineArn`.
	StorageVirtualMachineArn *string `field:"required" json:"storageVirtualMachineArn" yaml:"storageVirtualMachineArn"`
	// `AWS::DataSync::LocationFSxONTAP.Subdirectory`.
	Subdirectory *string `field:"optional" json:"subdirectory" yaml:"subdirectory"`
	// `AWS::DataSync::LocationFSxONTAP.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

