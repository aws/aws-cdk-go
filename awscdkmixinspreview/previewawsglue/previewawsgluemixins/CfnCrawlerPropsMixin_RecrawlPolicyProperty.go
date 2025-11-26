package previewawsgluemixins


// When crawling an Amazon S3 data source after the first crawl is complete, specifies whether to crawl the entire dataset again or to crawl only folders that were added since the last crawler run.
//
// For more information, see [Incremental Crawls in AWS Glue](https://docs.aws.amazon.com/glue/latest/dg/incremental-crawls.html) in the developer guide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   recrawlPolicyProperty := &RecrawlPolicyProperty{
//   	RecrawlBehavior: jsii.String("recrawlBehavior"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-recrawlpolicy.html
//
type CfnCrawlerPropsMixin_RecrawlPolicyProperty struct {
	// Specifies whether to crawl the entire dataset again or to crawl only folders that were added since the last crawler run.
	//
	// A value of `CRAWL_EVERYTHING` specifies crawling the entire dataset again.
	//
	// A value of `CRAWL_NEW_FOLDERS_ONLY` specifies crawling only folders that were added since the last crawler run.
	//
	// A value of `CRAWL_EVENT_MODE` specifies crawling only the changes identified by Amazon S3 events.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-recrawlpolicy.html#cfn-glue-crawler-recrawlpolicy-recrawlbehavior
	//
	RecrawlBehavior *string `field:"optional" json:"recrawlBehavior" yaml:"recrawlBehavior"`
}

