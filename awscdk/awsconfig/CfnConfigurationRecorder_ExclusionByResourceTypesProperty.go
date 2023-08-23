package awsconfig


// Specifies whether the configuration recorder excludes resource types from being recorded.
//
// Use the `resourceTypes` field to enter a comma-separated list of resource types to exclude as exemptions.
//
// > To use this option, you must set the `useOnly` field of [AWS::Config::ConfigurationRecorder RecordingStrategy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationrecorder-recordingstrategy.html) to `EXCLUSION_BY_RESOURCE_TYPES` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   exclusionByResourceTypesProperty := &ExclusionByResourceTypesProperty{
//   	ResourceTypes: []*string{
//   		jsii.String("resourceTypes"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationrecorder-exclusionbyresourcetypes.html
//
type CfnConfigurationRecorder_ExclusionByResourceTypesProperty struct {
	// A comma-separated list of resource types to exclude from recording by the configuration recorder.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationrecorder-exclusionbyresourcetypes.html#cfn-config-configurationrecorder-exclusionbyresourcetypes-resourcetypes
	//
	ResourceTypes *[]*string `field:"required" json:"resourceTypes" yaml:"resourceTypes"`
}

