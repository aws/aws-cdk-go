package awsglue


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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-icebergtarget.html#cfn-glue-crawler-icebergtarget-connectionname
	//
	ConnectionName *string `field:"optional" json:"connectionName" yaml:"connectionName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-icebergtarget.html#cfn-glue-crawler-icebergtarget-exclusions
	//
	Exclusions *[]*string `field:"optional" json:"exclusions" yaml:"exclusions"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-icebergtarget.html#cfn-glue-crawler-icebergtarget-maximumtraversaldepth
	//
	MaximumTraversalDepth *float64 `field:"optional" json:"maximumTraversalDepth" yaml:"maximumTraversalDepth"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-icebergtarget.html#cfn-glue-crawler-icebergtarget-paths
	//
	Paths *[]*string `field:"optional" json:"paths" yaml:"paths"`
}

