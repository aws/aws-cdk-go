package awssagemaker


// Configuration settings for the SageMaker Partner AI App.
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
	// The list of users that are given admin access to the SageMaker Partner AI App.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-partnerapp-partnerappconfig.html#cfn-sagemaker-partnerapp-partnerappconfig-adminusers
	//
	AdminUsers *[]*string `field:"optional" json:"adminUsers" yaml:"adminUsers"`
	// This is a map of required inputs for a SageMaker Partner AI App.
	//
	// Based on the application type, the map is populated with a key and value pair that is specific to the user and application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-partnerapp-partnerappconfig.html#cfn-sagemaker-partnerapp-partnerappconfig-arguments
	//
	Arguments interface{} `field:"optional" json:"arguments" yaml:"arguments"`
}

