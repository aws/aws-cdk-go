package awsresiliencehubv2


// Properties for defining a `CfnUserJourney`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnUserJourneyProps := &CfnUserJourneyProps{
//   	Name: jsii.String("name"),
//   	SystemIdentifier: jsii.String("systemIdentifier"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	PolicyArn: jsii.String("policyArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehubv2-userjourney.html
//
type CfnUserJourneyProps struct {
	// The name of the user journey.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehubv2-userjourney.html#cfn-resiliencehubv2-userjourney-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The system ARN or system ID that owns this user journey.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehubv2-userjourney.html#cfn-resiliencehubv2-userjourney-systemidentifier
	//
	SystemIdentifier *string `field:"required" json:"systemIdentifier" yaml:"systemIdentifier"`
	// The description of the user journey.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehubv2-userjourney.html#cfn-resiliencehubv2-userjourney-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The ARN of the resilience policy to associate with this user journey.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehubv2-userjourney.html#cfn-resiliencehubv2-userjourney-policyarn
	//
	PolicyArn *string `field:"optional" json:"policyArn" yaml:"policyArn"`
}

