package awsappstream


// Properties for defining a `CfnApplicationFleetAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationFleetAssociationProps := &CfnApplicationFleetAssociationProps{
//   	ApplicationArn: jsii.String("applicationArn"),
//   	FleetName: jsii.String("fleetName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-applicationfleetassociation.html
//
type CfnApplicationFleetAssociationProps struct {
	// The ARN of the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-applicationfleetassociation.html#cfn-appstream-applicationfleetassociation-applicationarn
	//
	ApplicationArn *string `field:"required" json:"applicationArn" yaml:"applicationArn"`
	// The name of the fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-applicationfleetassociation.html#cfn-appstream-applicationfleetassociation-fleetname
	//
	FleetName *string `field:"required" json:"fleetName" yaml:"fleetName"`
}

