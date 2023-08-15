package awss3

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Describes when noncurrent versions transition to a specified storage class.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var storageClass storageClass
//
//   noncurrentVersionTransition := &NoncurrentVersionTransition{
//   	StorageClass: storageClass,
//   	TransitionAfter: cdk.Duration_Minutes(jsii.Number(30)),
//
//   	// the properties below are optional
//   	NoncurrentVersionsToRetain: jsii.Number(123),
//   }
//
type NoncurrentVersionTransition struct {
	// The storage class to which you want the object to transition.
	StorageClass StorageClass `field:"required" json:"storageClass" yaml:"storageClass"`
	// Indicates the number of days after creation when objects are transitioned to the specified storage class.
	// Default: - No transition count.
	//
	TransitionAfter awscdk.Duration `field:"required" json:"transitionAfter" yaml:"transitionAfter"`
	// Indicates the number of noncurrent version objects to be retained.
	//
	// Can be up to 100 noncurrent versions retained.
	// Default: - No noncurrent version retained.
	//
	NoncurrentVersionsToRetain *float64 `field:"optional" json:"noncurrentVersionsToRetain" yaml:"noncurrentVersionsToRetain"`
}

