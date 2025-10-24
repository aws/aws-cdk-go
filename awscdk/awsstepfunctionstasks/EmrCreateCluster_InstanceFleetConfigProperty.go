package awsstepfunctionstasks


// The configuration that defines an instance fleet.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var configurationProperty_ ConfigurationProperty
//   var size Size
//
//   instanceFleetConfigProperty := &InstanceFleetConfigProperty{
//   	InstanceFleetType: awscdk.Aws_stepfunctions_tasks.EmrCreateCluster.InstanceRoleType_MASTER,
//
//   	// the properties below are optional
//   	InstanceTypeConfigs: []InstanceTypeConfigProperty{
//   		&InstanceTypeConfigProperty{
//   			InstanceType: jsii.String("instanceType"),
//
//   			// the properties below are optional
//   			BidPrice: jsii.String("bidPrice"),
//   			BidPriceAsPercentageOfOnDemandPrice: jsii.Number(123),
//   			Configurations: []ConfigurationProperty{
//   				&ConfigurationProperty{
//   					Classification: jsii.String("classification"),
//   					Configurations: []ConfigurationProperty{
//   						configurationProperty_,
//   					},
//   					Properties: map[string]*string{
//   						"propertiesKey": jsii.String("properties"),
//   					},
//   				},
//   			},
//   			EbsConfiguration: &EbsConfigurationProperty{
//   				EbsBlockDeviceConfigs: []EbsBlockDeviceConfigProperty{
//   					&EbsBlockDeviceConfigProperty{
//   						VolumeSpecification: &VolumeSpecificationProperty{
//   							VolumeSize: size,
//   							VolumeType: awscdk.*Aws_stepfunctions_tasks.EmrCreateCluster.EbsBlockDeviceVolumeType_GP3,
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
//   		OnDemandSpecification: &OnDemandProvisioningSpecificationProperty{
//   			AllocationStrategy: awscdk.*Aws_stepfunctions_tasks.EmrCreateCluster.OnDemandAllocationStrategy_LOWEST_PRICE,
//   		},
//   		SpotSpecification: &SpotProvisioningSpecificationProperty{
//   			TimeoutAction: awscdk.*Aws_stepfunctions_tasks.EmrCreateCluster.SpotTimeoutAction_SWITCH_TO_ON_DEMAND,
//
//   			// the properties below are optional
//   			AllocationStrategy: awscdk.*Aws_stepfunctions_tasks.EmrCreateCluster.SpotAllocationStrategy_CAPACITY_OPTIMIZED,
//   			BlockDurationMinutes: jsii.Number(123),
//   			Timeout: cdk.Duration_Minutes(jsii.Number(30)),
//   			TimeoutDurationMinutes: jsii.Number(123),
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
	//
	// If not specified or set to 0, only Spot Instances are provisioned for the instance fleet using `targetSpotCapacity`.
	//
	// At least one of `targetSpotCapacity` and `targetOnDemandCapacity` should be greater than 0.
	// For a master instance fleet, only one of `targetSpotCapacity` and `targetOnDemandCapacity` can be specified, and its value
	// must be 1.
	// Default: No targetOnDemandCapacity.
	//
	TargetOnDemandCapacity *float64 `field:"optional" json:"targetOnDemandCapacity" yaml:"targetOnDemandCapacity"`
	// The target capacity of Spot units for the instance fleet, which determines how many Spot instances to provision.
	//
	// If not specified or set to 0, only On-Demand Instances are provisioned for the instance fleet using `targetOnDemandCapacity`.
	//
	// At least one of `targetSpotCapacity` and `targetOnDemandCapacity` should be greater than 0.
	// For a master instance fleet, only one of `targetSpotCapacity` and `targetOnDemandCapacity` can be specified, and its value
	// must be 1.
	// Default: No targetSpotCapacity.
	//
	TargetSpotCapacity *float64 `field:"optional" json:"targetSpotCapacity" yaml:"targetSpotCapacity"`
}

