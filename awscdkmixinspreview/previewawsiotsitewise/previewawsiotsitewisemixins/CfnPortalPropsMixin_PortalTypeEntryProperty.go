package previewawsiotsitewisemixins


// Container associated a certain PortalType.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   portalTypeEntryProperty := &PortalTypeEntryProperty{
//   	PortalTools: []*string{
//   		jsii.String("portalTools"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-portal-portaltypeentry.html
//
type CfnPortalPropsMixin_PortalTypeEntryProperty struct {
	// The array of tools associated with the specified portal type.
	//
	// The possible values are `ASSISTANT` and `DASHBOARD` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-portal-portaltypeentry.html#cfn-iotsitewise-portal-portaltypeentry-portaltools
	//
	PortalTools *[]*string `field:"optional" json:"portalTools" yaml:"portalTools"`
}

