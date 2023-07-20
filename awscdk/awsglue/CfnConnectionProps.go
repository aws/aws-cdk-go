package awsglue


// Properties for defining a `CfnConnection`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var connectionProperties interface{}
//
//   cfnConnectionProps := &CfnConnectionProps{
//   	CatalogId: jsii.String("catalogId"),
//   	ConnectionInput: &ConnectionInputProperty{
//   		ConnectionType: jsii.String("connectionType"),
//
//   		// the properties below are optional
//   		ConnectionProperties: connectionProperties,
//   		Description: jsii.String("description"),
//   		MatchCriteria: []*string{
//   			jsii.String("matchCriteria"),
//   		},
//   		Name: jsii.String("name"),
//   		PhysicalConnectionRequirements: &PhysicalConnectionRequirementsProperty{
//   			AvailabilityZone: jsii.String("availabilityZone"),
//   			SecurityGroupIdList: []*string{
//   				jsii.String("securityGroupIdList"),
//   			},
//   			SubnetId: jsii.String("subnetId"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-connection.html
//
type CfnConnectionProps struct {
	// The ID of the data catalog to create the catalog object in.
	//
	// Currently, this should be the AWS account ID.
	//
	// > To specify the account ID, you can use the `Ref` intrinsic function with the `AWS::AccountId` pseudo parameter. For example: `!Ref AWS::AccountId` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-connection.html#cfn-glue-connection-catalogid
	//
	CatalogId *string `field:"required" json:"catalogId" yaml:"catalogId"`
	// The connection that you want to create.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-connection.html#cfn-glue-connection-connectioninput
	//
	ConnectionInput interface{} `field:"required" json:"connectionInput" yaml:"connectionInput"`
}

