package awsmediaconnect


// Properties for defining a `CfnFlowEntitlement`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFlowEntitlementProps := &cfnFlowEntitlementProps{
//   	description: jsii.String("description"),
//   	flowArn: jsii.String("flowArn"),
//   	name: jsii.String("name"),
//   	subscribers: []*string{
//   		jsii.String("subscribers"),
//   	},
//
//   	// the properties below are optional
//   	dataTransferSubscriberFeePercent: jsii.Number(123),
//   	encryption: &encryptionProperty{
//   		algorithm: jsii.String("algorithm"),
//   		roleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
//   		constantInitializationVector: jsii.String("constantInitializationVector"),
//   		deviceId: jsii.String("deviceId"),
//   		keyType: jsii.String("keyType"),
//   		region: jsii.String("region"),
//   		resourceId: jsii.String("resourceId"),
//   		secretArn: jsii.String("secretArn"),
//   		url: jsii.String("url"),
//   	},
//   	entitlementStatus: jsii.String("entitlementStatus"),
//   }
//
type CfnFlowEntitlementProps struct {
	// A description of the entitlement.
	//
	// This description appears only on the MediaConnect console and is not visible outside of the current AWS account.
	Description *string `field:"required" json:"description" yaml:"description"`
	// The Amazon Resource Name (ARN) of the flow.
	FlowArn *string `field:"required" json:"flowArn" yaml:"flowArn"`
	// The name of the entitlement.
	//
	// This value must be unique within the current flow.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The AWS account IDs that you want to share your content with.
	//
	// The receiving accounts (subscribers) will be allowed to create their own flows using your content as the source.
	Subscribers *[]*string `field:"required" json:"subscribers" yaml:"subscribers"`
	// The percentage of the entitlement data transfer fee that you want the subscriber to be responsible for.
	DataTransferSubscriberFeePercent *float64 `field:"optional" json:"dataTransferSubscriberFeePercent" yaml:"dataTransferSubscriberFeePercent"`
	// The type of encryption that MediaConnect will use on the output that is associated with the entitlement.
	Encryption interface{} `field:"optional" json:"encryption" yaml:"encryption"`
	// An indication of whether the new entitlement should be enabled or disabled as soon as it is created.
	//
	// If you don’t specify the entitlementStatus field in your request, MediaConnect sets it to ENABLED.
	EntitlementStatus *string `field:"optional" json:"entitlementStatus" yaml:"entitlementStatus"`
}

