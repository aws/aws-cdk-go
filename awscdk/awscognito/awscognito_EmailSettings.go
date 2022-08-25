package awscognito


// Email settings for the user pool.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   emailSettings := &emailSettings{
//   	from: jsii.String("from"),
//   	replyTo: jsii.String("replyTo"),
//   }
//
type EmailSettings struct {
	// The 'from' address on the emails received by the user.
	From *string `field:"optional" json:"from" yaml:"from"`
	// The 'replyTo' address on the emails received by the user as defined by IETF RFC-5322.
	//
	// When set, most email clients recognize to change 'to' line to this address when a reply is drafted.
	ReplyTo *string `field:"optional" json:"replyTo" yaml:"replyTo"`
}

