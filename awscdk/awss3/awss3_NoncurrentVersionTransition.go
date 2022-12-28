package awss3

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Describes when noncurrent versions transition to a specified storage class.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var duration duration
//   var storageClass storageClass
//
//   noncurrentVersionTransition := &noncurrentVersionTransition{
//   	storageClass: storageClass,
//   	transitionAfter: duration,
//
//   	// the properties below are optional
//   	noncurrentVersionsToRetain: jsii.Number(123),
//   }
//
// Experimental.
type NoncurrentVersionTransition struct {
	// The storage class to which you want the object to transition.
	// Experimental.
	StorageClass StorageClass `field:"required" json:"storageClass" yaml:"storageClass"`
	// Indicates the number of days after creation when objects are transitioned to the specified storage class.
	// Experimental.
	TransitionAfter awscdk.Duration `field:"required" json:"transitionAfter" yaml:"transitionAfter"`
	// Indicates the number of noncurrent version objects to be retained.
	//
	// Can be up to 100 noncurrent versions retained.
	// Experimental.
	NoncurrentVersionsToRetain *float64 `field:"optional" json:"noncurrentVersionsToRetain" yaml:"noncurrentVersionsToRetain"`
}

