package awsappsync

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// CachingConfig for AppSync resolvers.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cachingConfig := &CachingConfig{
//   	Ttl: cdk.Duration_Minutes(jsii.Number(30)),
//
//   	// the properties below are optional
//   	CachingKeys: []*string{
//   		jsii.String("cachingKeys"),
//   	},
//   }
//
type CachingConfig struct {
	// The TTL in seconds for a resolver that has caching enabled.
	//
	// Valid values are between 1 and 3600 seconds.
	Ttl awscdk.Duration `field:"required" json:"ttl" yaml:"ttl"`
	// The caching keys for a resolver that has caching enabled.
	//
	// Valid values are entries from the $context.arguments, $context.source, and $context.identity maps.
	// Default: - No caching keys.
	//
	CachingKeys *[]*string `field:"optional" json:"cachingKeys" yaml:"cachingKeys"`
}

