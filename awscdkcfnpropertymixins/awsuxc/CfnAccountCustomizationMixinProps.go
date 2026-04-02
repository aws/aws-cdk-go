package awsuxc


// Properties for CfnAccountCustomizationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnAccountCustomizationMixinProps := &CfnAccountCustomizationMixinProps{
//   	AccountColor: jsii.String("accountColor"),
//   	VisibleRegions: []*string{
//   		jsii.String("visibleRegions"),
//   	},
//   	VisibleServices: []*string{
//   		jsii.String("visibleServices"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-uxc-accountcustomization.html
//
type CfnAccountCustomizationMixinProps struct {
	// The color theme assigned to the account for visual identification in the AWS Console.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-uxc-accountcustomization.html#cfn-uxc-accountcustomization-accountcolor
	//
	AccountColor *string `field:"optional" json:"accountColor" yaml:"accountColor"`
	// A list of AWS region identifiers visible to the account in the AWS Console.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-uxc-accountcustomization.html#cfn-uxc-accountcustomization-visibleregions
	//
	VisibleRegions *[]*string `field:"optional" json:"visibleRegions" yaml:"visibleRegions"`
	// A list of AWS service identifiers visible to the account in the AWS Console.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-uxc-accountcustomization.html#cfn-uxc-accountcustomization-visibleservices
	//
	VisibleServices *[]*string `field:"optional" json:"visibleServices" yaml:"visibleServices"`
}

