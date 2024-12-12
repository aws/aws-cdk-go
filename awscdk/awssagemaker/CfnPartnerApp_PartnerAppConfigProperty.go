package awssagemaker


// A collection of configuration settings for the PartnerApp.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   partnerAppConfigProperty := &PartnerAppConfigProperty{
//   	AdminUsers: []*string{
//   		jsii.String("adminUsers"),
//   	},
//   	Arguments: map[string]*string{
//   		"argumentsKey": jsii.String("arguments"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-partnerapp-partnerappconfig.html
//
type CfnPartnerApp_PartnerAppConfigProperty struct {
	// A list of users with administrator privileges for the PartnerApp.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-partnerapp-partnerappconfig.html#cfn-sagemaker-partnerapp-partnerappconfig-adminusers
	//
	AdminUsers *[]*string `field:"optional" json:"adminUsers" yaml:"adminUsers"`
	// A list of arguments to pass to the PartnerApp.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-partnerapp-partnerappconfig.html#cfn-sagemaker-partnerapp-partnerappconfig-arguments
	//
	Arguments interface{} `field:"optional" json:"arguments" yaml:"arguments"`
}

