package awsgroundstation

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnMissionProfile`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMissionProfileProps := &CfnMissionProfileProps{
//   	DataflowEdges: []interface{}{
//   		&DataflowEdgeProperty{
//   			Destination: jsii.String("destination"),
//   			Source: jsii.String("source"),
//   		},
//   	},
//   	MinimumViableContactDurationSeconds: jsii.Number(123),
//   	Name: jsii.String("name"),
//   	TrackingConfigArn: jsii.String("trackingConfigArn"),
//
//   	// the properties below are optional
//   	ContactPostPassDurationSeconds: jsii.Number(123),
//   	ContactPrePassDurationSeconds: jsii.Number(123),
//   	StreamsKmsKey: &StreamsKmsKeyProperty{
//   		KmsAliasArn: jsii.String("kmsAliasArn"),
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
//   	},
//   	StreamsKmsRole: jsii.String("streamsKmsRole"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnMissionProfileProps struct {
	// A list containing lists of config ARNs.
	//
	// Each list of config ARNs is an edge, with a "from" config and a "to" config.
	DataflowEdges interface{} `field:"required" json:"dataflowEdges" yaml:"dataflowEdges"`
	// Minimum length of a contact in seconds that Ground Station will return when listing contacts.
	//
	// Ground Station will not return contacts shorter than this duration.
	MinimumViableContactDurationSeconds *float64 `field:"required" json:"minimumViableContactDurationSeconds" yaml:"minimumViableContactDurationSeconds"`
	// The name of the mission profile.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ARN of a tracking config objects that defines how to track the satellite through the sky during a contact.
	TrackingConfigArn *string `field:"required" json:"trackingConfigArn" yaml:"trackingConfigArn"`
	// Amount of time in seconds after a contact ends that you’d like to receive a CloudWatch Event indicating the pass has finished.
	//
	// For more information on CloudWatch Events, see the [What Is CloudWatch Events?](https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/WhatIsCloudWatchEvents.html)
	ContactPostPassDurationSeconds *float64 `field:"optional" json:"contactPostPassDurationSeconds" yaml:"contactPostPassDurationSeconds"`
	// Amount of time in seconds prior to contact start that you'd like to receive a CloudWatch Event indicating an upcoming pass.
	//
	// For more information on CloudWatch Events, see the [What Is CloudWatch Events?](https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/WhatIsCloudWatchEvents.html)
	ContactPrePassDurationSeconds *float64 `field:"optional" json:"contactPrePassDurationSeconds" yaml:"contactPrePassDurationSeconds"`
	// `AWS::GroundStation::MissionProfile.StreamsKmsKey`.
	StreamsKmsKey interface{} `field:"optional" json:"streamsKmsKey" yaml:"streamsKmsKey"`
	// `AWS::GroundStation::MissionProfile.StreamsKmsRole`.
	StreamsKmsRole *string `field:"optional" json:"streamsKmsRole" yaml:"streamsKmsRole"`
	// Tags assigned to the mission profile.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

