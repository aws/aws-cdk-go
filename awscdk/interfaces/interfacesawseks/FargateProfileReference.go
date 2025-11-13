package interfacesawseks


// A reference to a FargateProfile resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fargateProfileReference := &FargateProfileReference{
//   	ClusterName: jsii.String("clusterName"),
//   	FargateProfileArn: jsii.String("fargateProfileArn"),
//   	FargateProfileName: jsii.String("fargateProfileName"),
//   }
//
type FargateProfileReference struct {
	// The ClusterName of the FargateProfile resource.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// The ARN of the FargateProfile resource.
	FargateProfileArn *string `field:"required" json:"fargateProfileArn" yaml:"fargateProfileArn"`
	// The FargateProfileName of the FargateProfile resource.
	FargateProfileName *string `field:"required" json:"fargateProfileName" yaml:"fargateProfileName"`
}

