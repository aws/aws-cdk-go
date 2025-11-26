package previewawsapsmixins


// Configuration settings for a scraper component.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   componentConfigProperty := &ComponentConfigProperty{
//   	Options: map[string]*string{
//   		"optionsKey": jsii.String("options"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-scraper-componentconfig.html
//
type CfnScraperPropsMixin_ComponentConfigProperty struct {
	// Configuration options for the scraper component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-scraper-componentconfig.html#cfn-aps-scraper-componentconfig-options
	//
	Options interface{} `field:"optional" json:"options" yaml:"options"`
}

