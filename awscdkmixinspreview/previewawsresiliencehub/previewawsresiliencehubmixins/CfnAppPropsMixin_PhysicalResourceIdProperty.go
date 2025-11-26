package previewawsresiliencehubmixins


// Defines a physical resource identifier.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   physicalResourceIdProperty := &PhysicalResourceIdProperty{
//   	AwsAccountId: jsii.String("awsAccountId"),
//   	AwsRegion: jsii.String("awsRegion"),
//   	Identifier: jsii.String("identifier"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehub-app-physicalresourceid.html
//
type CfnAppPropsMixin_PhysicalResourceIdProperty struct {
	// The AWS account that owns the physical resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehub-app-physicalresourceid.html#cfn-resiliencehub-app-physicalresourceid-awsaccountid
	//
	AwsAccountId *string `field:"optional" json:"awsAccountId" yaml:"awsAccountId"`
	// The AWS Region that the physical resource is located in.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehub-app-physicalresourceid.html#cfn-resiliencehub-app-physicalresourceid-awsregion
	//
	AwsRegion *string `field:"optional" json:"awsRegion" yaml:"awsRegion"`
	// Identifier of the physical resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehub-app-physicalresourceid.html#cfn-resiliencehub-app-physicalresourceid-identifier
	//
	Identifier *string `field:"optional" json:"identifier" yaml:"identifier"`
	// Specifies the type of physical resource identifier.
	//
	// - **Arn** - The resource identifier is an Amazon Resource Name (ARN) and it can identify the following list of resources:
	//
	// - `AWS::ECS::Service`
	// - `AWS::EFS::FileSystem`
	// - `AWS::ElasticLoadBalancingV2::LoadBalancer`
	// - `AWS::Lambda::Function`
	// - `AWS::SNS::Topic`
	// - **Native** - The resource identifier is an AWS Resilience Hub -native identifier and it can identify the following list of resources:
	//
	// - `AWS::ApiGateway::RestApi`
	// - `AWS::ApiGatewayV2::Api`
	// - `AWS::AutoScaling::AutoScalingGroup`
	// - `AWS::DocDB::DBCluster`
	// - `AWS::DocDB::DBGlobalCluster`
	// - `AWS::DocDB::DBInstance`
	// - `AWS::DynamoDB::GlobalTable`
	// - `AWS::DynamoDB::Table`
	// - `AWS::EC2::EC2Fleet`
	// - `AWS::EC2::Instance`
	// - `AWS::EC2::NatGateway`
	// - `AWS::EC2::Volume`
	// - `AWS::ElasticLoadBalancing::LoadBalancer`
	// - `AWS::RDS::DBCluster`
	// - `AWS::RDS::DBInstance`
	// - `AWS::RDS::GlobalCluster`
	// - `AWS::Route53::RecordSet`
	// - `AWS::S3::Bucket`
	// - `AWS::SQS::Queue`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehub-app-physicalresourceid.html#cfn-resiliencehub-app-physicalresourceid-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

