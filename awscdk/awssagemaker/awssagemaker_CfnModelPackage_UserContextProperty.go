package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   userContextProperty := &userContextProperty{
//   	domainId: jsii.String("domainId"),
//   	userProfileArn: jsii.String("userProfileArn"),
//   	userProfileName: jsii.String("userProfileName"),
//   }
//
type CfnModelPackage_UserContextProperty struct {
	// `CfnModelPackage.UserContextProperty.DomainId`.
	DomainId *string `field:"optional" json:"domainId" yaml:"domainId"`
	// `CfnModelPackage.UserContextProperty.UserProfileArn`.
	UserProfileArn *string `field:"optional" json:"userProfileArn" yaml:"userProfileArn"`
	// `CfnModelPackage.UserContextProperty.UserProfileName`.
	UserProfileName *string `field:"optional" json:"userProfileName" yaml:"userProfileName"`
}

