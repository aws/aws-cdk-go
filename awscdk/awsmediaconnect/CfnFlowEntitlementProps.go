package awsmediaconnect


// Properties for defining a `CfnFlowEntitlement`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFlowEntitlementProps := &CfnFlowEntitlementProps{
//   	Description: jsii.String("description"),
//   	FlowArn: jsii.String("flowArn"),
//   	Name: jsii.String("name"),
//   	Subscribers: []*string{
//   		jsii.String("subscribers"),
//   	},
//
//   	// the properties below are optional
//   	DataTransferSubscriberFeePercent: jsii.Number(123),
//   	Encryption: &EncryptionProperty{
//   		Algorithm: jsii.String("algorithm"),
//   		RoleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
//   		ConstantInitializationVector: jsii.String("constantInitializationVector"),
//   		DeviceId: jsii.String("deviceId"),
//   		KeyType: jsii.String("keyType"),
//   		Region: jsii.String("region"),
//   		ResourceId: jsii.String("resourceId"),
//   		SecretArn: jsii.String("secretArn"),
//   		Url: jsii.String("url"),
//   	},
//   	EntitlementStatus: jsii.String("entitlementStatus"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowentitlement.html
//
type CfnFlowEntitlementProps struct {
	// A description of the entitlement.
	//
	// This description appears only on the MediaConnect console and is not visible outside of the current AWS account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowentitlement.html#cfn-mediaconnect-flowentitlement-description
	//
	Description *string `field:"required" json:"description" yaml:"description"`
	// The Amazon Resource Name (ARN) of the flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowentitlement.html#cfn-mediaconnect-flowentitlement-flowarn
	//
	FlowArn *string `field:"required" json:"flowArn" yaml:"flowArn"`
	// The name of the entitlement.
	//
	// This value must be unique within the current flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowentitlement.html#cfn-mediaconnect-flowentitlement-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The AWS account IDs that you want to share your content with.
	//
	// The receiving accounts (subscribers) will be allowed to create their own flows using your content as the source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowentitlement.html#cfn-mediaconnect-flowentitlement-subscribers
	//
	Subscribers *[]*string `field:"required" json:"subscribers" yaml:"subscribers"`
	// The percentage of the entitlement data transfer fee that you want the subscriber to be responsible for.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowentitlement.html#cfn-mediaconnect-flowentitlement-datatransfersubscriberfeepercent
	//
	// Default: - 0.
	//
	DataTransferSubscriberFeePercent *float64 `field:"optional" json:"dataTransferSubscriberFeePercent" yaml:"dataTransferSubscriberFeePercent"`
	// Encryption information.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowentitlement.html#cfn-mediaconnect-flowentitlement-encryption
	//
	Encryption interface{} `field:"optional" json:"encryption" yaml:"encryption"`
	// An indication of whether the new entitlement should be enabled or disabled as soon as it is created.
	//
	// If you don’t specify the entitlementStatus field in your request, MediaConnect sets it to ENABLED.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowentitlement.html#cfn-mediaconnect-flowentitlement-entitlementstatus
	//
	EntitlementStatus *string `field:"optional" json:"entitlementStatus" yaml:"entitlementStatus"`
}

