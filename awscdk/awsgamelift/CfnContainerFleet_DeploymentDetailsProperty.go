package awsgamelift


// Information about the most recent deployment for the container fleet.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deploymentDetailsProperty := &DeploymentDetailsProperty{
//   	LatestDeploymentId: jsii.String("latestDeploymentId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containerfleet-deploymentdetails.html
//
type CfnContainerFleet_DeploymentDetailsProperty struct {
	// A unique identifier for a fleet deployment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containerfleet-deploymentdetails.html#cfn-gamelift-containerfleet-deploymentdetails-latestdeploymentid
	//
	LatestDeploymentId *string `field:"optional" json:"latestDeploymentId" yaml:"latestDeploymentId"`
}

