package awsmedialivealpha


// Properties for creating a MediaLive Input.
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
type InputProps struct {
	// Input configuration that defines the input type and source settings.
	// Experimental.
	Input InputConfiguration `field:"required" json:"input" yaml:"input"`
	// Input name.
	// Default: - auto-generated.
	//
	// Experimental.
	InputName *string `field:"optional" json:"inputName" yaml:"inputName"`
	// The network location of the input — AWS cloud or on-premises (MediaLive Anywhere).
	// Default: - AWS, applied by MediaLive.
	//
	// Experimental.
	InputNetworkLocation InputNetworkLocation `field:"optional" json:"inputNetworkLocation" yaml:"inputNetworkLocation"`
	// Tags to add to the input.
	// Default: - no tags.
	//
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

