package awsdevicefarm

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnProject`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnProjectProps := &CfnProjectProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	DefaultJobTimeoutMinutes: jsii.Number(123),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnProjectProps struct {
	// The project's name.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Sets the execution timeout value (in minutes) for a project.
	//
	// All test runs in this project use the specified execution timeout value unless overridden when scheduling a run.
	DefaultJobTimeoutMinutes *float64 `field:"optional" json:"defaultJobTimeoutMinutes" yaml:"defaultJobTimeoutMinutes"`
	// The tags to add to the resource.
	//
	// A tag is an array of key-value pairs. Tag keys can have a maximum character length of 128 characters. Tag values can have a maximum length of 256 characters.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

