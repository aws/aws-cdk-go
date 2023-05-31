package awss3

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// The intelligent tiering configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var duration duration
//
//   intelligentTieringConfiguration := &IntelligentTieringConfiguration{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	ArchiveAccessTierTime: duration,
//   	DeepArchiveAccessTierTime: duration,
//   	Prefix: jsii.String("prefix"),
//   	Tags: []tag{
//   		&tag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// Experimental.
type IntelligentTieringConfiguration struct {
	// Configuration name.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// When enabled, Intelligent-Tiering will automatically move objects that haven’t been accessed for a minimum of 90 days to the Archive Access tier.
	// Experimental.
	ArchiveAccessTierTime awscdk.Duration `field:"optional" json:"archiveAccessTierTime" yaml:"archiveAccessTierTime"`
	// When enabled, Intelligent-Tiering will automatically move objects that haven’t been accessed for a minimum of 180 days to the Deep Archive Access tier.
	// Experimental.
	DeepArchiveAccessTierTime awscdk.Duration `field:"optional" json:"deepArchiveAccessTierTime" yaml:"deepArchiveAccessTierTime"`
	// Add a filter to limit the scope of this configuration to a single prefix.
	// Experimental.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// You can limit the scope of this rule to the key value pairs added below.
	// Experimental.
	Tags *[]*Tag `field:"optional" json:"tags" yaml:"tags"`
}

