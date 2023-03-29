package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customContentConfigurationProperty := &CustomContentConfigurationProperty{
//   	ContentType: jsii.String("contentType"),
//   	ContentUrl: jsii.String("contentUrl"),
//   	ImageScaling: jsii.String("imageScaling"),
//   }
//
type CfnDashboard_CustomContentConfigurationProperty struct {
	// `CfnDashboard.CustomContentConfigurationProperty.ContentType`.
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// `CfnDashboard.CustomContentConfigurationProperty.ContentUrl`.
	ContentUrl *string `field:"optional" json:"contentUrl" yaml:"contentUrl"`
	// `CfnDashboard.CustomContentConfigurationProperty.ImageScaling`.
	ImageScaling *string `field:"optional" json:"imageScaling" yaml:"imageScaling"`
}

