package awsevidently

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnSegment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSegmentProps := &cfnSegmentProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	pattern: jsii.String("pattern"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnSegmentProps struct {
	// `AWS::Evidently::Segment.Name`.
	Name *string `field:"required" json:"name" yaml:"name"`
	// `AWS::Evidently::Segment.Description`.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// `AWS::Evidently::Segment.Pattern`.
	Pattern *string `field:"optional" json:"pattern" yaml:"pattern"`
	// `AWS::Evidently::Segment.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

