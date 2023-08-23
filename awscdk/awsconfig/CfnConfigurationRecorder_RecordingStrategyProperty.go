package awsconfig


// Specifies the recording strategy of the configuration recorder.
//
// Valid values include: `ALL_SUPPORTED_RESOURCE_TYPES` , `INCLUSION_BY_RESOURCE_TYPES` , and `EXCLUSION_BY_RESOURCE_TYPES` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   recordingStrategyProperty := &RecordingStrategyProperty{
//   	UseOnly: jsii.String("useOnly"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationrecorder-recordingstrategy.html
//
type CfnConfigurationRecorder_RecordingStrategyProperty struct {
	// The recording strategy for the configuration recorder.
	//
	// - If you set this option to `ALL_SUPPORTED_RESOURCE_TYPES` , AWS Config records configuration changes for all supported regional resource types. You also must set the `allSupported` field of [AWS::Config::ConfigurationRecorder RecordingGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationrecorder-recordinggroup.html) to `true` .
	//
	// When AWS Config adds support for a new type of regional resource, AWS Config automatically starts recording resources of that type. For a list of supported resource types, see [Supported Resource Types](https://docs.aws.amazon.com/config/latest/developerguide/resource-config-reference.html#supported-resources) in the *AWS Config developer guide* .
	// - If you set this option to `INCLUSION_BY_RESOURCE_TYPES` , AWS Config records configuration changes for only the resource types that you specify in the `resourceTypes` field of [AWS::Config::ConfigurationRecorder RecordingGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationrecorder-recordinggroup.html) .
	// - If you set this option to `EXCLUSION_BY_RESOURCE_TYPES` , AWS Config records configuration changes for all supported resource types, except the resource types that you specify as exemptions to exclude from being recorded in the `resourceTypes` field of [AWS::Config::ConfigurationRecorder ExclusionByResourceTypes](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationrecorder-exclusionbyresourcetypes.html) .
	//
	// > The `recordingStrategy` field is optional when you set the `allSupported` field of [AWS::Config::ConfigurationRecorder RecordingGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationrecorder-recordinggroup.html) to `true` .
	// >
	// > The `recordingStrategy` field is optional when you list resource types in the `resourceTypes` field of [AWS::Config::ConfigurationRecorder RecordingGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationrecorder-recordinggroup.html) .
	// >
	// > The `recordingStrategy` field is required if you list resource types to exclude from recording in the `resourceTypes` field of [AWS::Config::ConfigurationRecorder ExclusionByResourceTypes](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationrecorder-exclusionbyresourcetypes.html) . > If you choose `EXCLUSION_BY_RESOURCE_TYPES` for the recording strategy, the `exclusionByResourceTypes` field will override other properties in the request.
	// >
	// > For example, even if you set `includeGlobalResourceTypes` to false, global resource types will still be automatically recorded in this option unless those resource types are specifically listed as exemptions in the `resourceTypes` field of `exclusionByResourceTypes` .
	// >
	// > By default, if you choose the `EXCLUSION_BY_RESOURCE_TYPES` recording strategy, when AWS Config adds support for a new resource type in the Region where you set up the configuration recorder, including global resource types, AWS Config starts recording resources of that type automatically.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationrecorder-recordingstrategy.html#cfn-config-configurationrecorder-recordingstrategy-useonly
	//
	UseOnly *string `field:"required" json:"useOnly" yaml:"useOnly"`
}

