package awsconfig


// Specifies which resource types AWS Config records for configuration changes.
//
// By default, AWS Config records configuration changes for all current and future supported resource types in the AWS Region where you have enabled AWS Config , excluding the global IAM resource types: IAM users, groups, roles, and customer managed policies.
//
// In the recording group, you specify whether you want to record all supported current and future supported resource types or to include or exclude specific resources types. For a list of supported resource types, see [Supported Resource Types](https://docs.aws.amazon.com/config/latest/developerguide/resource-config-reference.html#supported-resources) in the *AWS Config developer guide* .
//
// If you don't want AWS Config to record all current and future supported resource types (excluding the global IAM resource types), use one of the following recording strategies:
//
// - *Record all current and future resource types with exclusions* ( `EXCLUSION_BY_RESOURCE_TYPES` ), or
// - *Record specific resource types* ( `INCLUSION_BY_RESOURCE_TYPES` ).
//
// If you use the recording strategy to *Record all current and future resource types* ( `ALL_SUPPORTED_RESOURCE_TYPES` ), you can use the flag `includeGlobalResourceTypes` to include the global IAM resource types in your recording.
//
// > *Aurora global clusters are recorded in all enabled Regions*
// >
// > The `AWS::RDS::GlobalCluster` resource type will be recorded in all supported AWS Config Regions where the configuration recorder is enabled.
// >
// > If you do not want to record `AWS::RDS::GlobalCluster` in all enabled Regions, use the `EXCLUSION_BY_RESOURCE_TYPES` or `INCLUSION_BY_RESOURCE_TYPES` recording strategy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   recordingGroupProperty := &RecordingGroupProperty{
//   	AllSupported: jsii.Boolean(false),
//   	ExclusionByResourceTypes: &ExclusionByResourceTypesProperty{
//   		ResourceTypes: []*string{
//   			jsii.String("resourceTypes"),
//   		},
//   	},
//   	IncludeGlobalResourceTypes: jsii.Boolean(false),
//   	RecordingStrategy: &RecordingStrategyProperty{
//   		UseOnly: jsii.String("useOnly"),
//   	},
//   	ResourceTypes: []*string{
//   		jsii.String("resourceTypes"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationrecorder-recordinggroup.html
//
type CfnConfigurationRecorder_RecordingGroupProperty struct {
	// Specifies whether AWS Config records configuration changes for all supported resource types, excluding the global IAM resource types.
	//
	// If you set this field to `true` , when AWS Config adds support for a new resource type, AWS Config starts recording resources of that type automatically.
	//
	// If you set this field to `true` , you cannot enumerate specific resource types to record in the `resourceTypes` field of [RecordingGroup](https://docs.aws.amazon.com/config/latest/APIReference/API_RecordingGroup.html) , or to exclude in the `resourceTypes` field of [ExclusionByResourceTypes](https://docs.aws.amazon.com/config/latest/APIReference/API_ExclusionByResourceTypes.html) .
	//
	// > *Region availability*
	// >
	// > Check [Resource Coverage by Region Availability](https://docs.aws.amazon.com/config/latest/developerguide/what-is-resource-config-coverage.html) to see if a resource type is supported in the AWS Region where you set up AWS Config .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationrecorder-recordinggroup.html#cfn-config-configurationrecorder-recordinggroup-allsupported
	//
	AllSupported interface{} `field:"optional" json:"allSupported" yaml:"allSupported"`
	// An object that specifies how AWS Config excludes resource types from being recorded by the configuration recorder.
	//
	// To use this option, you must set the `useOnly` field of [AWS::Config::ConfigurationRecorder RecordingStrategy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationrecorder-recordingstrategy.html) to `EXCLUSION_BY_RESOURCE_TYPES` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationrecorder-recordinggroup.html#cfn-config-configurationrecorder-recordinggroup-exclusionbyresourcetypes
	//
	ExclusionByResourceTypes interface{} `field:"optional" json:"exclusionByResourceTypes" yaml:"exclusionByResourceTypes"`
	// This option is a bundle which only applies to the global IAM resource types: IAM users, groups, roles, and customer managed policies.
	//
	// These global IAM resource types can only be recorded by AWS Config in Regions where AWS Config was available before February 2022. You cannot be record the global IAM resouce types in Regions supported by AWS Config after February 2022. This list where you cannot record the global IAM resource types includes the following Regions:
	//
	// - Asia Pacific (Hyderabad)
	// - Asia Pacific (Melbourne)
	// - Europe (Spain)
	// - Europe (Zurich)
	// - Israel (Tel Aviv)
	// - Middle East (UAE)
	//
	// > *Aurora global clusters are recorded in all enabled Regions*
	// >
	// > The `AWS::RDS::GlobalCluster` resource type will be recorded in all supported AWS Config Regions where the configuration recorder is enabled, even if `includeGlobalResourceTypes` is not set to `true` . The `includeGlobalResourceTypes` option is a bundle which only applies to IAM users, groups, roles, and customer managed policies.
	// >
	// > If you do not want to record `AWS::RDS::GlobalCluster` in all enabled Regions, use one of the following recording strategies:
	// >
	// > - *Record all current and future resource types with exclusions* ( `EXCLUSION_BY_RESOURCE_TYPES` ), or
	// > - *Record specific resource types* ( `INCLUSION_BY_RESOURCE_TYPES` ).
	// >
	// > For more information, see [Selecting Which Resources are Recorded](https://docs.aws.amazon.com/config/latest/developerguide/select-resources.html#select-resources-all) in the *AWS Config developer guide* . > Before you set this field to `true` , set the `allSupported` field of [RecordingGroup](https://docs.aws.amazon.com/config/latest/APIReference/API_RecordingGroup.html) to `true` . Optionally, you can set the `useOnly` field of [RecordingStrategy](https://docs.aws.amazon.com/config/latest/APIReference/API_RecordingStrategy.html) to `ALL_SUPPORTED_RESOURCE_TYPES` . > *Overriding fields*
	// >
	// > If you set this field to `false` but list global IAM resource types in the `resourceTypes` field of [RecordingGroup](https://docs.aws.amazon.com/config/latest/APIReference/API_RecordingGroup.html) , AWS Config will still record configuration changes for those specified resource types *regardless* of if you set the `includeGlobalResourceTypes` field to false.
	// >
	// > If you do not want to record configuration changes to the global IAM resource types (IAM users, groups, roles, and customer managed policies), make sure to not list them in the `resourceTypes` field in addition to setting the `includeGlobalResourceTypes` field to false.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationrecorder-recordinggroup.html#cfn-config-configurationrecorder-recordinggroup-includeglobalresourcetypes
	//
	IncludeGlobalResourceTypes interface{} `field:"optional" json:"includeGlobalResourceTypes" yaml:"includeGlobalResourceTypes"`
	// An object that specifies the recording strategy for the configuration recorder.
	//
	// - If you set the `useOnly` field of [RecordingStrategy](https://docs.aws.amazon.com/config/latest/APIReference/API_RecordingStrategy.html) to `ALL_SUPPORTED_RESOURCE_TYPES` , AWS Config records configuration changes for all supported resource types, excluding the global IAM resource types. You also must set the `allSupported` field of [RecordingGroup](https://docs.aws.amazon.com/config/latest/APIReference/API_RecordingGroup.html) to `true` . When AWS Config adds support for a new resource type, AWS Config automatically starts recording resources of that type.
	// - If you set the `useOnly` field of [RecordingStrategy](https://docs.aws.amazon.com/config/latest/APIReference/API_RecordingStrategy.html) to `INCLUSION_BY_RESOURCE_TYPES` , AWS Config records configuration changes for only the resource types you specify in the `resourceTypes` field of [RecordingGroup](https://docs.aws.amazon.com/config/latest/APIReference/API_RecordingGroup.html) .
	// - If you set the `useOnly` field of [RecordingStrategy](https://docs.aws.amazon.com/config/latest/APIReference/API_RecordingStrategy.html) to `EXCLUSION_BY_RESOURCE_TYPES` , AWS Config records configuration changes for all supported resource types except the resource types that you specify to exclude from being recorded in the `resourceTypes` field of [ExclusionByResourceTypes](https://docs.aws.amazon.com/config/latest/APIReference/API_ExclusionByResourceTypes.html) .
	//
	// > *Required and optional fields*
	// >
	// > The `recordingStrategy` field is optional when you set the `allSupported` field of [RecordingGroup](https://docs.aws.amazon.com/config/latest/APIReference/API_RecordingGroup.html) to `true` .
	// >
	// > The `recordingStrategy` field is optional when you list resource types in the `resourceTypes` field of [RecordingGroup](https://docs.aws.amazon.com/config/latest/APIReference/API_RecordingGroup.html) .
	// >
	// > The `recordingStrategy` field is required if you list resource types to exclude from recording in the `resourceTypes` field of [ExclusionByResourceTypes](https://docs.aws.amazon.com/config/latest/APIReference/API_ExclusionByResourceTypes.html) . > *Overriding fields*
	// >
	// > If you choose `EXCLUSION_BY_RESOURCE_TYPES` for the recording strategy, the `exclusionByResourceTypes` field will override other properties in the request.
	// >
	// > For example, even if you set `includeGlobalResourceTypes` to false, global IAM resource types will still be automatically recorded in this option unless those resource types are specifically listed as exclusions in the `resourceTypes` field of `exclusionByResourceTypes` . > *Global resources types and the resource exclusion recording strategy*
	// >
	// > By default, if you choose the `EXCLUSION_BY_RESOURCE_TYPES` recording strategy, when AWS Config adds support for a new resource type in the Region where you set up the configuration recorder, including global resource types, AWS Config starts recording resources of that type automatically.
	// >
	// > Unless specifically listed as exclusions, `AWS::RDS::GlobalCluster` will be recorded automatically in all supported AWS Config Regions were the configuration recorder is enabled.
	// >
	// > IAM users, groups, roles, and customer managed policies will be recorded in the Region where you set up the configuration recorder if that is a Region where AWS Config was available before February 2022. You cannot be record the global IAM resouce types in Regions supported by AWS Config after February 2022. This list where you cannot record the global IAM resource types includes the following Regions:
	// >
	// > - Asia Pacific (Hyderabad)
	// > - Asia Pacific (Melbourne)
	// > - Europe (Spain)
	// > - Europe (Zurich)
	// > - Israel (Tel Aviv)
	// > - Middle East (UAE).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationrecorder-recordinggroup.html#cfn-config-configurationrecorder-recordinggroup-recordingstrategy
	//
	RecordingStrategy interface{} `field:"optional" json:"recordingStrategy" yaml:"recordingStrategy"`
	// A comma-separated list that specifies which resource types AWS Config records.
	//
	// For a list of valid `resourceTypes` values, see the *Resource Type Value* column in [Supported AWS resource Types](https://docs.aws.amazon.com/config/latest/developerguide/resource-config-reference.html#supported-resources) in the *AWS Config developer guide* .
	//
	// > *Required and optional fields*
	// >
	// > Optionally, you can set the `useOnly` field of [RecordingStrategy](https://docs.aws.amazon.com/config/latest/APIReference/API_RecordingStrategy.html) to `INCLUSION_BY_RESOURCE_TYPES` .
	// >
	// > To record all configuration changes, set the `allSupported` field of [RecordingGroup](https://docs.aws.amazon.com/config/latest/APIReference/API_RecordingGroup.html) to `true` , and either omit this field or don't specify any resource types in this field. If you set the `allSupported` field to `false` and specify values for `resourceTypes` , when AWS Config adds support for a new type of resource, it will not record resources of that type unless you manually add that type to your recording group. > *Region availability*
	// >
	// > Before specifying a resource type for AWS Config to track, check [Resource Coverage by Region Availability](https://docs.aws.amazon.com/config/latest/developerguide/what-is-resource-config-coverage.html) to see if the resource type is supported in the AWS Region where you set up AWS Config . If a resource type is supported by AWS Config in at least one Region, you can enable the recording of that resource type in all Regions supported by AWS Config , even if the specified resource type is not supported in the AWS Region where you set up AWS Config .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationrecorder-recordinggroup.html#cfn-config-configurationrecorder-recordinggroup-resourcetypes
	//
	ResourceTypes *[]*string `field:"optional" json:"resourceTypes" yaml:"resourceTypes"`
}

