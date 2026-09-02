package awsmedialivealpha


// Properties for creating a MediaLive Input Security Group.
//
// Example:
//   var stack Stack
//   var passphrase ISecret
//
//
//   sg := medialive.NewInputSecurityGroup(stack, jsii.String("SrtSg"), &InputSecurityGroupProps{
//   	AllowlistRules: []*string{
//   		jsii.String("203.0.113.0/24"),
//   	},
//   })
//
//   medialive.NewInput(stack, jsii.String("SrtListenerInput"), &InputProps{
//   	InputName: jsii.String("srt-listener"),
//   	Input: medialive.InputConfiguration_SrtListener(&SrtListenerInputProps{
//   		InputSecurityGroups: []IInputSecurityGroupRef{
//   			sg,
//   		},
//   		MinimumLatency: awscdk.Duration_Millis(jsii.Number(500)),
//   		StreamId: jsii.String("my-stream-id"),
//   		Decryption: &SrtDecryptionProps{
//   			Algorithm: medialive.SrtDecryptionAlgorithm_AES256(),
//   			PassphraseSecret: passphrase,
//   		},
//   	}),
//   })
//
// Experimental.
type InputSecurityGroupProps struct {
	// The list of IPv4 CIDR addresses to allow.
	// Experimental.
	AllowlistRules *[]*string `field:"required" json:"allowlistRules" yaml:"allowlistRules"`
	// Tags to add to the input security group.
	// Default: - no tags.
	//
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

