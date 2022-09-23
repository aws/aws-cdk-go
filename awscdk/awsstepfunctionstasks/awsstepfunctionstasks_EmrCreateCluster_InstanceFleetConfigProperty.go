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
//   instanceFleetConfigProperty := &instanceFleetConfigProperty{
//   	instanceFleetType: awscdk.Aws_stepfunctions_tasks.emrCreateCluster.instanceRoleType_MASTER,
//
//   	// the properties below are optional
//   	instanceTypeConfigs: []instanceTypeConfigProperty{
//   		&instanceTypeConfigProperty{
//   			instanceType: jsii.String("instanceType"),
//
//   			// the properties below are optional
//   			bidPrice: jsii.String("bidPrice"),
//   			bidPriceAsPercentageOfOnDemandPrice: jsii.Number(123),
//   			configurations: []*configurationProperty{
//   				&configurationProperty{
//   					classification: jsii.String("classification"),
//   					configurations: []*configurationProperty{
//   						configurationProperty_,
//   					},
//   					properties: map[string]*string{
//   						"propertiesKey": jsii.String("properties"),
//   					},
//   				},
//   			},
//   			ebsConfiguration: &ebsConfigurationProperty{
//   				ebsBlockDeviceConfigs: []ebsBlockDeviceConfigProperty{
//   					&ebsBlockDeviceConfigProperty{
//   						volumeSpecification: &volumeSpecificationProperty{
//   							volumeSize: size,
//   							volumeType: awscdk.*Aws_stepfunctions_tasks.*emrCreateCluster.ebsBlockDeviceVolumeType_GP2,
//
//   							// the properties below are optional
//   							iops: jsii.Number(123),
//   						},
//
//   						// the properties below are optional
//   						volumesPerInstance: jsii.Number(123),
//   					},
//   				},
//   				ebsOptimized: jsii.Boolean(false),
//   			},
//   			weightedCapacity: jsii.Number(123),
//   		},
//   	},
//   	launchSpecifications: &instanceFleetProvisioningSpecificationsProperty{
//   		spotSpecification: &spotProvisioningSpecificationProperty{
//   			timeoutAction: awscdk.*Aws_stepfunctions_tasks.*emrCreateCluster.spotTimeoutAction_SWITCH_TO_ON_DEMAND,
//   			timeoutDurationMinutes: jsii.Number(123),
//
//   			// the properties below are optional
//   			allocationStrategy: awscdk.*Aws_stepfunctions_tasks.*emrCreateCluster.spotAllocationStrategy_CAPACITY_OPTIMIZED,
//   			blockDurationMinutes: jsii.Number(123),
//   		},
//   	},
//   	name: jsii.String("name"),
//   	targetOnDemandCapacity: jsii.Number(123),
//   	targetSpotCapacity: jsii.Number(123),
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
	InstanceTypeConfigs *[]*EmrCreateCluster_InstanceTypeConfigProperty `field:"optional" json:"instanceTypeConfigs" yaml:"instanceTypeConfigs"`
	// The launch specification for the instance fleet.
	LaunchSpecifications *EmrCreateCluster_InstanceFleetProvisioningSpecificationsProperty `field:"optional" json:"launchSpecifications" yaml:"launchSpecifications"`
	// The friendly name of the instance fleet.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The target capacity of On-Demand units for the instance fleet, which determines how many On-Demand instances to provision.
	TargetOnDemandCapacity *float64 `field:"optional" json:"targetOnDemandCapacity" yaml:"targetOnDemandCapacity"`
	// The target capacity of Spot units for the instance fleet, which determines how many Spot instances to provision.
	TargetSpotCapacity *float64 `field:"optional" json:"targetSpotCapacity" yaml:"targetSpotCapacity"`
}

