package awsresiliencehubv2


// Resource configuration for an input source.
//
// Provide exactly one field.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceConfigurationProperty := &ResourceConfigurationProperty{
//   	CfnStackArn: jsii.String("cfnStackArn"),
//   	DesignFileS3Url: jsii.String("designFileS3Url"),
//   	Eks: &EksSourceProperty{
//   		ClusterArn: jsii.String("clusterArn"),
//   		Namespaces: []*string{
//   			jsii.String("namespaces"),
//   		},
//   	},
//   	ResourceTags: []interface{}{
//   		&ResourceTagProperty{
//   			Key: jsii.String("key"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   	TfStateFileUrl: jsii.String("tfStateFileUrl"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-service-resourceconfiguration.html
//
type CfnService_ResourceConfigurationProperty struct {
	// ARN of a CloudFormation stack.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-service-resourceconfiguration.html#cfn-resiliencehubv2-service-resourceconfiguration-cfnstackarn
	//
	CfnStackArn *string `field:"optional" json:"cfnStackArn" yaml:"cfnStackArn"`
	// S3 URL of a design file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-service-resourceconfiguration.html#cfn-resiliencehubv2-service-resourceconfiguration-designfiles3url
	//
	DesignFileS3Url *string `field:"optional" json:"designFileS3Url" yaml:"designFileS3Url"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-service-resourceconfiguration.html#cfn-resiliencehubv2-service-resourceconfiguration-eks
	//
	Eks interface{} `field:"optional" json:"eks" yaml:"eks"`
	// Resource tags to discover resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-service-resourceconfiguration.html#cfn-resiliencehubv2-service-resourceconfiguration-resourcetags
	//
	ResourceTags interface{} `field:"optional" json:"resourceTags" yaml:"resourceTags"`
	// URL of a Terraform state file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-service-resourceconfiguration.html#cfn-resiliencehubv2-service-resourceconfiguration-tfstatefileurl
	//
	TfStateFileUrl *string `field:"optional" json:"tfStateFileUrl" yaml:"tfStateFileUrl"`
}

