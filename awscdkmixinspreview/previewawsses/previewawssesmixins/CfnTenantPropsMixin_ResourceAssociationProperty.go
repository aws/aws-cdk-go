package previewawssesmixins


// The resource to associate with the tenant.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   resourceAssociationProperty := &ResourceAssociationProperty{
//   	ResourceArn: jsii.String("resourceArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-tenant-resourceassociation.html
//
type CfnTenantPropsMixin_ResourceAssociationProperty struct {
	// The Amazon Resource Name (ARN) of the resource associated with the tenant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-tenant-resourceassociation.html#cfn-ses-tenant-resourceassociation-resourcearn
	//
	ResourceArn *string `field:"optional" json:"resourceArn" yaml:"resourceArn"`
}

