package awss3

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// The intelligent tiering configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   intelligentTieringConfiguration := &intelligentTieringConfiguration{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	archiveAccessTierTime: cdk.duration.minutes(jsii.Number(30)),
//   	deepArchiveAccessTierTime: cdk.*duration.minutes(jsii.Number(30)),
//   	prefix: jsii.String("prefix"),
//   	tags: []tag{
//   		&tag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type IntelligentTieringConfiguration struct {
	// Configuration name.
	Name *string `field:"required" json:"name" yaml:"name"`
	// When enabled, Intelligent-Tiering will automatically move objects that haven’t been accessed for a minimum of 90 days to the Archive Access tier.
	ArchiveAccessTierTime awscdk.Duration `field:"optional" json:"archiveAccessTierTime" yaml:"archiveAccessTierTime"`
	// When enabled, Intelligent-Tiering will automatically move objects that haven’t been accessed for a minimum of 180 days to the Deep Archive Access tier.
	DeepArchiveAccessTierTime awscdk.Duration `field:"optional" json:"deepArchiveAccessTierTime" yaml:"deepArchiveAccessTierTime"`
	// Add a filter to limit the scope of this configuration to a single prefix.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// You can limit the scope of this rule to the key value pairs added below.
	Tags *[]*Tag `field:"optional" json:"tags" yaml:"tags"`
}

