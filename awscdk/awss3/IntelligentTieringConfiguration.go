package awss3

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// The intelligent tiering configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   intelligentTieringConfiguration := &IntelligentTieringConfiguration{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	ArchiveAccessTierTime: cdk.Duration_Minutes(jsii.Number(30)),
//   	DeepArchiveAccessTierTime: cdk.Duration_*Minutes(jsii.Number(30)),
//   	Prefix: jsii.String("prefix"),
//   	Tags: []tag{
//   		&tag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type IntelligentTieringConfiguration struct {
	// Configuration name.
	Name *string `field:"required" json:"name" yaml:"name"`
	// When enabled, Intelligent-Tiering will automatically move objects that haven’t been accessed for a minimum of 90 days to the Archive Access tier.
	// Default: Objects will not move to Glacier.
	//
	ArchiveAccessTierTime awscdk.Duration `field:"optional" json:"archiveAccessTierTime" yaml:"archiveAccessTierTime"`
	// When enabled, Intelligent-Tiering will automatically move objects that haven’t been accessed for a minimum of 180 days to the Deep Archive Access tier.
	// Default: Objects will not move to Glacier Deep Access.
	//
	DeepArchiveAccessTierTime awscdk.Duration `field:"optional" json:"deepArchiveAccessTierTime" yaml:"deepArchiveAccessTierTime"`
	// Add a filter to limit the scope of this configuration to a single prefix.
	// Default: this configuration will apply to **all** objects in the bucket.
	//
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// You can limit the scope of this rule to the key value pairs added below.
	// Default: No filtering will be performed on tags.
	//
	Tags *[]*Tag `field:"optional" json:"tags" yaml:"tags"`
}

