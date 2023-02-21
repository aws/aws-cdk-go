package awsappstream


// Properties for defining a `CfnEntitlement`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEntitlementProps := &cfnEntitlementProps{
//   	appVisibility: jsii.String("appVisibility"),
//   	attributes: []interface{}{
//   		&attributeProperty{
//   			name: jsii.String("name"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	name: jsii.String("name"),
//   	stackName: jsii.String("stackName"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   }
//
type CfnEntitlementProps struct {
	// Specifies whether to entitle all apps or only selected apps.
	AppVisibility *string `field:"required" json:"appVisibility" yaml:"appVisibility"`
	// The attributes of the entitlement.
	Attributes interface{} `field:"required" json:"attributes" yaml:"attributes"`
	// The name of the entitlement.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The name of the stack.
	StackName *string `field:"required" json:"stackName" yaml:"stackName"`
	// The description of the entitlement.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

