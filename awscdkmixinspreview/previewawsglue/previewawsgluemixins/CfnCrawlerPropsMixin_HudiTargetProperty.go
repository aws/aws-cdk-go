package previewawsgluemixins


// Specifies an Apache Hudi data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   hudiTargetProperty := &HudiTargetProperty{
//   	ConnectionName: jsii.String("connectionName"),
//   	Exclusions: []*string{
//   		jsii.String("exclusions"),
//   	},
//   	MaximumTraversalDepth: jsii.Number(123),
//   	Paths: []*string{
//   		jsii.String("paths"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-huditarget.html
//
type CfnCrawlerPropsMixin_HudiTargetProperty struct {
	// The name of the connection to use to connect to the Hudi target.
	//
	// If your Hudi files are stored in buckets that require VPC authorization, you can set their connection properties here.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-huditarget.html#cfn-glue-crawler-huditarget-connectionname
	//
	ConnectionName *string `field:"optional" json:"connectionName" yaml:"connectionName"`
	// A list of glob patterns used to exclude from the crawl.
	//
	// For more information, see [Catalog Tables with a Crawler](https://docs.aws.amazon.com/glue/latest/dg/add-crawler.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-huditarget.html#cfn-glue-crawler-huditarget-exclusions
	//
	Exclusions *[]*string `field:"optional" json:"exclusions" yaml:"exclusions"`
	// The maximum depth of Amazon S3 paths that the crawler can traverse to discover the Hudi metadata folder in your Amazon S3 path.
	//
	// Used to limit the crawler run time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-huditarget.html#cfn-glue-crawler-huditarget-maximumtraversaldepth
	//
	MaximumTraversalDepth *float64 `field:"optional" json:"maximumTraversalDepth" yaml:"maximumTraversalDepth"`
	// An array of Amazon S3 location strings for Hudi, each indicating the root folder with which the metadata files for a Hudi table resides.
	//
	// The Hudi folder may be located in a child folder of the root folder.
	//
	// The crawler will scan all folders underneath a path for a Hudi folder.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-huditarget.html#cfn-glue-crawler-huditarget-paths
	//
	Paths *[]*string `field:"optional" json:"paths" yaml:"paths"`
}

