package previewawsarcregionswitchmixins


// Configuration for handling failures when performing operations on Aurora global databases.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   globalAuroraUngracefulProperty := &GlobalAuroraUngracefulProperty{
//   	Ungraceful: jsii.String("ungraceful"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-globalauroraungraceful.html
//
type CfnPlanPropsMixin_GlobalAuroraUngracefulProperty struct {
	// The settings for ungraceful execution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-globalauroraungraceful.html#cfn-arcregionswitch-plan-globalauroraungraceful-ungraceful
	//
	Ungraceful *string `field:"optional" json:"ungraceful" yaml:"ungraceful"`
}

