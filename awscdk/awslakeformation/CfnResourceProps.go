package awslakeformation


// Properties for defining a `CfnResource`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResourceProps := &CfnResourceProps{
//   	ResourceArn: jsii.String("resourceArn"),
//   	UseServiceLinkedRole: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	HybridAccessEnabled: jsii.Boolean(false),
//   	RoleArn: jsii.String("roleArn"),
//   	WithFederation: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-resource.html
//
type CfnResourceProps struct {
	// The Amazon Resource Name (ARN) of the resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-resource.html#cfn-lakeformation-resource-resourcearn
	//
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
	// Designates a trusted caller, an IAM principal, by registering this caller with the Data Catalog .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-resource.html#cfn-lakeformation-resource-useservicelinkedrole
	//
	UseServiceLinkedRole interface{} `field:"required" json:"useServiceLinkedRole" yaml:"useServiceLinkedRole"`
	// Indicates whether the data access of tables pointing to the location can be managed by both Lake Formation permissions as well as Amazon S3 bucket policies.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-resource.html#cfn-lakeformation-resource-hybridaccessenabled
	//
	HybridAccessEnabled interface{} `field:"optional" json:"hybridAccessEnabled" yaml:"hybridAccessEnabled"`
	// The IAM role that registered a resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-resource.html#cfn-lakeformation-resource-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Allows Lake Formation to assume a role to access tables in a federated database.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-resource.html#cfn-lakeformation-resource-withfederation
	//
	WithFederation interface{} `field:"optional" json:"withFederation" yaml:"withFederation"`
}

