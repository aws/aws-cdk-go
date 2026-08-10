package awsusernotifications


// Properties for defining a `CfnManagedNotificationConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnManagedNotificationConfigurationProps := &CfnManagedNotificationConfigurationProps{
//   	Category: jsii.String("category"),
//   	SubCategory: jsii.String("subCategory"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-usernotifications-managednotificationconfiguration.html
//
type CfnManagedNotificationConfigurationProps struct {
	// The category of the ManagedNotificationConfiguration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-usernotifications-managednotificationconfiguration.html#cfn-usernotifications-managednotificationconfiguration-category
	//
	Category *string `field:"optional" json:"category" yaml:"category"`
	// The subCategory of the ManagedNotificationConfiguration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-usernotifications-managednotificationconfiguration.html#cfn-usernotifications-managednotificationconfiguration-subcategory
	//
	SubCategory *string `field:"optional" json:"subCategory" yaml:"subCategory"`
}

