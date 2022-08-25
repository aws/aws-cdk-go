package awsbatch


// The authorization configuration details for the Amazon EFS file system.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   authorizationConfigProperty := &authorizationConfigProperty{
//   	accessPointId: jsii.String("accessPointId"),
//   	iam: jsii.String("iam"),
//   }
//
type CfnJobDefinition_AuthorizationConfigProperty struct {
	// The Amazon EFS access point ID to use.
	//
	// If an access point is specified, the root directory value specified in the `EFSVolumeConfiguration` must either be omitted or set to `/` which will enforce the path set on the EFS access point. If an access point is used, transit encryption must be enabled in the `EFSVolumeConfiguration` . For more information, see [Working with Amazon EFS access points](https://docs.aws.amazon.com/efs/latest/ug/efs-access-points.html) in the *Amazon Elastic File System User Guide* .
	AccessPointId *string `field:"optional" json:"accessPointId" yaml:"accessPointId"`
	// Whether or not to use the AWS Batch job IAM role defined in a job definition when mounting the Amazon EFS file system.
	//
	// If enabled, transit encryption must be enabled in the `EFSVolumeConfiguration` . If this parameter is omitted, the default value of `DISABLED` is used. For more information, see [Using Amazon EFS access points](https://docs.aws.amazon.com/batch/latest/userguide/efs-volumes.html#efs-volume-accesspoints) in the *AWS Batch User Guide* . EFS IAM authorization requires that `TransitEncryption` be `ENABLED` and that a `JobRoleArn` is specified.
	Iam *string `field:"optional" json:"iam" yaml:"iam"`
}

