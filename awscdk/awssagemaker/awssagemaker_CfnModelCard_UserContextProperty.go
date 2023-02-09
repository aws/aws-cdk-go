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
type CfnModelCard_UserContextProperty struct {
	// `CfnModelCard.UserContextProperty.DomainId`.
	DomainId *string `field:"optional" json:"domainId" yaml:"domainId"`
	// `CfnModelCard.UserContextProperty.UserProfileArn`.
	UserProfileArn *string `field:"optional" json:"userProfileArn" yaml:"userProfileArn"`
	// `CfnModelCard.UserContextProperty.UserProfileName`.
	UserProfileName *string `field:"optional" json:"userProfileName" yaml:"userProfileName"`
}

