package awscodepipeline

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// An interface representing resources generated in order to support the cross-region capabilities of CodePipeline.
//
// You get instances of this interface from the `Pipeline#crossRegionSupport` property.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket bucket
//   var stack stack
//
//   crossRegionSupport := &CrossRegionSupport{
//   	ReplicationBucket: bucket,
//   	Stack: stack,
//   }
//
type CrossRegionSupport struct {
	// The replication Bucket used by CodePipeline to operate in this region.
	//
	// Belongs to `stack`.
	ReplicationBucket awss3.IBucket `field:"required" json:"replicationBucket" yaml:"replicationBucket"`
	// The Stack that has been created to house the replication Bucket required for this  region.
	Stack awscdk.Stack `field:"required" json:"stack" yaml:"stack"`
}

