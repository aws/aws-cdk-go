package awsconfig


// Specifies which AWS resource types AWS Config records for configuration changes.
//
// In the recording group, you specify whether you want to record all supported resource types or only specific types of resources.
//
// By default, AWS Config records the configuration changes for all supported types of *regional resources* that AWS Config discovers in the region in which it is running. Regional resources are tied to a region and can be used only in that region. Examples of regional resources are EC2 instances and EBS volumes.
//
// You can also have AWS Config record supported types of *global resources* . Global resources are not tied to a specific region and can be used in all regions. The global resource types that AWS Config supports include IAM users, groups, roles, and customer managed policies.
//
// > Global resource types onboarded to AWS Config recording after February 2022 will only be recorded in the service's home region for the commercial partition and AWS GovCloud (US) West for the GovCloud partition. You can view the Configuration Items for these new global resource types only in their home region and AWS GovCloud (US) West.
// >
// > Supported global resource types onboarded before February 2022 such as `AWS::IAM::Group` , `AWS::IAM::Policy` , `AWS::IAM::Role` , `AWS::IAM::User` remain unchanged, and they will continue to deliver Configuration Items in all supported regions in AWS Config . The change will only affect new global resource types onboarded after February 2022.
// >
// > To record global resource types onboarded after February 2022, enable All Supported Resource Types in the home region of the global resource type you want to record.
//
// If you don't want AWS Config to record all resources, you can specify which types of resources it will record with the `resourceTypes` parameter.
//
// For a list of supported resource types, see [Supported Resource Types](https://docs.aws.amazon.com/config/latest/developerguide/resource-config-reference.html#supported-resources) .
//
// For more information and a table of the Home Regions for Global Resource Types Onboarded after February 2022, see [Selecting Which Resources AWS Config Records](https://docs.aws.amazon.com/config/latest/developerguide/select-resources.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   recordingGroupProperty := &RecordingGroupProperty{
//   	AllSupported: jsii.Boolean(false),
//   	IncludeGlobalResourceTypes: jsii.Boolean(false),
//   	ResourceTypes: []*string{
//   		jsii.String("resourceTypes"),
//   	},
//   }
//
type CfnConfigurationRecorder_RecordingGroupProperty struct {
	// Specifies whether AWS Config records configuration changes for every supported type of regional resource.
	//
	// If you set this option to `true` , when AWS Config adds support for a new type of regional resource, it starts recording resources of that type automatically.
	//
	// If you set this option to `true` , you cannot enumerate a list of `resourceTypes` .
	AllSupported interface{} `field:"optional" json:"allSupported" yaml:"allSupported"`
	// Specifies whether AWS Config includes all supported types of global resources (for example, IAM resources) with the resources that it records.
	//
	// Before you can set this option to `true` , you must set the `AllSupported` option to `true` .
	//
	// If you set this option to `true` , when AWS Config adds support for a new type of global resource, it starts recording resources of that type automatically.
	//
	// The configuration details for any global resource are the same in all regions. To prevent duplicate configuration items, you should consider customizing AWS Config in only one region to record global resources.
	IncludeGlobalResourceTypes interface{} `field:"optional" json:"includeGlobalResourceTypes" yaml:"includeGlobalResourceTypes"`
	// A comma-separated list that specifies the types of AWS resources for which AWS Config records configuration changes (for example, `AWS::EC2::Instance` or `AWS::CloudTrail::Trail` ).
	//
	// To record all configuration changes, you must set the `AllSupported` option to `false` .
	//
	// If you set this option to `true` , when AWS Config adds support for a new type of resource, it will not record resources of that type unless you manually add that type to your recording group.
	//
	// For a list of valid `resourceTypes` values, see the *resourceType Value* column in [Supported AWS Resource Types](https://docs.aws.amazon.com/config/latest/developerguide/resource-config-reference.html#supported-resources) .
	ResourceTypes *[]*string `field:"optional" json:"resourceTypes" yaml:"resourceTypes"`
}

