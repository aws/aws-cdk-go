package awsec2


// Describes an potential intermediate component of a feasible path.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   alternatePathHintProperty := &AlternatePathHintProperty{
//   	ComponentArn: jsii.String("componentArn"),
//   	ComponentId: jsii.String("componentId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-alternatepathhint.html
//
type CfnNetworkInsightsAnalysis_AlternatePathHintProperty struct {
	// The Amazon Resource Name (ARN) of the component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-alternatepathhint.html#cfn-ec2-networkinsightsanalysis-alternatepathhint-componentarn
	//
	ComponentArn *string `field:"optional" json:"componentArn" yaml:"componentArn"`
	// The ID of the component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-alternatepathhint.html#cfn-ec2-networkinsightsanalysis-alternatepathhint-componentid
	//
	ComponentId *string `field:"optional" json:"componentId" yaml:"componentId"`
}

