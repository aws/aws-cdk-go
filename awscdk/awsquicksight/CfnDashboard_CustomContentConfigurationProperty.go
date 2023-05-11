package awsquicksight


// The configuration of a `CustomContentVisual` .
//
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
	// The content type of the custom content visual.
	//
	// You can use this to have the visual render as an image.
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// The input URL that links to the custom content that you want in the custom visual.
	ContentUrl *string `field:"optional" json:"contentUrl" yaml:"contentUrl"`
	// The sizing options for the size of the custom content visual.
	//
	// This structure is required when the `ContentType` of the visual is `'IMAGE'` .
	ImageScaling *string `field:"optional" json:"imageScaling" yaml:"imageScaling"`
}

