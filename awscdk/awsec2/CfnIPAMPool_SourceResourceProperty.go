package awsec2


// The resource associated with this pool's space.
//
// Depending on the ResourceType, setting a SourceResource changes which space can be provisioned in this pool and which types of resources can receive allocations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceResourceProperty := &SourceResourceProperty{
//   	ResourceId: jsii.String("resourceId"),
//   	ResourceOwner: jsii.String("resourceOwner"),
//   	ResourceRegion: jsii.String("resourceRegion"),
//   	ResourceType: jsii.String("resourceType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ipampool-sourceresource.html
//
type CfnIPAMPool_SourceResourceProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ipampool-sourceresource.html#cfn-ec2-ipampool-sourceresource-resourceid
	//
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ipampool-sourceresource.html#cfn-ec2-ipampool-sourceresource-resourceowner
	//
	ResourceOwner *string `field:"required" json:"resourceOwner" yaml:"resourceOwner"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ipampool-sourceresource.html#cfn-ec2-ipampool-sourceresource-resourceregion
	//
	ResourceRegion *string `field:"required" json:"resourceRegion" yaml:"resourceRegion"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ipampool-sourceresource.html#cfn-ec2-ipampool-sourceresource-resourcetype
	//
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
}

