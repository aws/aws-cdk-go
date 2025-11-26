package previewawsec2mixins


// Properties for CfnCapacityReservationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCapacityReservationMixinProps := &CfnCapacityReservationMixinProps{
//   	AvailabilityZone: jsii.String("availabilityZone"),
//   	AvailabilityZoneId: jsii.String("availabilityZoneId"),
//   	EbsOptimized: jsii.Boolean(false),
//   	EndDate: jsii.String("endDate"),
//   	EndDateType: jsii.String("endDateType"),
//   	EphemeralStorage: jsii.Boolean(false),
//   	InstanceCount: jsii.Number(123),
//   	InstanceMatchCriteria: jsii.String("instanceMatchCriteria"),
//   	InstancePlatform: jsii.String("instancePlatform"),
//   	InstanceType: jsii.String("instanceType"),
//   	OutPostArn: jsii.String("outPostArn"),
//   	PlacementGroupArn: jsii.String("placementGroupArn"),
//   	TagSpecifications: []TagSpecificationProperty{
//   		&TagSpecificationProperty{
//   			ResourceType: jsii.String("resourceType"),
//   			Tags: []CfnTag{
//   				&CfnTag{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	Tenancy: jsii.String("tenancy"),
//   	UnusedReservationBillingOwnerId: jsii.String("unusedReservationBillingOwnerId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-capacityreservation.html
//
type CfnCapacityReservationMixinProps struct {
	// The Availability Zone in which to create the Capacity Reservation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-capacityreservation.html#cfn-ec2-capacityreservation-availabilityzone
	//
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The ID of the Availability Zone in which the capacity is reserved.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-capacityreservation.html#cfn-ec2-capacityreservation-availabilityzoneid
	//
	AvailabilityZoneId *string `field:"optional" json:"availabilityZoneId" yaml:"availabilityZoneId"`
	// Indicates whether the Capacity Reservation supports EBS-optimized instances.
	//
	// This optimization provides dedicated throughput to Amazon EBS and an optimized configuration stack to provide optimal I/O performance. This optimization isn't available with all instance types. Additional usage charges apply when using an EBS- optimized instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-capacityreservation.html#cfn-ec2-capacityreservation-ebsoptimized
	//
	EbsOptimized interface{} `field:"optional" json:"ebsOptimized" yaml:"ebsOptimized"`
	// The date and time at which the Capacity Reservation expires.
	//
	// When a Capacity Reservation expires, the reserved capacity is released and you can no longer launch instances into it. The Capacity Reservation's state changes to `expired` when it reaches its end date and time.
	//
	// You must provide an `EndDate` value if `EndDateType` is `limited` . Omit `EndDate` if `EndDateType` is `unlimited` .
	//
	// If the `EndDateType` is `limited` , the Capacity Reservation is cancelled within an hour from the specified time. For example, if you specify 5/31/2019, 13:30:55, the Capacity Reservation is guaranteed to end between 13:30:55 and 14:30:55 on 5/31/2019.
	//
	// If you are requesting a future-dated Capacity Reservation, you can't specify an end date and time that is within the commitment duration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-capacityreservation.html#cfn-ec2-capacityreservation-enddate
	//
	EndDate *string `field:"optional" json:"endDate" yaml:"endDate"`
	// Indicates the way in which the Capacity Reservation ends.
	//
	// A Capacity Reservation can have one of the following end types:
	//
	// - `unlimited` - The Capacity Reservation remains active until you explicitly cancel it. Do not provide an `EndDate` if the `EndDateType` is `unlimited` .
	// - `limited` - The Capacity Reservation expires automatically at a specified date and time. You must provide an `EndDate` value if the `EndDateType` value is `limited` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-capacityreservation.html#cfn-ec2-capacityreservation-enddatetype
	//
	EndDateType *string `field:"optional" json:"endDateType" yaml:"endDateType"`
	// *Deprecated.*.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-capacityreservation.html#cfn-ec2-capacityreservation-ephemeralstorage
	//
	EphemeralStorage interface{} `field:"optional" json:"ephemeralStorage" yaml:"ephemeralStorage"`
	// The number of instances for which to reserve capacity.
	//
	// > You can request future-dated Capacity Reservations for an instance count with a minimum of 64 vCPUs. For example, if you request a future-dated Capacity Reservation for `m5.xlarge` instances, you must request at least 25 instances ( *16 * m5.xlarge = 64 vCPUs* ).
	//
	// Valid range: 1 - 1000.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-capacityreservation.html#cfn-ec2-capacityreservation-instancecount
	//
	InstanceCount *float64 `field:"optional" json:"instanceCount" yaml:"instanceCount"`
	// Indicates the type of instance launches that the Capacity Reservation accepts. The options include:.
	//
	// - `open` - The Capacity Reservation automatically matches all instances that have matching attributes (instance type, platform, and Availability Zone). Instances that have matching attributes run in the Capacity Reservation automatically without specifying any additional parameters.
	// - `targeted` - The Capacity Reservation only accepts instances that have matching attributes (instance type, platform, and Availability Zone), and explicitly target the Capacity Reservation. This ensures that only permitted instances can use the reserved capacity.
	//
	// > If you are requesting a future-dated Capacity Reservation, you must specify `targeted` .
	//
	// Default: `open`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-capacityreservation.html#cfn-ec2-capacityreservation-instancematchcriteria
	//
	InstanceMatchCriteria *string `field:"optional" json:"instanceMatchCriteria" yaml:"instanceMatchCriteria"`
	// The type of operating system for which to reserve capacity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-capacityreservation.html#cfn-ec2-capacityreservation-instanceplatform
	//
	InstancePlatform *string `field:"optional" json:"instancePlatform" yaml:"instancePlatform"`
	// The instance type for which to reserve capacity.
	//
	// > You can request future-dated Capacity Reservations for instance types in the C, M, R, I, T, and G instance families only.
	//
	// For more information, see [Instance types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html) in the *Amazon EC2 User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-capacityreservation.html#cfn-ec2-capacityreservation-instancetype
	//
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// > Not supported for future-dated Capacity Reservations.
	//
	// The Amazon Resource Name (ARN) of the Outpost on which to create the Capacity Reservation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-capacityreservation.html#cfn-ec2-capacityreservation-outpostarn
	//
	OutPostArn *string `field:"optional" json:"outPostArn" yaml:"outPostArn"`
	// > Not supported for future-dated Capacity Reservations.
	//
	// The Amazon Resource Name (ARN) of the cluster placement group in which to create the Capacity Reservation. For more information, see [Capacity Reservations for cluster placement groups](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/cr-cpg.html) in the *Amazon EC2 User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-capacityreservation.html#cfn-ec2-capacityreservation-placementgrouparn
	//
	PlacementGroupArn *string `field:"optional" json:"placementGroupArn" yaml:"placementGroupArn"`
	// The tags to apply to the Capacity Reservation during launch.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-capacityreservation.html#cfn-ec2-capacityreservation-tagspecifications
	//
	TagSpecifications *[]*CfnCapacityReservationPropsMixin_TagSpecificationProperty `field:"optional" json:"tagSpecifications" yaml:"tagSpecifications"`
	// Indicates the tenancy of the Capacity Reservation. A Capacity Reservation can have one of the following tenancy settings:.
	//
	// - `default` - The Capacity Reservation is created on hardware that is shared with other AWS accounts .
	// - `dedicated` - The Capacity Reservation is created on single-tenant hardware that is dedicated to a single AWS account .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-capacityreservation.html#cfn-ec2-capacityreservation-tenancy
	//
	Tenancy *string `field:"optional" json:"tenancy" yaml:"tenancy"`
	// The ID of the AWS account to which to assign billing of the unused capacity of the Capacity Reservation.
	//
	// A request will be sent to the specified account. That account must accept the request for the billing to be assigned to their account. For more information, see [Billing assignment for shared Amazon EC2 Capacity Reservations](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/assign-billing.html) .
	//
	// You can assign billing only for shared Capacity Reservations. To share a Capacity Reservation, you must add it to a resource share. For more information, see [AWS::RAM::ResourceShare](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ram-resourceshare.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-capacityreservation.html#cfn-ec2-capacityreservation-unusedreservationbillingownerid
	//
	UnusedReservationBillingOwnerId *string `field:"optional" json:"unusedReservationBillingOwnerId" yaml:"unusedReservationBillingOwnerId"`
}

