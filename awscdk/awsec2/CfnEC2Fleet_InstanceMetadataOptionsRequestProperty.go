package awsec2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   instanceMetadataOptionsRequestProperty := &InstanceMetadataOptionsRequestProperty{
//   	HttpEndpoint: jsii.String("httpEndpoint"),
//   	HttpPutResponseHopLimit: jsii.Number(123),
//   	HttpTokens: jsii.String("httpTokens"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-instancemetadataoptionsrequest.html
//
type CfnEC2Fleet_InstanceMetadataOptionsRequestProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-instancemetadataoptionsrequest.html#cfn-ec2-ec2fleet-instancemetadataoptionsrequest-httpendpoint
	//
	HttpEndpoint *string `field:"optional" json:"httpEndpoint" yaml:"httpEndpoint"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-instancemetadataoptionsrequest.html#cfn-ec2-ec2fleet-instancemetadataoptionsrequest-httpputresponsehoplimit
	//
	HttpPutResponseHopLimit *float64 `field:"optional" json:"httpPutResponseHopLimit" yaml:"httpPutResponseHopLimit"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-instancemetadataoptionsrequest.html#cfn-ec2-ec2fleet-instancemetadataoptionsrequest-httptokens
	//
	HttpTokens *string `field:"optional" json:"httpTokens" yaml:"httpTokens"`
}

