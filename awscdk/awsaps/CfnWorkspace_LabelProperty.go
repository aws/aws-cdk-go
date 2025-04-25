package awsaps


// Series label.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   labelProperty := &LabelProperty{
//   	Name: jsii.String("name"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-workspace-label.html
//
type CfnWorkspace_LabelProperty struct {
	// Name of the label.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-workspace-label.html#cfn-aps-workspace-label-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Value of the label.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-workspace-label.html#cfn-aps-workspace-label-value
	//
	Value *string `field:"required" json:"value" yaml:"value"`
}

