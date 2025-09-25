package awseks


// A reference to a AccessEntry resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessEntryReference := &AccessEntryReference{
//   	AccessEntryArn: jsii.String("accessEntryArn"),
//   	ClusterName: jsii.String("clusterName"),
//   	PrincipalArn: jsii.String("principalArn"),
//   }
//
type AccessEntryReference struct {
	// The ARN of the AccessEntry resource.
	AccessEntryArn *string `field:"required" json:"accessEntryArn" yaml:"accessEntryArn"`
	// The ClusterName of the AccessEntry resource.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// The PrincipalArn of the AccessEntry resource.
	PrincipalArn *string `field:"required" json:"principalArn" yaml:"principalArn"`
}

