package awsmediapackagev2alpha


// Details about the content that you want MediaPackage to pass through in the manifest to the playback device.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import mediapackagev2_alpha "github.com/aws/aws-cdk-go/awsmediapackagev2alpha"
//
//   dashProgramInformation := &DashProgramInformation{
//   	Copyright: jsii.String("copyright"),
//   	LanguageCode: jsii.String("languageCode"),
//   	MoreInformationUrl: jsii.String("moreInformationUrl"),
//   	Source: jsii.String("source"),
//   	Title: jsii.String("title"),
//   }
//
// Experimental.
type DashProgramInformation struct {
	// A copyright statement about the content.
	// Default: - No copyright information.
	//
	// Experimental.
	Copyright *string `field:"optional" json:"copyright" yaml:"copyright"`
	// The language code for this manifest.
	// Default: - No language code specified.
	//
	// Experimental.
	LanguageCode *string `field:"optional" json:"languageCode" yaml:"languageCode"`
	// An absolute URL that contains more information about this content.
	// Default: - No additional information URL.
	//
	// Experimental.
	MoreInformationUrl *string `field:"optional" json:"moreInformationUrl" yaml:"moreInformationUrl"`
	// Information about the content provider.
	// Default: - No source information.
	//
	// Experimental.
	Source *string `field:"optional" json:"source" yaml:"source"`
	// The title for the manifest.
	// Default: - No title specified.
	//
	// Experimental.
	Title *string `field:"optional" json:"title" yaml:"title"`
}

