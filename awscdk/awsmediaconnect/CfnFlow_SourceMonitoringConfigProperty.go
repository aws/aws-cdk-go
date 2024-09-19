package awsmediaconnect


// The settings for source monitoring.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceMonitoringConfigProperty := &SourceMonitoringConfigProperty{
//   	ThumbnailState: jsii.String("thumbnailState"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-sourcemonitoringconfig.html
//
type CfnFlow_SourceMonitoringConfigProperty struct {
	// The state of thumbnail monitoring.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-sourcemonitoringconfig.html#cfn-mediaconnect-flow-sourcemonitoringconfig-thumbnailstate
	//
	ThumbnailState *string `field:"required" json:"thumbnailState" yaml:"thumbnailState"`
}

