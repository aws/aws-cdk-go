package awsglue


// Specifies Apache Iceberg data store targets.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   icebergTargetProperty := &IcebergTargetProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-icebergtarget.html
//
type CfnCrawler_IcebergTargetProperty struct {
	// The name of the connection to use to connect to the Iceberg target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-icebergtarget.html#cfn-glue-crawler-icebergtarget-connectionname
	//
	ConnectionName *string `field:"optional" json:"connectionName" yaml:"connectionName"`
	// A list of global patterns used to exclude from the crawl.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-icebergtarget.html#cfn-glue-crawler-icebergtarget-exclusions
	//
	Exclusions *[]*string `field:"optional" json:"exclusions" yaml:"exclusions"`
	// The maximum depth of Amazon S3 paths that the crawler can traverse to discover the Iceberg metadata folder in your Amazon S3 path.
	//
	// Used to limit the crawler run time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-icebergtarget.html#cfn-glue-crawler-icebergtarget-maximumtraversaldepth
	//
	MaximumTraversalDepth *float64 `field:"optional" json:"maximumTraversalDepth" yaml:"maximumTraversalDepth"`
	// One or more Amazon S3 paths that contains Iceberg metadata folders as s3://bucket/prefix .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-icebergtarget.html#cfn-glue-crawler-icebergtarget-paths
	//
	Paths *[]*string `field:"optional" json:"paths" yaml:"paths"`
}

