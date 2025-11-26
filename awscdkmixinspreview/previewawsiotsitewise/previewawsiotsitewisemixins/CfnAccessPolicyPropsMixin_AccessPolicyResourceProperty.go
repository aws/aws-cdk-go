package previewawsiotsitewisemixins


// The AWS IoT SiteWise Monitor resource for this access policy.
//
// Choose either a portal or a project.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   accessPolicyResourceProperty := &AccessPolicyResourceProperty{
//   	Portal: &PortalProperty{
//   		Id: jsii.String("id"),
//   	},
//   	Project: &ProjectProperty{
//   		Id: jsii.String("id"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-accesspolicy-accesspolicyresource.html
//
type CfnAccessPolicyPropsMixin_AccessPolicyResourceProperty struct {
	// Identifies an AWS IoT SiteWise Monitor portal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-accesspolicy-accesspolicyresource.html#cfn-iotsitewise-accesspolicy-accesspolicyresource-portal
	//
	Portal interface{} `field:"optional" json:"portal" yaml:"portal"`
	// Identifies a specific AWS IoT SiteWise Monitor project.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-accesspolicy-accesspolicyresource.html#cfn-iotsitewise-accesspolicy-accesspolicyresource-project
	//
	Project interface{} `field:"optional" json:"project" yaml:"project"`
}

