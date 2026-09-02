package awsssm


// An Azure subscription with its ID and optional display name.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   azureSubscriptionProperty := &AzureSubscriptionProperty{
//   	Id: jsii.String("id"),
//
//   	// the properties below are optional
//   	DisplayName: jsii.String("displayName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-cloudconnector-azuresubscription.html
//
type CfnCloudConnector_AzureSubscriptionProperty struct {
	// The Azure subscription ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-cloudconnector-azuresubscription.html#cfn-ssm-cloudconnector-azuresubscription-id
	//
	Id *string `field:"required" json:"id" yaml:"id"`
	// The display name of the Azure subscription.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-cloudconnector-azuresubscription.html#cfn-ssm-cloudconnector-azuresubscription-displayname
	//
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
}

