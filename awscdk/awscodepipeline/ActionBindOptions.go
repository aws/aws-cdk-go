package awscodepipeline

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket bucket
//   var role role
//
//   actionBindOptions := &ActionBindOptions{
//   	Bucket: bucket,
//   	Role: role,
//   }
//
type ActionBindOptions struct {
	Bucket awss3.IBucket `field:"required" json:"bucket" yaml:"bucket"`
	Role awsiam.IRole `field:"required" json:"role" yaml:"role"`
}

