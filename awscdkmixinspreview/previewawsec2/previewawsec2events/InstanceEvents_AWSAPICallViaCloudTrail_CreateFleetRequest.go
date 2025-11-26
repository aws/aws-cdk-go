package previewawsec2events


// Type definition for CreateFleetRequest.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var overrides interface{}
//
//   createFleetRequest := &CreateFleetRequest{
//   	ClientToken: []*string{
//   		jsii.String("clientToken"),
//   	},
//   	ExistingInstances: &ExistingInstances{
//   		AvailabilityZone: []*string{
//   			jsii.String("availabilityZone"),
//   		},
//   		Count: []*string{
//   			jsii.String("count"),
//   		},
//   		InstanceType: []*string{
//   			jsii.String("instanceType"),
//   		},
//   		MarketOption: []*string{
//   			jsii.String("marketOption"),
//   		},
//   		OperatingSystem: []*string{
//   			jsii.String("operatingSystem"),
//   		},
//   		Tag: []*string{
//   			jsii.String("tag"),
//   		},
//   	},
//   	LaunchTemplateConfigs: &LaunchTemplateConfigs{
//   		LaunchTemplateSpecification: &LaunchTemplateSpecification{
//   			LaunchTemplateId: []*string{
//   				jsii.String("launchTemplateId"),
//   			},
//   			Version: []*string{
//   				jsii.String("version"),
//   			},
//   		},
//   		Overrides: []interface{}{
//   			overrides,
//   		},
//   		Tag: []*string{
//   			jsii.String("tag"),
//   		},
//   	},
//   	OnDemandOptions: &OnDemandOptions{
//   		AllocationStrategy: []*string{
//   			jsii.String("allocationStrategy"),
//   		},
//   		InstancePoolConstraintFilterDisabled: []*string{
//   			jsii.String("instancePoolConstraintFilterDisabled"),
//   		},
//   		MaxInstanceCount: []*string{
//   			jsii.String("maxInstanceCount"),
//   		},
//   		MaxTargetCapacity: []*string{
//   			jsii.String("maxTargetCapacity"),
//   		},
//   	},
//   	SpotOptions: &SpotOptions{
//   		AllocationStrategy: []*string{
//   			jsii.String("allocationStrategy"),
//   		},
//   		InstancePoolConstraintFilterDisabled: []*string{
//   			jsii.String("instancePoolConstraintFilterDisabled"),
//   		},
//   		InstancePoolsToUseCount: []*string{
//   			jsii.String("instancePoolsToUseCount"),
//   		},
//   		MaxInstanceCount: []*string{
//   			jsii.String("maxInstanceCount"),
//   		},
//   		MaxTargetCapacity: []*string{
//   			jsii.String("maxTargetCapacity"),
//   		},
//   	},
//   	TagSpecification: &TagSpecification{
//   		ResourceType: []*string{
//   			jsii.String("resourceType"),
//   		},
//   		Tag: &Tag{
//   			Key: []*string{
//   				jsii.String("key"),
//   			},
//   			Tag: []*string{
//   				jsii.String("tag"),
//   			},
//   			Value: []*string{
//   				jsii.String("value"),
//   			},
//   		},
//   	},
//   	TargetCapacitySpecification: &TargetCapacitySpecification{
//   		DefaultTargetCapacityType: []*string{
//   			jsii.String("defaultTargetCapacityType"),
//   		},
//   		OnDemandTargetCapacity: []*string{
//   			jsii.String("onDemandTargetCapacity"),
//   		},
//   		SpotTargetCapacity: []*string{
//   			jsii.String("spotTargetCapacity"),
//   		},
//   		TotalTargetCapacity: []*string{
//   			jsii.String("totalTargetCapacity"),
//   		},
//   	},
//   	Type: []*string{
//   		jsii.String("type"),
//   	},
//   }
//
// Experimental.
type InstanceEvents_AWSAPICallViaCloudTrail_CreateFleetRequest struct {
	// ClientToken property.
	//
	// Specify an array of string values to match this event if the actual value of ClientToken is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ClientToken *[]*string `field:"optional" json:"clientToken" yaml:"clientToken"`
	// ExistingInstances property.
	//
	// Specify an array of string values to match this event if the actual value of ExistingInstances is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ExistingInstances *InstanceEvents_AWSAPICallViaCloudTrail_ExistingInstances `field:"optional" json:"existingInstances" yaml:"existingInstances"`
	// LaunchTemplateConfigs property.
	//
	// Specify an array of string values to match this event if the actual value of LaunchTemplateConfigs is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LaunchTemplateConfigs *InstanceEvents_AWSAPICallViaCloudTrail_LaunchTemplateConfigs `field:"optional" json:"launchTemplateConfigs" yaml:"launchTemplateConfigs"`
	// OnDemandOptions property.
	//
	// Specify an array of string values to match this event if the actual value of OnDemandOptions is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	OnDemandOptions *InstanceEvents_AWSAPICallViaCloudTrail_OnDemandOptions `field:"optional" json:"onDemandOptions" yaml:"onDemandOptions"`
	// SpotOptions property.
	//
	// Specify an array of string values to match this event if the actual value of SpotOptions is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SpotOptions *InstanceEvents_AWSAPICallViaCloudTrail_SpotOptions `field:"optional" json:"spotOptions" yaml:"spotOptions"`
	// TagSpecification property.
	//
	// Specify an array of string values to match this event if the actual value of TagSpecification is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TagSpecification *InstanceEvents_AWSAPICallViaCloudTrail_TagSpecification `field:"optional" json:"tagSpecification" yaml:"tagSpecification"`
	// TargetCapacitySpecification property.
	//
	// Specify an array of string values to match this event if the actual value of TargetCapacitySpecification is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TargetCapacitySpecification *InstanceEvents_AWSAPICallViaCloudTrail_TargetCapacitySpecification `field:"optional" json:"targetCapacitySpecification" yaml:"targetCapacitySpecification"`
	// Type property.
	//
	// Specify an array of string values to match this event if the actual value of Type is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Type *[]*string `field:"optional" json:"type" yaml:"type"`
}

