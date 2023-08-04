package awsstepfunctionstasks


// The configuration that defines an instance fleet.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var configurationProperty_ configurationProperty
//   var size size
//
//   instanceFleetConfigProperty := &InstanceFleetConfigProperty{
//   	InstanceFleetType: awscdk.Aws_stepfunctions_tasks.EmrCreateCluster.InstanceRoleType_MASTER,
//
//   	// the properties below are optional
//   	InstanceTypeConfigs: []instanceTypeConfigProperty{
//   		&instanceTypeConfigProperty{
//   			InstanceType: jsii.String("instanceType"),
//
//   			// the properties below are optional
//   			BidPrice: jsii.String("bidPrice"),
//   			BidPriceAsPercentageOfOnDemandPrice: jsii.Number(123),
//   			Configurations: []*configurationProperty{
//   				&configurationProperty{
//   					Classification: jsii.String("classification"),
//   					Configurations: []*configurationProperty{
//   						configurationProperty_,
//   					},
//   					Properties: map[string]*string{
//   						"propertiesKey": jsii.String("properties"),
//   					},
//   				},
//   			},
//   			EbsConfiguration: &EbsConfigurationProperty{
//   				EbsBlockDeviceConfigs: []ebsBlockDeviceConfigProperty{
//   					&ebsBlockDeviceConfigProperty{
//   						VolumeSpecification: &VolumeSpecificationProperty{
//   							VolumeSize: size,
//   							VolumeType: awscdk.*Aws_stepfunctions_tasks.EmrCreateCluster.EbsBlockDeviceVolumeType_GP2,
//
//   							// the properties below are optional
//   							Iops: jsii.Number(123),
//   						},
//
//   						// the properties below are optional
//   						VolumesPerInstance: jsii.Number(123),
//   					},
//   				},
//   				EbsOptimized: jsii.Boolean(false),
//   			},
//   			WeightedCapacity: jsii.Number(123),
//   		},
//   	},
//   	LaunchSpecifications: &InstanceFleetProvisioningSpecificationsProperty{
//   		SpotSpecification: &SpotProvisioningSpecificationProperty{
//   			TimeoutAction: awscdk.*Aws_stepfunctions_tasks.EmrCreateCluster.SpotTimeoutAction_SWITCH_TO_ON_DEMAND,
//   			TimeoutDurationMinutes: jsii.Number(123),
//
//   			// the properties below are optional
//   			AllocationStrategy: awscdk.*Aws_stepfunctions_tasks.EmrCreateCluster.SpotAllocationStrategy_CAPACITY_OPTIMIZED,
//   			BlockDurationMinutes: jsii.Number(123),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	TargetOnDemandCapacity: jsii.Number(123),
//   	TargetSpotCapacity: jsii.Number(123),
//   }
//
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_InstanceFleetConfig.html
//
type EmrCreateCluster_InstanceFleetConfigProperty struct {
	// The node type that the instance fleet hosts.
	//
	// Valid values are MASTER,CORE,and TASK.
	InstanceFleetType EmrCreateCluster_InstanceRoleType `field:"required" json:"instanceFleetType" yaml:"instanceFleetType"`
	// The instance type configurations that define the EC2 instances in the instance fleet.
	// Default: No instanceTpeConfigs.
	//
	InstanceTypeConfigs *[]*EmrCreateCluster_InstanceTypeConfigProperty `field:"optional" json:"instanceTypeConfigs" yaml:"instanceTypeConfigs"`
	// The launch specification for the instance fleet.
	// Default: No launchSpecifications.
	//
	LaunchSpecifications *EmrCreateCluster_InstanceFleetProvisioningSpecificationsProperty `field:"optional" json:"launchSpecifications" yaml:"launchSpecifications"`
	// The friendly name of the instance fleet.
	// Default: No name.
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The target capacity of On-Demand units for the instance fleet, which determines how many On-Demand instances to provision.
	// Default: No targetOnDemandCapacity.
	//
	TargetOnDemandCapacity *float64 `field:"optional" json:"targetOnDemandCapacity" yaml:"targetOnDemandCapacity"`
	// The target capacity of Spot units for the instance fleet, which determines how many Spot instances to provision.
	// Default: No targetSpotCapacity.
	//
	TargetSpotCapacity *float64 `field:"optional" json:"targetSpotCapacity" yaml:"targetSpotCapacity"`
}

