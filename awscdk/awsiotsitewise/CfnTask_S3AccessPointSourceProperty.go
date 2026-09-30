package awsiotsitewise


// Configures a mount that reads from an Amazon S3 access point.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3AccessPointSourceProperty := &S3AccessPointSourceProperty{
//   	AccessPointArn: jsii.String("accessPointArn"),
//
//   	// the properties below are optional
//   	Prefix: jsii.String("prefix"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-task-s3accesspointsource.html
//
type CfnTask_S3AccessPointSourceProperty struct {
	// The Amazon Resource Name (ARN) of the Amazon S3 access point.
	//
	// The mount reads objects from the bucket associated with this access point. Access is governed by the access point policy and the task execution role's IAM permissions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-task-s3accesspointsource.html#cfn-iotsitewise-task-s3accesspointsource-accesspointarn
	//
	AccessPointArn *string `field:"required" json:"accessPointArn" yaml:"accessPointArn"`
	// An object key name prefix.
	//
	// If specified, the mount includes only objects whose keys begin with this prefix. To include all objects at the access point, omit this field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-task-s3accesspointsource.html#cfn-iotsitewise-task-s3accesspointsource-prefix
	//
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

