package awsemrcontainers


// The job driver for job type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   sparkSqlJobDriverProperty := &SparkSqlJobDriverProperty{
//   	EntryPoint: jsii.String("entryPoint"),
//   	SparkSqlParameters: jsii.String("sparkSqlParameters"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-jobtemplate-sparksqljobdriver.html
//
type CfnJobTemplatePropsMixin_SparkSqlJobDriverProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-jobtemplate-sparksqljobdriver.html#cfn-emrcontainers-jobtemplate-sparksqljobdriver-entrypoint
	//
	EntryPoint *string `field:"optional" json:"entryPoint" yaml:"entryPoint"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-jobtemplate-sparksqljobdriver.html#cfn-emrcontainers-jobtemplate-sparksqljobdriver-sparksqlparameters
	//
	SparkSqlParameters *string `field:"optional" json:"sparkSqlParameters" yaml:"sparkSqlParameters"`
}

