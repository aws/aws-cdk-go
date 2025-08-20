package awsglue


// An object that describes the VPC configuration for a table optimizer.
//
// This configuration is necessary to perform optimization on tables that are in a customer VPC.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcConfigurationProperty := &VpcConfigurationProperty{
//   	GlueConnectionName: jsii.String("glueConnectionName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-tableoptimizer-vpcconfiguration.html
//
type CfnTableOptimizer_VpcConfigurationProperty struct {
	// The name of the AWS Glue connection used for the VPC for the table optimizer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-tableoptimizer-vpcconfiguration.html#cfn-glue-tableoptimizer-vpcconfiguration-glueconnectionname
	//
	GlueConnectionName *string `field:"optional" json:"glueConnectionName" yaml:"glueConnectionName"`
}

