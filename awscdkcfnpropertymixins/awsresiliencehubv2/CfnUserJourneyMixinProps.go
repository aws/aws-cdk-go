package awsresiliencehubv2


// Properties for CfnUserJourneyPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnUserJourneyMixinProps := &CfnUserJourneyMixinProps{
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	PolicyArn: jsii.String("policyArn"),
//   	SystemIdentifier: jsii.String("systemIdentifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehubv2-userjourney.html
//
type CfnUserJourneyMixinProps struct {
	// The description of the user journey.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehubv2-userjourney.html#cfn-resiliencehubv2-userjourney-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the user journey.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehubv2-userjourney.html#cfn-resiliencehubv2-userjourney-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The ARN of the resilience policy to associate with this user journey.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehubv2-userjourney.html#cfn-resiliencehubv2-userjourney-policyarn
	//
	PolicyArn *string `field:"optional" json:"policyArn" yaml:"policyArn"`
	// The system ARN or system ID that owns this user journey.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehubv2-userjourney.html#cfn-resiliencehubv2-userjourney-systemidentifier
	//
	SystemIdentifier *string `field:"optional" json:"systemIdentifier" yaml:"systemIdentifier"`
}

