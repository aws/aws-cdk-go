package awsiotsitewise


// The data source configuration for a mount.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   mountSourceProperty := &MountSourceProperty{
//   	S3AccessPoint: &S3AccessPointSourceProperty{
//   		AccessPointArn: jsii.String("accessPointArn"),
//   		Prefix: jsii.String("prefix"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-task-mountsource.html
//
type CfnTaskPropsMixin_MountSourceProperty struct {
	// Configures a mount that reads from an Amazon S3 access point.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-task-mountsource.html#cfn-iotsitewise-task-mountsource-s3accesspoint
	//
	S3AccessPoint interface{} `field:"optional" json:"s3AccessPoint" yaml:"s3AccessPoint"`
}

