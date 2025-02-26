package awseks


// Represents the properties required to create an Amazon EKS access entry.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var accessPolicy accessPolicy
//   var cluster cluster
//
//   accessEntryProps := &AccessEntryProps{
//   	AccessPolicies: []iAccessPolicy{
//   		accessPolicy,
//   	},
//   	Cluster: cluster,
//   	Principal: jsii.String("principal"),
//
//   	// the properties below are optional
//   	AccessEntryName: jsii.String("accessEntryName"),
//   	AccessEntryType: awscdk.Aws_eks.AccessEntryType_STANDARD,
//   }
//
type AccessEntryProps struct {
	// The access policies that define the permissions and scope for the access entry.
	AccessPolicies *[]IAccessPolicy `field:"required" json:"accessPolicies" yaml:"accessPolicies"`
	// The Amazon EKS cluster to which the access entry applies.
	Cluster ICluster `field:"required" json:"cluster" yaml:"cluster"`
	// The Amazon Resource Name (ARN) of the principal (user or role) to associate the access entry with.
	Principal *string `field:"required" json:"principal" yaml:"principal"`
	// The name of the AccessEntry.
	// Default: - No access entry name is provided.
	//
	AccessEntryName *string `field:"optional" json:"accessEntryName" yaml:"accessEntryName"`
	// The type of the AccessEntry.
	// Default: STANDARD.
	//
	AccessEntryType AccessEntryType `field:"optional" json:"accessEntryType" yaml:"accessEntryType"`
}

