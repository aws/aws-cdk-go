package interfacesawsglue


// A reference to a Crawler resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   crawlerReference := &CrawlerReference{
//   	CrawlerName: jsii.String("crawlerName"),
//   }
//
type CrawlerReference struct {
	// The Name of the Crawler resource.
	CrawlerName *string `field:"required" json:"crawlerName" yaml:"crawlerName"`
}

