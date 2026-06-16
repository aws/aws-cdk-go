package awsresiliencehubv2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   associatedSystemProperty := &AssociatedSystemProperty{
//   	SystemArn: jsii.String("systemArn"),
//
//   	// the properties below are optional
//   	UserJourneyIds: []*string{
//   		jsii.String("userJourneyIds"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-service-associatedsystem.html
//
type CfnService_AssociatedSystemProperty struct {
	// The system ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-service-associatedsystem.html#cfn-resiliencehubv2-service-associatedsystem-systemarn
	//
	SystemArn *string `field:"required" json:"systemArn" yaml:"systemArn"`
	// User journey IDs associated with this system.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-service-associatedsystem.html#cfn-resiliencehubv2-service-associatedsystem-userjourneyids
	//
	UserJourneyIds *[]*string `field:"optional" json:"userJourneyIds" yaml:"userJourneyIds"`
}

