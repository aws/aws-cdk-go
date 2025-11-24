package mixinsawsmediaconnect


// Properties for CfnFlowEntitlementPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnFlowEntitlementMixinProps := &CfnFlowEntitlementMixinProps{
//   	DataTransferSubscriberFeePercent: jsii.Number(123),
//   	Description: jsii.String("description"),
//   	Encryption: &EncryptionProperty{
//   		Algorithm: jsii.String("algorithm"),
//   		ConstantInitializationVector: jsii.String("constantInitializationVector"),
//   		DeviceId: jsii.String("deviceId"),
//   		KeyType: jsii.String("keyType"),
//   		Region: jsii.String("region"),
//   		ResourceId: jsii.String("resourceId"),
//   		RoleArn: jsii.String("roleArn"),
//   		SecretArn: jsii.String("secretArn"),
//   		Url: jsii.String("url"),
//   	},
//   	EntitlementStatus: jsii.String("entitlementStatus"),
//   	FlowArn: jsii.String("flowArn"),
//   	Name: jsii.String("name"),
//   	Subscribers: []*string{
//   		jsii.String("subscribers"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowentitlement.html
//
type CfnFlowEntitlementMixinProps struct {
	// The percentage of the entitlement data transfer fee that you want the subscriber to be responsible for.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowentitlement.html#cfn-mediaconnect-flowentitlement-datatransfersubscriberfeepercent
	//
	// Default: - 0.
	//
	DataTransferSubscriberFeePercent *float64 `field:"optional" json:"dataTransferSubscriberFeePercent" yaml:"dataTransferSubscriberFeePercent"`
	// A description of the entitlement.
	//
	// This description appears only on the MediaConnect console and is not visible outside of the current AWS account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowentitlement.html#cfn-mediaconnect-flowentitlement-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Information about the encryption of the flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowentitlement.html#cfn-mediaconnect-flowentitlement-encryption
	//
	Encryption interface{} `field:"optional" json:"encryption" yaml:"encryption"`
	// An indication of whether the new entitlement should be enabled or disabled as soon as it is created.
	//
	// If you don’t specify the entitlementStatus field in your request, MediaConnect sets it to ENABLED.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowentitlement.html#cfn-mediaconnect-flowentitlement-entitlementstatus
	//
	EntitlementStatus *string `field:"optional" json:"entitlementStatus" yaml:"entitlementStatus"`
	// The Amazon Resource Name (ARN) of the flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowentitlement.html#cfn-mediaconnect-flowentitlement-flowarn
	//
	FlowArn *string `field:"optional" json:"flowArn" yaml:"flowArn"`
	// The name of the entitlement.
	//
	// This value must be unique within the current flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowentitlement.html#cfn-mediaconnect-flowentitlement-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The AWS account IDs that you want to share your content with.
	//
	// The receiving accounts (subscribers) will be allowed to create their own flows using your content as the source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowentitlement.html#cfn-mediaconnect-flowentitlement-subscribers
	//
	Subscribers *[]*string `field:"optional" json:"subscribers" yaml:"subscribers"`
}

