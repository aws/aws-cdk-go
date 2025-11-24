package mixinsawsiotsitewise


// Identifies an AWS IoT SiteWise Monitor portal.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   portalProperty := &PortalProperty{
//   	Id: jsii.String("id"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-accesspolicy-portal.html
//
type CfnAccessPolicyPropsMixin_PortalProperty struct {
	// The ID of the portal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-accesspolicy-portal.html#cfn-iotsitewise-accesspolicy-portal-id
	//
	Id *string `field:"optional" json:"id" yaml:"id"`
}

