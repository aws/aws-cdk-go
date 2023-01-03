package awsdatasync


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sMBProperty := &sMBProperty{
//   	mountOptions: &smbMountOptionsProperty{
//   		version: jsii.String("version"),
//   	},
//   	password: jsii.String("password"),
//   	user: jsii.String("user"),
//
//   	// the properties below are optional
//   	domain: jsii.String("domain"),
//   }
//
type CfnLocationFSxONTAP_SMBProperty struct {
	// `CfnLocationFSxONTAP.SMBProperty.MountOptions`.
	MountOptions interface{} `field:"required" json:"mountOptions" yaml:"mountOptions"`
	// `CfnLocationFSxONTAP.SMBProperty.Password`.
	Password *string `field:"required" json:"password" yaml:"password"`
	// `CfnLocationFSxONTAP.SMBProperty.User`.
	User *string `field:"required" json:"user" yaml:"user"`
	// `CfnLocationFSxONTAP.SMBProperty.Domain`.
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
}

