package awsconfig


// Specifies the types of AWS resource for which AWS Config records configuration changes.
//
// In the recording group, you specify whether all supported types or specific types of resources are recorded.
//
// By default, AWS Config records configuration changes for all supported types of regional resources that AWS Config discovers in the region in which it is running. Regional resources are tied to a region and can be used only in that region. Examples of regional resources are EC2 instances and EBS volumes.
//
// You can also have AWS Config record configuration changes for supported types of global resources (for example, IAM resources). Global resources are not tied to an individual region and can be used in all regions.
//
// > The configuration details for any global resource are the same in all regions. If you customize AWS Config in multiple regions to record global resources, it will create multiple configuration items each time a global resource changes: one configuration item for each region. These configuration items will contain identical data. To prevent duplicate configuration items, you should consider customizing AWS Config in only one region to record global resources, unless you want the configuration items to be available in multiple regions.
//
// If you don't want AWS Config to record all resources, you can specify which types of resources it will record with the `resourceTypes` parameter.
//
// For a list of supported resource types, see [Supported Resource Types](https://docs.aws.amazon.com/config/latest/developerguide/resource-config-reference.html#supported-resources) .
//
// For more information, see [Selecting Which Resources AWS Config Records](https://docs.aws.amazon.com/config/latest/developerguide/select-resources.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   recordingGroupProperty := &recordingGroupProperty{
//   	allSupported: jsii.Boolean(false),
//   	includeGlobalResourceTypes: jsii.Boolean(false),
//   	resourceTypes: []*string{
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

