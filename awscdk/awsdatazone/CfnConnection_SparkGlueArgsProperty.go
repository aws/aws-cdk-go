package awsdatazone


// The Spark AWS Glue args.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sparkGlueArgsProperty := &SparkGlueArgsProperty{
//   	Connection: jsii.String("connection"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-sparkglueargs.html
//
type CfnConnection_SparkGlueArgsProperty struct {
	// The connection in the Spark AWS Glue args.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-sparkglueargs.html#cfn-datazone-connection-sparkglueargs-connection
	//
	Connection *string `field:"optional" json:"connection" yaml:"connection"`
}

