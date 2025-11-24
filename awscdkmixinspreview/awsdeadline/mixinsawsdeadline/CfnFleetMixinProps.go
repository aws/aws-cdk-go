package mixinsawsdeadline

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnFleetPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnFleetMixinProps := &CfnFleetMixinProps{
//   	Configuration: &FleetConfigurationProperty{
//   		CustomerManaged: &CustomerManagedFleetConfigurationProperty{
//   			Mode: jsii.String("mode"),
//   			StorageProfileId: jsii.String("storageProfileId"),
//   			TagPropagationMode: jsii.String("tagPropagationMode"),
//   			WorkerCapabilities: &CustomerManagedWorkerCapabilitiesProperty{
//   				AcceleratorCount: &AcceleratorCountRangeProperty{
//   					Max: jsii.Number(123),
//   					Min: jsii.Number(123),
//   				},
//   				AcceleratorTotalMemoryMiB: &AcceleratorTotalMemoryMiBRangeProperty{
//   					Max: jsii.Number(123),
//   					Min: jsii.Number(123),
//   				},
//   				AcceleratorTypes: []*string{
//   					jsii.String("acceleratorTypes"),
//   				},
//   				CpuArchitectureType: jsii.String("cpuArchitectureType"),
//   				CustomAmounts: []interface{}{
//   					&FleetAmountCapabilityProperty{
//   						Max: jsii.Number(123),
//   						Min: jsii.Number(123),
//   						Name: jsii.String("name"),
//   					},
//   				},
//   				CustomAttributes: []interface{}{
//   					&FleetAttributeCapabilityProperty{
//   						Name: jsii.String("name"),
//   						Values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   				},
//   				MemoryMiB: &MemoryMiBRangeProperty{
//   					Max: jsii.Number(123),
//   					Min: jsii.Number(123),
//   				},
//   				OsFamily: jsii.String("osFamily"),
//   				VCpuCount: &VCpuCountRangeProperty{
//   					Max: jsii.Number(123),
//   					Min: jsii.Number(123),
//   				},
//   			},
//   		},
//   		ServiceManagedEc2: &ServiceManagedEc2FleetConfigurationProperty{
//   			InstanceCapabilities: &ServiceManagedEc2InstanceCapabilitiesProperty{
//   				AcceleratorCapabilities: &AcceleratorCapabilitiesProperty{
//   					Count: &AcceleratorCountRangeProperty{
//   						Max: jsii.Number(123),
//   						Min: jsii.Number(123),
//   					},
//   					Selections: []interface{}{
//   						&AcceleratorSelectionProperty{
//   							Name: jsii.String("name"),
//   							Runtime: jsii.String("runtime"),
//   						},
//   					},
//   				},
//   				AllowedInstanceTypes: []*string{
//   					jsii.String("allowedInstanceTypes"),
//   				},
//   				CpuArchitectureType: jsii.String("cpuArchitectureType"),
//   				CustomAmounts: []interface{}{
//   					&FleetAmountCapabilityProperty{
//   						Max: jsii.Number(123),
//   						Min: jsii.Number(123),
//   						Name: jsii.String("name"),
//   					},
//   				},
//   				CustomAttributes: []interface{}{
//   					&FleetAttributeCapabilityProperty{
//   						Name: jsii.String("name"),
//   						Values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   				},
//   				ExcludedInstanceTypes: []*string{
//   					jsii.String("excludedInstanceTypes"),
//   				},
//   				MemoryMiB: &MemoryMiBRangeProperty{
//   					Max: jsii.Number(123),
//   					Min: jsii.Number(123),
//   				},
//   				OsFamily: jsii.String("osFamily"),
//   				RootEbsVolume: &Ec2EbsVolumeProperty{
//   					Iops: jsii.Number(123),
//   					SizeGiB: jsii.Number(123),
//   					ThroughputMiB: jsii.Number(123),
//   				},
//   				VCpuCount: &VCpuCountRangeProperty{
//   					Max: jsii.Number(123),
//   					Min: jsii.Number(123),
//   				},
//   			},
//   			InstanceMarketOptions: &ServiceManagedEc2InstanceMarketOptionsProperty{
//   				Type: jsii.String("type"),
//   			},
//   			StorageProfileId: jsii.String("storageProfileId"),
//   			VpcConfiguration: &VpcConfigurationProperty{
//   				ResourceConfigurationArns: []*string{
//   					jsii.String("resourceConfigurationArns"),
//   				},
//   			},
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	DisplayName: jsii.String("displayName"),
//   	FarmId: jsii.String("farmId"),
//   	HostConfiguration: &HostConfigurationProperty{
//   		ScriptBody: jsii.String("scriptBody"),
//   		ScriptTimeoutSeconds: jsii.Number(123),
//   	},
//   	MaxWorkerCount: jsii.Number(123),
//   	MinWorkerCount: jsii.Number(123),
//   	RoleArn: jsii.String("roleArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-fleet.html
//
type CfnFleetMixinProps struct {
	// The configuration details for the fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-fleet.html#cfn-deadline-fleet-configuration
	//
	Configuration interface{} `field:"optional" json:"configuration" yaml:"configuration"`
	// A description that helps identify what the fleet is used for.
	//
	// > This field can store any content. Escape or encode this content before displaying it on a webpage or any other system that might interpret the content of this field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-fleet.html#cfn-deadline-fleet-description
	//
	// Default: - "".
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The display name of the fleet summary to update.
	//
	// > This field can store any content. Escape or encode this content before displaying it on a webpage or any other system that might interpret the content of this field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-fleet.html#cfn-deadline-fleet-displayname
	//
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The farm ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-fleet.html#cfn-deadline-fleet-farmid
	//
	FarmId *string `field:"optional" json:"farmId" yaml:"farmId"`
	// Provides a script that runs as a worker is starting up that you can use to provide additional configuration for workers in your fleet.
	//
	// To remove a script from a fleet, use the [UpdateFleet](https://docs.aws.amazon.com/deadline-cloud/latest/APIReference/API_UpdateFleet.html) operation with the `hostConfiguration` `scriptBody` parameter set to an empty string ("").
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-fleet.html#cfn-deadline-fleet-hostconfiguration
	//
	HostConfiguration interface{} `field:"optional" json:"hostConfiguration" yaml:"hostConfiguration"`
	// The maximum number of workers specified in the fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-fleet.html#cfn-deadline-fleet-maxworkercount
	//
	MaxWorkerCount *float64 `field:"optional" json:"maxWorkerCount" yaml:"maxWorkerCount"`
	// The minimum number of workers in the fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-fleet.html#cfn-deadline-fleet-minworkercount
	//
	// Default: - 0.
	//
	MinWorkerCount *float64 `field:"optional" json:"minWorkerCount" yaml:"minWorkerCount"`
	// The IAM role that workers in the fleet use when processing jobs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-fleet.html#cfn-deadline-fleet-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The tags to add to your fleet.
	//
	// Each tag consists of a tag key and a tag value. Tag keys and values are both required, but tag values can be empty strings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-fleet.html#cfn-deadline-fleet-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

