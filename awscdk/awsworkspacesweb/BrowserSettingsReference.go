package awsworkspacesweb


// A reference to a BrowserSettings resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   browserSettingsReference := &BrowserSettingsReference{
//   	BrowserSettingsArn: jsii.String("browserSettingsArn"),
//   }
//
type BrowserSettingsReference struct {
	// The BrowserSettingsArn of the BrowserSettings resource.
	BrowserSettingsArn *string `field:"required" json:"browserSettingsArn" yaml:"browserSettingsArn"`
}

