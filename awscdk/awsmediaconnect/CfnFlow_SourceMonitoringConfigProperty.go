package awsmediaconnect


// The `SourceMonitoringConfig` property type specifies the source monitoring settings for an AWS::MediaConnect::Flow.
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
	// The current state of the thumbnail monitoring.
	//
	// - If you don't explicitly specify a value when creating a flow, no thumbnail state will be set.
	// - If you update an existing flow and remove a previously set thumbnail state, the value will change to `DISABLED` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-sourcemonitoringconfig.html#cfn-mediaconnect-flow-sourcemonitoringconfig-thumbnailstate
	//
	ThumbnailState *string `field:"required" json:"thumbnailState" yaml:"thumbnailState"`
}

