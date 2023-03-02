package awscodepipeline

import (
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awss3"
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
//   actionBindOptions := &actionBindOptions{
//   	bucket: bucket,
//   	role: role,
//   }
//
// Experimental.
type ActionBindOptions struct {
	// Experimental.
	Bucket awss3.IBucket `field:"required" json:"bucket" yaml:"bucket"`
	// Experimental.
	Role awsiam.IRole `field:"required" json:"role" yaml:"role"`
}

