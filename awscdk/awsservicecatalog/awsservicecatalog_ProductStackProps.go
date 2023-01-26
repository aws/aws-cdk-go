package awsservicecatalog

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Product stack props.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket bucket
//
//   productStackProps := &productStackProps{
//   	assetBucket: bucket,
//   }
//
type ProductStackProps struct {
	// A Bucket can be passed to store assets, enabling ProductStack Asset support.
	AssetBucket awss3.IBucket `field:"optional" json:"assetBucket" yaml:"assetBucket"`
}

