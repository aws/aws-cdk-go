package previewawsgroundstationmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnMissionProfilePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnMissionProfileMixinProps := &CfnMissionProfileMixinProps{
//   	ContactPostPassDurationSeconds: jsii.Number(123),
//   	ContactPrePassDurationSeconds: jsii.Number(123),
//   	DataflowEdges: []interface{}{
//   		&DataflowEdgeProperty{
//   			Destination: jsii.String("destination"),
//   			Source: jsii.String("source"),
//   		},
//   	},
//   	MinimumViableContactDurationSeconds: jsii.Number(123),
//   	Name: jsii.String("name"),
//   	StreamsKmsKey: &StreamsKmsKeyProperty{
//   		KmsAliasArn: jsii.String("kmsAliasArn"),
//   		KmsAliasName: jsii.String("kmsAliasName"),
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
//   	},
//   	StreamsKmsRole: jsii.String("streamsKmsRole"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TelemetrySinkConfigArn: jsii.String("telemetrySinkConfigArn"),
//   	TrackingConfigArn: jsii.String("trackingConfigArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-groundstation-missionprofile.html
//
type CfnMissionProfileMixinProps struct {
	// Amount of time in seconds after a contact ends that you’d like to receive a Ground Station Contact State Change indicating the pass has finished.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-groundstation-missionprofile.html#cfn-groundstation-missionprofile-contactpostpassdurationseconds
	//
	ContactPostPassDurationSeconds *float64 `field:"optional" json:"contactPostPassDurationSeconds" yaml:"contactPostPassDurationSeconds"`
	// Amount of time in seconds prior to contact start that you'd like to receive a Ground Station Contact State Change Event indicating an upcoming pass.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-groundstation-missionprofile.html#cfn-groundstation-missionprofile-contactprepassdurationseconds
	//
	ContactPrePassDurationSeconds *float64 `field:"optional" json:"contactPrePassDurationSeconds" yaml:"contactPrePassDurationSeconds"`
	// A list containing lists of config ARNs.
	//
	// Each list of config ARNs is an edge, with a "from" config and a "to" config.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-groundstation-missionprofile.html#cfn-groundstation-missionprofile-dataflowedges
	//
	DataflowEdges interface{} `field:"optional" json:"dataflowEdges" yaml:"dataflowEdges"`
	// Minimum length of a contact in seconds that Ground Station will return when listing contacts.
	//
	// Ground Station will not return contacts shorter than this duration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-groundstation-missionprofile.html#cfn-groundstation-missionprofile-minimumviablecontactdurationseconds
	//
	MinimumViableContactDurationSeconds *float64 `field:"optional" json:"minimumViableContactDurationSeconds" yaml:"minimumViableContactDurationSeconds"`
	// The name of the mission profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-groundstation-missionprofile.html#cfn-groundstation-missionprofile-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// KMS key to use for encrypting streams.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-groundstation-missionprofile.html#cfn-groundstation-missionprofile-streamskmskey
	//
	StreamsKmsKey interface{} `field:"optional" json:"streamsKmsKey" yaml:"streamsKmsKey"`
	// Role to use for encrypting streams with KMS key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-groundstation-missionprofile.html#cfn-groundstation-missionprofile-streamskmsrole
	//
	StreamsKmsRole *string `field:"optional" json:"streamsKmsRole" yaml:"streamsKmsRole"`
	// Tags assigned to the mission profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-groundstation-missionprofile.html#cfn-groundstation-missionprofile-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// ARN of a Config resource of type TelemetrySinkConfig used for telemetry data sink configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-groundstation-missionprofile.html#cfn-groundstation-missionprofile-telemetrysinkconfigarn
	//
	TelemetrySinkConfigArn *string `field:"optional" json:"telemetrySinkConfigArn" yaml:"telemetrySinkConfigArn"`
	// The ARN of a tracking config objects that defines how to track the satellite through the sky during a contact.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-groundstation-missionprofile.html#cfn-groundstation-missionprofile-trackingconfigarn
	//
	TrackingConfigArn *string `field:"optional" json:"trackingConfigArn" yaml:"trackingConfigArn"`
}

