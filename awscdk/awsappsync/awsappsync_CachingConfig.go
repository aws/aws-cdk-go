package awsappsync

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// CachingConfig for AppSync resolvers.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var duration duration
//
//   cachingConfig := &cachingConfig{
//   	ttl: duration,
//
//   	// the properties below are optional
//   	cachingKeys: []*string{
//   		jsii.String("cachingKeys"),
//   	},
//   }
//
// Experimental.
type CachingConfig struct {
	// The TTL in seconds for a resolver that has caching enabled.
	//
	// Valid values are between 1 and 3600 seconds.
	// Experimental.
	Ttl awscdk.Duration `field:"required" json:"ttl" yaml:"ttl"`
	// The caching keys for a resolver that has caching enabled.
	//
	// Valid values are entries from the $context.arguments, $context.source, and $context.identity maps.
	// Experimental.
	CachingKeys *[]*string `field:"optional" json:"cachingKeys" yaml:"cachingKeys"`
}

