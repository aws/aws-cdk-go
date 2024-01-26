package awscodebuild


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   projectFleetProperty := &ProjectFleetProperty{
//   	FleetArn: jsii.String("fleetArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-projectfleet.html
//
type CfnProject_ProjectFleetProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-projectfleet.html#cfn-codebuild-project-projectfleet-fleetarn
	//
	FleetArn *string `field:"optional" json:"fleetArn" yaml:"fleetArn"`
}

