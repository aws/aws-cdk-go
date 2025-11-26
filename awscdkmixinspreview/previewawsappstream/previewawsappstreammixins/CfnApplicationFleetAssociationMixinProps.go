package previewawsappstreammixins


// Properties for CfnApplicationFleetAssociationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnApplicationFleetAssociationMixinProps := &CfnApplicationFleetAssociationMixinProps{
//   	ApplicationArn: jsii.String("applicationArn"),
//   	FleetName: jsii.String("fleetName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-applicationfleetassociation.html
//
type CfnApplicationFleetAssociationMixinProps struct {
	// The ARN of the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-applicationfleetassociation.html#cfn-appstream-applicationfleetassociation-applicationarn
	//
	ApplicationArn *string `field:"optional" json:"applicationArn" yaml:"applicationArn"`
	// The name of the fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-applicationfleetassociation.html#cfn-appstream-applicationfleetassociation-fleetname
	//
	FleetName *string `field:"optional" json:"fleetName" yaml:"fleetName"`
}

