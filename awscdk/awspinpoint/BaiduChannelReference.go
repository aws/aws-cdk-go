package awspinpoint


// A reference to a BaiduChannel resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   baiduChannelReference := &BaiduChannelReference{
//   	BaiduChannelId: jsii.String("baiduChannelId"),
//   }
//
type BaiduChannelReference struct {
	// The Id of the BaiduChannel resource.
	BaiduChannelId *string `field:"required" json:"baiduChannelId" yaml:"baiduChannelId"`
}

