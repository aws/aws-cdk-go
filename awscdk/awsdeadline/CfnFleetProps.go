package awsdeadline

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnFleet`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFleetProps := &CfnFleetProps{
//   	Configuration: &FleetConfigurationProperty{
//   		CustomerManaged: &CustomerManagedFleetConfigurationProperty{
//   			Mode: jsii.String("mode"),
//   			WorkerCapabilities: &CustomerManagedWorkerCapabilitiesProperty{
//   				CpuArchitectureType: jsii.String("cpuArchitectureType"),
//   				MemoryMiB: &MemoryMiBRangeProperty{
//   					Min: jsii.Number(123),
//
//   					// the properties below are optional
//   					Max: jsii.Number(123),
//   				},
//   				OsFamily: jsii.String("osFamily"),
//   				VCpuCount: &VCpuCountRangeProperty{
//   					Min: jsii.Number(123),
//
//   					// the properties below are optional
//   					Max: jsii.Number(123),
//   				},
//
//   				// the properties below are optional
//   				AcceleratorCount: &AcceleratorCountRangeProperty{
//   					Min: jsii.Number(123),
//
//   					// the properties below are optional
//   					Max: jsii.Number(123),
//   				},
//   				AcceleratorTotalMemoryMiB: &AcceleratorTotalMemoryMiBRangeProperty{
//   					Min: jsii.Number(123),
//
//   					// the properties below are optional
//   					Max: jsii.Number(123),
//   				},
//   				AcceleratorTypes: []*string{
//   					jsii.String("acceleratorTypes"),
//   				},
//   				CustomAmounts: []interface{}{
//   					&FleetAmountCapabilityProperty{
//   						Min: jsii.Number(123),
//   						Name: jsii.String("name"),
//
//   						// the properties below are optional
//   						Max: jsii.Number(123),
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
//   			},
//
//   			// the properties below are optional
//   			StorageProfileId: jsii.String("storageProfileId"),
//   		},
//   		ServiceManagedEc2: &ServiceManagedEc2FleetConfigurationProperty{
//   			InstanceCapabilities: &ServiceManagedEc2InstanceCapabilitiesProperty{
//   				CpuArchitectureType: jsii.String("cpuArchitectureType"),
//   				MemoryMiB: &MemoryMiBRangeProperty{
//   					Min: jsii.Number(123),
//
//   					// the properties below are optional
//   					Max: jsii.Number(123),
//   				},
//   				OsFamily: jsii.String("osFamily"),
//   				VCpuCount: &VCpuCountRangeProperty{
//   					Min: jsii.Number(123),
//
//   					// the properties below are optional
//   					Max: jsii.Number(123),
//   				},
//
//   				// the properties below are optional
//   				AcceleratorCapabilities: &AcceleratorCapabilitiesProperty{
//   					Selections: []interface{}{
//   						&AcceleratorSelectionProperty{
//   							Name: jsii.String("name"),
//
//   							// the properties below are optional
//   							Runtime: jsii.String("runtime"),
//   						},
//   					},
//
//   					// the properties below are optional
//   					Count: &AcceleratorCountRangeProperty{
//   						Min: jsii.Number(123),
//
//   						// the properties below are optional
//   						Max: jsii.Number(123),
//   					},
//   				},
//   				AllowedInstanceTypes: []*string{
//   					jsii.String("allowedInstanceTypes"),
//   				},
//   				CustomAmounts: []interface{}{
//   					&FleetAmountCapabilityProperty{
//   						Min: jsii.Number(123),
//   						Name: jsii.String("name"),
//
//   						// the properties below are optional
//   						Max: jsii.Number(123),
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
//   				RootEbsVolume: &Ec2EbsVolumeProperty{
//   					Iops: jsii.Number(123),
//   					SizeGiB: jsii.Number(123),
//   					ThroughputMiB: jsii.Number(123),
//   				},
//   			},
//   			InstanceMarketOptions: &ServiceManagedEc2InstanceMarketOptionsProperty{
//   				Type: jsii.String("type"),
//   			},
//   		},
//   	},
//   	DisplayName: jsii.String("displayName"),
//   	FarmId: jsii.String("farmId"),
//   	MaxWorkerCount: jsii.Number(123),
//   	RoleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	MinWorkerCount: jsii.Number(123),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-fleet.html
//
type CfnFleetProps struct {
	// The configuration details for the fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-fleet.html#cfn-deadline-fleet-configuration
	//
	Configuration interface{} `field:"required" json:"configuration" yaml:"configuration"`
	// The display name of the fleet summary to update.
	//
	// > This field can store any content. Escape or encode this content before displaying it on a webpage or any other system that might interpret the content of this field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-fleet.html#cfn-deadline-fleet-displayname
	//
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The farm ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-fleet.html#cfn-deadline-fleet-farmid
	//
	FarmId *string `field:"required" json:"farmId" yaml:"farmId"`
	// The maximum number of workers specified in the fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-fleet.html#cfn-deadline-fleet-maxworkercount
	//
	MaxWorkerCount *float64 `field:"required" json:"maxWorkerCount" yaml:"maxWorkerCount"`
	// The IAM role that workers in the fleet use when processing jobs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-fleet.html#cfn-deadline-fleet-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// A description that helps identify what the fleet is used for.
	//
	// > This field can store any content. Escape or encode this content before displaying it on a webpage or any other system that might interpret the content of this field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-fleet.html#cfn-deadline-fleet-description
	//
	// Default: - "".
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The minimum number of workers in the fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-fleet.html#cfn-deadline-fleet-minworkercount
	//
	// Default: - 0.
	//
	MinWorkerCount *float64 `field:"optional" json:"minWorkerCount" yaml:"minWorkerCount"`
	// The tags to add to your fleet.
	//
	// Each tag consists of a tag key and a tag value. Tag keys and values are both required, but tag values can be empty strings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-fleet.html#cfn-deadline-fleet-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

