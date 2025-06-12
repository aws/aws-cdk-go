package awsecs


// The authorization configuration details for the Amazon EFS file system.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   authorizationConfigProperty := &AuthorizationConfigProperty{
//   	AccessPointId: jsii.String("accessPointId"),
//   	Iam: jsii.String("iam"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-authorizationconfig.html
//
type CfnTaskDefinition_AuthorizationConfigProperty struct {
	// The Amazon EFS access point ID to use.
	//
	// If an access point is specified, the root directory value specified in the `EFSVolumeConfiguration` must either be omitted or set to `/` which will enforce the path set on the EFS access point. If an access point is used, transit encryption must be on in the `EFSVolumeConfiguration` . For more information, see [Working with Amazon EFS access points](https://docs.aws.amazon.com/efs/latest/ug/efs-access-points.html) in the *Amazon Elastic File System User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-authorizationconfig.html#cfn-ecs-taskdefinition-authorizationconfig-accesspointid
	//
	AccessPointId *string `field:"optional" json:"accessPointId" yaml:"accessPointId"`
	// Determines whether to use the Amazon ECS task role defined in a task definition when mounting the Amazon EFS file system.
	//
	// If it is turned on, transit encryption must be turned on in the `EFSVolumeConfiguration` . If this parameter is omitted, the default value of `DISABLED` is used. For more information, see [Using Amazon EFS access points](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/efs-volumes.html#efs-volume-accesspoints) in the *Amazon Elastic Container Service Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-authorizationconfig.html#cfn-ecs-taskdefinition-authorizationconfig-iam
	//
	Iam *string `field:"optional" json:"iam" yaml:"iam"`
}

