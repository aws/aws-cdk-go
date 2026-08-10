package awsmediapackage


// Properties for defining a `CfnHarvestJob`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnHarvestJobProps := &CfnHarvestJobProps{
//   	EndTime: jsii.String("endTime"),
//   	Id: jsii.String("id"),
//   	OriginEndpointId: jsii.String("originEndpointId"),
//   	S3Destination: &S3DestinationProperty{
//   		BucketName: jsii.String("bucketName"),
//   		ManifestKey: jsii.String("manifestKey"),
//   		RoleArn: jsii.String("roleArn"),
//   	},
//   	StartTime: jsii.String("startTime"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-harvestjob.html
//
type CfnHarvestJobProps struct {
	// The end of the time-window which will be harvested.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-harvestjob.html#cfn-mediapackage-harvestjob-endtime
	//
	EndTime *string `field:"required" json:"endTime" yaml:"endTime"`
	// The ID of the HarvestJob.
	//
	// The ID must be unique within the region and it cannot be changed after the HarvestJob is submitted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-harvestjob.html#cfn-mediapackage-harvestjob-id
	//
	Id *string `field:"required" json:"id" yaml:"id"`
	// The ID of the OriginEndpoint that the HarvestJob will harvest from.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-harvestjob.html#cfn-mediapackage-harvestjob-originendpointid
	//
	OriginEndpointId *string `field:"required" json:"originEndpointId" yaml:"originEndpointId"`
	// Configuration parameters for where in an S3 bucket to place the harvested content.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-harvestjob.html#cfn-mediapackage-harvestjob-s3destination
	//
	S3Destination interface{} `field:"required" json:"s3Destination" yaml:"s3Destination"`
	// The start of the time-window which will be harvested.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-harvestjob.html#cfn-mediapackage-harvestjob-starttime
	//
	StartTime *string `field:"required" json:"startTime" yaml:"startTime"`
}

