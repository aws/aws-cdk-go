package awscodepipeline

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awss3"
)

// An interface representing resources generated in order to support the cross-region capabilities of CodePipeline.
//
// You get instances of this interface from the {@link Pipeline#crossRegionSupport} property.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
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
// Experimental.
type CrossRegionSupport struct {
	// The replication Bucket used by CodePipeline to operate in this region.
	//
	// Belongs to {@link stack}.
	// Experimental.
	ReplicationBucket awss3.IBucket `field:"required" json:"replicationBucket" yaml:"replicationBucket"`
	// The Stack that has been created to house the replication Bucket required for this  region.
	// Experimental.
	Stack awscdk.Stack `field:"required" json:"stack" yaml:"stack"`
}

