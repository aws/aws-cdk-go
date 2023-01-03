package awsservicecatalog

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Result of binding `Template` into a `Product`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket bucket
//
//   cloudFormationTemplateConfig := &cloudFormationTemplateConfig{
//   	httpUrl: jsii.String("httpUrl"),
//
//   	// the properties below are optional
//   	assetBucket: bucket,
//   }
//
type CloudFormationTemplateConfig struct {
	// The http url of the template in S3.
	HttpUrl *string `field:"required" json:"httpUrl" yaml:"httpUrl"`
	// The S3 bucket containing product stack assets.
	AssetBucket awss3.IBucket `field:"optional" json:"assetBucket" yaml:"assetBucket"`
}

