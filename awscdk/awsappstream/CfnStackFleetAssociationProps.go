package awsappstream


// Properties for defining a `CfnStackFleetAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnStackFleetAssociationProps := &CfnStackFleetAssociationProps{
//   	FleetName: jsii.String("fleetName"),
//   	StackName: jsii.String("stackName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-stackfleetassociation.html
//
type CfnStackFleetAssociationProps struct {
	// The name of the fleet.
	//
	// To associate a fleet with a stack, you must specify a dependency on the fleet resource. For more information, see [DependsOn Attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-stackfleetassociation.html#cfn-appstream-stackfleetassociation-fleetname
	//
	FleetName *string `field:"required" json:"fleetName" yaml:"fleetName"`
	// The name of the stack.
	//
	// To associate a fleet with a stack, you must specify a dependency on the stack resource. For more information, see [DependsOn Attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-stackfleetassociation.html#cfn-appstream-stackfleetassociation-stackname
	//
	StackName *string `field:"required" json:"stackName" yaml:"stackName"`
}

