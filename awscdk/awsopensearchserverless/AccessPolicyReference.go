package awsopensearchserverless


// A reference to a AccessPolicy resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessPolicyReference := &AccessPolicyReference{
//   	AccessPolicyName: jsii.String("accessPolicyName"),
//   	Type: jsii.String("type"),
//   }
//
type AccessPolicyReference struct {
	// The Name of the AccessPolicy resource.
	AccessPolicyName *string `field:"required" json:"accessPolicyName" yaml:"accessPolicyName"`
	// The Type of the AccessPolicy resource.
	Type *string `field:"required" json:"type" yaml:"type"`
}

