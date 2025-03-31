package awsmediaconnect


// The priority you want to assign to a source.
//
// You can have a primary stream and a backup stream or two equally prioritized streams.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourcePriorityProperty := &SourcePriorityProperty{
//   	PrimarySource: jsii.String("primarySource"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-bridge-sourcepriority.html
//
type CfnBridge_SourcePriorityProperty struct {
	// The name of the source you choose as the primary source for this flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-bridge-sourcepriority.html#cfn-mediaconnect-bridge-sourcepriority-primarysource
	//
	PrimarySource *string `field:"optional" json:"primarySource" yaml:"primarySource"`
}

