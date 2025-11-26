package previewawsworkspacesinstancesmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   capacityReservationTargetProperty := &CapacityReservationTargetProperty{
//   	CapacityReservationId: jsii.String("capacityReservationId"),
//   	CapacityReservationResourceGroupArn: jsii.String("capacityReservationResourceGroupArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesinstances-workspaceinstance-capacityreservationtarget.html
//
type CfnWorkspaceInstancePropsMixin_CapacityReservationTargetProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesinstances-workspaceinstance-capacityreservationtarget.html#cfn-workspacesinstances-workspaceinstance-capacityreservationtarget-capacityreservationid
	//
	CapacityReservationId *string `field:"optional" json:"capacityReservationId" yaml:"capacityReservationId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesinstances-workspaceinstance-capacityreservationtarget.html#cfn-workspacesinstances-workspaceinstance-capacityreservationtarget-capacityreservationresourcegrouparn
	//
	CapacityReservationResourceGroupArn *string `field:"optional" json:"capacityReservationResourceGroupArn" yaml:"capacityReservationResourceGroupArn"`
}

