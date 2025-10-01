package awssagemaker


// A reference to a UserProfile resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   userProfileReference := &UserProfileReference{
//   	DomainId: jsii.String("domainId"),
//   	UserProfileArn: jsii.String("userProfileArn"),
//   	UserProfileName: jsii.String("userProfileName"),
//   }
//
type UserProfileReference struct {
	// The DomainId of the UserProfile resource.
	DomainId *string `field:"required" json:"domainId" yaml:"domainId"`
	// The ARN of the UserProfile resource.
	UserProfileArn *string `field:"required" json:"userProfileArn" yaml:"userProfileArn"`
	// The UserProfileName of the UserProfile resource.
	UserProfileName *string `field:"required" json:"userProfileName" yaml:"userProfileName"`
}

