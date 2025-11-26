package previewawsquicksightmixins


// The custom icon content for the table link content configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   tableFieldCustomIconContentProperty := &TableFieldCustomIconContentProperty{
//   	Icon: jsii.String("icon"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-tablefieldcustomiconcontent.html
//
type CfnTemplatePropsMixin_TableFieldCustomIconContentProperty struct {
	// The icon set type (link) of the custom icon content for table URL link content.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-tablefieldcustomiconcontent.html#cfn-quicksight-template-tablefieldcustomiconcontent-icon
	//
	Icon *string `field:"optional" json:"icon" yaml:"icon"`
}

