package awsquicksight


// VPC connection properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcConnectionPropertiesProperty := &vpcConnectionPropertiesProperty{
//   	vpcConnectionArn: jsii.String("vpcConnectionArn"),
//   }
//
type CfnDataSource_VpcConnectionPropertiesProperty struct {
	// The Amazon Resource Name (ARN) for the VPC connection.
	VpcConnectionArn *string `field:"required" json:"vpcConnectionArn" yaml:"vpcConnectionArn"`
}

