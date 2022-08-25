package awsdatasync


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   nFSProperty := &nFSProperty{
//   	mountOptions: &nfsMountOptionsProperty{
//   		version: jsii.String("version"),
//   	},
//   }
//
type CfnLocationFSxONTAP_NFSProperty struct {
	// `CfnLocationFSxONTAP.NFSProperty.MountOptions`.
	MountOptions interface{} `field:"required" json:"mountOptions" yaml:"mountOptions"`
}

