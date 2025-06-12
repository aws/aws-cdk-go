package awsmedialive


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   multicastSourceCreateRequestProperty := &MulticastSourceCreateRequestProperty{
//   	SourceIp: jsii.String("sourceIp"),
//   	Url: jsii.String("url"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-multicastsourcecreaterequest.html
//
type CfnInput_MulticastSourceCreateRequestProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-multicastsourcecreaterequest.html#cfn-medialive-input-multicastsourcecreaterequest-sourceip
	//
	SourceIp *string `field:"optional" json:"sourceIp" yaml:"sourceIp"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-multicastsourcecreaterequest.html#cfn-medialive-input-multicastsourcecreaterequest-url
	//
	Url *string `field:"optional" json:"url" yaml:"url"`
}

