package awsquicksight


// A QuickSight resource attached to the space.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   spaceResourceProperty := &SpaceResourceProperty{
//   	ResourceArn: jsii.String("resourceArn"),
//   	ResourceType: jsii.String("resourceType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-space-spaceresource.html
//
type CfnSpace_SpaceResourceProperty struct {
	// The ARN of the QuickSight resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-space-spaceresource.html#cfn-quicksight-space-spaceresource-resourcearn
	//
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
	// The type of QuickSight resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-space-spaceresource.html#cfn-quicksight-space-spaceresource-resourcetype
	//
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
}

