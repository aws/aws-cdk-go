package awss3


// Properties for defining a `CfnMultiRegionAccessPoint`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMultiRegionAccessPointProps := &cfnMultiRegionAccessPointProps{
//   	regions: []interface{}{
//   		&regionProperty{
//   			bucket: jsii.String("bucket"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	name: jsii.String("name"),
//   	publicAccessBlockConfiguration: &publicAccessBlockConfigurationProperty{
//   		blockPublicAcls: jsii.Boolean(false),
//   		blockPublicPolicy: jsii.Boolean(false),
//   		ignorePublicAcls: jsii.Boolean(false),
//   		restrictPublicBuckets: jsii.Boolean(false),
//   	},
//   }
//
type CfnMultiRegionAccessPointProps struct {
	// A collection of the Regions and buckets associated with the Multi-Region Access Point.
	Regions interface{} `field:"required" json:"regions" yaml:"regions"`
	// The name of the Multi-Region Access Point.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The PublicAccessBlock configuration that you want to apply to this Multi-Region Access Point.
	//
	// You can enable the configuration options in any combination. For more information about when Amazon S3 considers an object public, see [The Meaning of "Public"](https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status) in the *Amazon S3 User Guide* .
	PublicAccessBlockConfiguration interface{} `field:"optional" json:"publicAccessBlockConfiguration" yaml:"publicAccessBlockConfiguration"`
}

