package awsevs


// An object that represents a host.
//
// > You cannot use `dedicatedHostId` and `placementGroupId` together in the same `HostInfoForCreate` object. This results in a `ValidationException` response.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hostInfoForCreateProperty := &HostInfoForCreateProperty{
//   	HostName: jsii.String("hostName"),
//   	InstanceType: jsii.String("instanceType"),
//   	KeyName: jsii.String("keyName"),
//
//   	// the properties below are optional
//   	DedicatedHostId: jsii.String("dedicatedHostId"),
//   	PlacementGroupId: jsii.String("placementGroupId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evs-environment-hostinfoforcreate.html
//
type CfnEnvironment_HostInfoForCreateProperty struct {
	// The DNS hostname of the host.
	//
	// DNS hostnames for hosts must be unique across Amazon EVS environments and within VCF.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evs-environment-hostinfoforcreate.html#cfn-evs-environment-hostinfoforcreate-hostname
	//
	HostName *string `field:"required" json:"hostName" yaml:"hostName"`
	// The EC2 instance type that represents the host.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evs-environment-hostinfoforcreate.html#cfn-evs-environment-hostinfoforcreate-instancetype
	//
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// The name of the SSH key that is used to access the host.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evs-environment-hostinfoforcreate.html#cfn-evs-environment-hostinfoforcreate-keyname
	//
	KeyName *string `field:"required" json:"keyName" yaml:"keyName"`
	// The unique ID of the Amazon EC2 Dedicated Host.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evs-environment-hostinfoforcreate.html#cfn-evs-environment-hostinfoforcreate-dedicatedhostid
	//
	DedicatedHostId *string `field:"optional" json:"dedicatedHostId" yaml:"dedicatedHostId"`
	// The unique ID of the placement group where the host is placed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evs-environment-hostinfoforcreate.html#cfn-evs-environment-hostinfoforcreate-placementgroupid
	//
	PlacementGroupId *string `field:"optional" json:"placementGroupId" yaml:"placementGroupId"`
}

