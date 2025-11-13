package interfacesawsaps


// A reference to a Scraper resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scraperReference := &ScraperReference{
//   	ScraperArn: jsii.String("scraperArn"),
//   }
//
type ScraperReference struct {
	// The Arn of the Scraper resource.
	ScraperArn *string `field:"required" json:"scraperArn" yaml:"scraperArn"`
}

