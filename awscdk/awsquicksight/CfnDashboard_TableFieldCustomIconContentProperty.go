package awsquicksight


// The custom icon content for the table link content configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tableFieldCustomIconContentProperty := &TableFieldCustomIconContentProperty{
//   	Icon: jsii.String("icon"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-tablefieldcustomiconcontent.html
//
type CfnDashboard_TableFieldCustomIconContentProperty struct {
	// The icon set type (link) of the custom icon content for table URL link content.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-tablefieldcustomiconcontent.html#cfn-quicksight-dashboard-tablefieldcustomiconcontent-icon
	//
	Icon *string `field:"optional" json:"icon" yaml:"icon"`
}

