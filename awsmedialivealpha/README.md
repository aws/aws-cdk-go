# AWS::MediaLive Construct Library

<!--BEGIN STABILITY BANNER-->---


![cdk-constructs: Experimental](https://img.shields.io/badge/cdk--constructs-experimental-important.svg?style=for-the-badge)

> The APIs of higher level constructs in this module are experimental and under active development.
> They are subject to non-backward compatible changes or removal in any future version. These are
> not subject to the [Semantic Versioning](https://semver.org/) model and breaking changes will be
> announced in the release notes. This means that while you may use them, you may need to update
> your source code when upgrading to a newer version of this package.

---
<!--END STABILITY BANNER-->

## AWS Elemental MediaLive

AWS Elemental MediaLive is a real-time video service that lets you create live outputs for broadcast and streaming delivery.

This package contains constructs for working with AWS Elemental MediaLive, including Inputs, Input Security Groups, Channels, and MediaLive Anywhere resources (Networks, Clusters, Channel Placement Groups, SDI Sources).

For further information on AWS Elemental MediaLive, see [the documentation](https://docs.aws.amazon.com/medialive/latest/ug/what-is.html). See [supported codecs per output group](https://docs.aws.amazon.com/medialive/latest/ug/outputs-supported-codecs.html).

The following example creates an SRT caller input, encodes it to H.264 + AAC, and outputs HLS segments to an S3 bucket:

```go
var stack Stack
var bucket IBucket


input := medialive.NewInput(stack, jsii.String("SrtInput"), &InputProps{
	InputName: jsii.String("my-srt-input"),
	Input: medialive.InputConfiguration_SrtCaller([]SrtCallerSourceProps{
		&SrtCallerSourceProps{
			SrtListenerAddress: jsii.String("203.0.113.10"),
			SrtListenerPort: jsii.Number(5000),
		},
	}),
})

video := medialive.EncodeConfiguration_Video(&VideoEncodeProps{
	Name: jsii.String("video_720p"),
	Codec: medialive.VideoCodecSettings_H264(&H264SettingsProps{
		RateControl: medialive.H264RateControl_Cbr(&CbrRateControlProps{
			Bitrate: awscdk.Bitrate_Mbps(jsii.Number(3)),
		}),
		Framerate: medialive.Framerate_FPS_30(),
	}),
	Width: jsii.Number(1280),
	Height: jsii.Number(720),
})

audio := medialive.EncodeConfiguration_Audio(&AudioEncodeProps{
	Name: jsii.String("audio_aac"),
	Codec: medialive.AudioCodecSettings_Aac(&AacSettingsProps{
		Bitrate: awscdk.Bitrate_Kbps(jsii.Number(192)),
	}),
})

medialive.NewChannel(stack, jsii.String("Channel"), &ChannelProps{
	Inputs: []InputAttachment{
		&InputAttachment{
			Input: *Input,
		},
	},
	OutputGroups: []OutputGroupConfiguration{
		medialive.OutputGroupConfiguration_Hls(&HlsOutputGroupProps{
			Name: jsii.String("hls"),
			Destinations: []OutputDestination{
				medialive.OutputDestination_ToBucket(bucket, jsii.String("live/stream")),
			},
			Outputs: []HlsOutputDefinition{
				&HlsOutputDefinition{
					Encodes: []EncodeConfiguration{
						video,
						audio,
					},
					OutputName: jsii.String("hls_out"),
				},
			},
		}),
	},
})
```

## Input

An input represents the upstream source that feeds a MediaLive channel. Use `InputConfiguration` factory methods to create different input types.

### SRT Caller

MediaLive connects to a remote SRT listener:

```go
var stack Stack

medialive.NewInput(stack, jsii.String("SrtInput"), &InputProps{
	InputName: jsii.String("srt-caller"),
	Input: medialive.InputConfiguration_SrtCaller([]SrtCallerSourceProps{
		&SrtCallerSourceProps{
			SrtListenerAddress: jsii.String("203.0.113.10"),
			SrtListenerPort: jsii.Number(5000),
		},
	}),
})
```

### SRT Listener

MediaLive listens for an incoming SRT connection. SRT listener inputs require an input security
group. To receive encrypted content, supply a `decryption` block referencing a Secrets Manager
secret that holds the passphrase — the secret is passed by reference, so MediaLive resolves the ARN
at synth time:

```go
var stack Stack
var passphrase ISecret


sg := medialive.NewInputSecurityGroup(stack, jsii.String("SrtSg"), &InputSecurityGroupProps{
	AllowlistRules: []*string{
		jsii.String("203.0.113.0/24"),
	},
})

medialive.NewInput(stack, jsii.String("SrtListenerInput"), &InputProps{
	InputName: jsii.String("srt-listener"),
	Input: medialive.InputConfiguration_SrtListener(&SrtListenerInputProps{
		InputSecurityGroups: []IInputSecurityGroupRef{
			sg,
		},
		MinimumLatency: awscdk.Duration_Millis(jsii.Number(500)),
		StreamId: jsii.String("my-stream-id"),
		Decryption: &SrtDecryptionProps{
			Algorithm: medialive.SrtDecryptionAlgorithm_AES256(),
			PassphraseSecret: passphrase,
		},
	}),
})
```

### AWS Elemental MediaConnect Router

Creates a MediaConnect Router Input with automatic encryption:

```go
var stack Stack

medialive.NewInput(stack, jsii.String("RouterInput"), &InputProps{
	InputName: jsii.String("mc-router"),
	Input: medialive.InputConfiguration_MediaConnectRouter(),
})
```

An input created this way is the only kind `@aws-cdk/aws-mediaconnect-alpha`'s `RouterOutputConfiguration.mediaLiveInput()` can deliver to — pointing it at any other input type synths but fails at deploy.

### MP4 File from S3

Use `InputSource.fromBucket()` to reference an S3 object:

```go
var stack Stack
var bucket IBucket


medialive.NewInput(stack, jsii.String("FileInput"), &InputProps{
	InputName: jsii.String("mp4-file"),
	Input: medialive.InputConfiguration_Mp4File([]InputSource{
		medialive.InputSource_FromBucket(bucket, jsii.String("media/input.mp4")),
	}),
})
```

### Importing an Existing Input

```go
var stack Stack

input := medialive.Input_FromInputArn(stack, jsii.String("Imported"), jsii.String("arn:aws:medialive:us-east-1:123456789012:input:1234567"))
```

## Input Security Group

An input security group controls which IPv4 CIDR blocks can push content to a push-type input.

```go
var stack Stack

sg := medialive.NewInputSecurityGroup(stack, jsii.String("SG"), &InputSecurityGroupProps{
	AllowlistRules: []*string{
		jsii.String("203.0.113.0/24"),
	},
})
```

### Importing an Existing Input Security Group

```go
var stack Stack

sg := medialive.InputSecurityGroup_FromInputSecurityGroupArn(stack, jsii.String("Imported"), jsii.String("arn:aws:medialive:us-east-1:123456789012:inputSecurityGroup:1234567"))
```

## Channel

A channel takes one or more inputs, encodes them, and produces output groups. If no `role` is provided, the channel auto-creates an IAM role with the `medialive.amazonaws.com` service principal.

Minimal example — single input, single HLS output:

```go
var stack Stack
var input IInput
var bucket IBucket


video := medialive.EncodeConfiguration_Video(&VideoEncodeProps{
	Name: jsii.String("video_720p"),
	Codec: medialive.VideoCodecSettings_H264(&H264SettingsProps{
		RateControl: medialive.H264RateControl_Cbr(&CbrRateControlProps{
			Bitrate: awscdk.Bitrate_Mbps(jsii.Number(3)),
		}),
		Framerate: medialive.Framerate_FPS_30(),
	}),
	Width: jsii.Number(1280),
	Height: jsii.Number(720),
})

audio := medialive.EncodeConfiguration_Audio(&AudioEncodeProps{
	Name: jsii.String("audio_aac"),
	Codec: medialive.AudioCodecSettings_Aac(&AacSettingsProps{
		Bitrate: awscdk.Bitrate_Kbps(jsii.Number(192)),
	}),
})

medialive.NewChannel(stack, jsii.String("Channel"), &ChannelProps{
	Inputs: []InputAttachment{
		&InputAttachment{
			Input: *Input,
		},
	},
	OutputGroups: []OutputGroupConfiguration{
		medialive.OutputGroupConfiguration_Hls(&HlsOutputGroupProps{
			Name: jsii.String("hls"),
			Destinations: []OutputDestination{
				medialive.OutputDestination_ToBucket(bucket, jsii.String("live/stream")),
			},
			Outputs: []HlsOutputDefinition{
				&HlsOutputDefinition{
					Encodes: []EncodeConfiguration{
						video,
						audio,
					},
					OutputName: jsii.String("hls_out"),
				},
			},
		}),
	},
})
```

### STANDARD Channel with MediaPackage V2

A STANDARD channel runs two pipelines for redundancy. Each output group needs two destinations — one per pipeline.

```go
var stack Stack
var input IInput
var mpChannel IChannel


hdVideo := medialive.EncodeConfiguration_Video(&VideoEncodeProps{
	Name: jsii.String("video_1080p"),
	Codec: medialive.VideoCodecSettings_H265(&H265SettingsProps{
		RateControl: medialive.H265RateControl_Qvbr(&QvbrRateControlProps{
			MaxBitrate: awscdk.Bitrate_Mbps(jsii.Number(8)),
			QvbrQualityLevel: jsii.Number(7),
		}),
		Framerate: medialive.Framerate_FPS_30(),
	}),
	Width: jsii.Number(1920),
	Height: jsii.Number(1080),
})

sdVideo := medialive.EncodeConfiguration_Video(&VideoEncodeProps{
	Name: jsii.String("video_480p"),
	Codec: medialive.VideoCodecSettings_*H265(&H265SettingsProps{
		RateControl: medialive.H265RateControl_*Qvbr(&QvbrRateControlProps{
			MaxBitrate: awscdk.Bitrate_*Mbps(jsii.Number(2)),
			QvbrQualityLevel: jsii.Number(7),
		}),
		Framerate: medialive.Framerate_FPS_30(),
	}),
	Width: jsii.Number(854),
	Height: jsii.Number(480),
})

audio := medialive.EncodeConfiguration_Audio(&AudioEncodeProps{
	Name: jsii.String("audio_aac"),
	Codec: medialive.AudioCodecSettings_Aac(&AacSettingsProps{
		Bitrate: awscdk.Bitrate_Kbps(jsii.Number(192)),
	}),
})

medialive.NewChannel(stack, jsii.String("Channel"), &ChannelProps{
	ChannelClass: medialive.ChannelClass_STANDARD(),
	Inputs: []InputAttachment{
		&InputAttachment{
			Input: *Input,
		},
	},
	OutputGroups: []OutputGroupConfiguration{
		medialive.OutputGroupConfiguration_MediaPackageV2(&MediaPackageV2OutputGroupProps{
			Name: jsii.String("emp"),
			Channel: mpChannel,
			Outputs: []MediaPackageV2OutputDefinition{
				&MediaPackageV2OutputDefinition{
					Encode: hdVideo,
					OutputName: jsii.String("hd"),
				},
				&MediaPackageV2OutputDefinition{
					Encode: sdVideo,
					OutputName: jsii.String("sd"),
				},
				&MediaPackageV2OutputDefinition{
					Encode: audio,
					OutputName: jsii.String("audio"),
				},
			},
		}),
	},
})
```

### Global Configuration

`globalConfiguration` sets channel-wide behaviour: how the pipelines are locked together and the output timing source. All fields are optional and fall back to MediaLive defaults.

```go
var stack Stack
var input IInput
var bucket IBucket
var video EncodeConfiguration
var audio EncodeConfiguration


medialive.NewChannel(stack, jsii.String("Channel"), &ChannelProps{
	Inputs: []InputAttachment{
		&InputAttachment{
			Input: *Input,
		},
	},
	TimecodeConfig: &TimecodeConfig{
		Source: medialive.TimecodeSource_EMBEDDED(),
	},
	GlobalConfiguration: &GlobalConfiguration{
		OutputLocking: medialive.OutputLocking_Epoch(),
		OutputTimingSource: medialive.OutputTimingSource_INPUT_CLOCK(),
	},
	OutputGroups: []OutputGroupConfiguration{
		medialive.OutputGroupConfiguration_Hls(&HlsOutputGroupProps{
			Name: jsii.String("hls"),
			Destinations: []OutputDestination{
				medialive.OutputDestination_ToBucket(bucket, jsii.String("live/stream")),
			},
			Outputs: []HlsOutputDefinition{
				&HlsOutputDefinition{
					Encodes: []EncodeConfiguration{
						video,
						audio,
					},
					OutputName: jsii.String("hls_out"),
				},
			},
		}),
	},
})
```

#### Output locking

`outputLocking` synchronises the frames emitted by a channel's two pipelines. Pick a strategy with
the `OutputLocking` factory:

* `OutputLocking.pipeline()` — synchronise each pipeline's output to the other. Choose how with
  `method`: `PipelineLockingMethod.SOURCE_TIMECODE` (default, needs reliable embedded timecodes) or
  `PipelineLockingMethod.VIDEO_ALIGNMENT` (visual content matching, no timecodes required).
* `OutputLocking.epoch()` — synchronise to the Unix epoch (optionally a `customEpoch`/`jamSyncTime`).
  Requires `outputTimingSource: OutputTimingSource.INPUT_CLOCK` (enforced at synth).
* `OutputLocking.disabled()` — no synchronisation.

```go
// Video-aligned pipeline locking — useful when sources lack reliable timecodes
locking := medialive.OutputLocking_Pipeline(&PipelineOutputLockingProps{
	Method: medialive.PipelineLockingMethod_VIDEO_ALIGNMENT(),
})
```

#### Input-loss behavior

`inputLossBehavior` controls what MediaLive emits when the input is lost: a black period, then a
repeated frame, then either a solid colour or a slate image. Provide the slate as a
[`FileLocation`](#file-locations).

```go
var slateBucket IBucket


inputLoss := &InputLossBehavior{
	BlackFrame: awscdk.Duration_Seconds(jsii.Number(1)),
	RepeatFrame: awscdk.Duration_*Seconds(jsii.Number(5)),
	ImageType: medialive.InputLossImageType_SLATE(),
	ImageSlate: medialive.FileLocation_FromBucket(slateBucket, jsii.String("slates/offline.png")),
}
```

## File locations

Several channel features reference a file MediaLive reads at runtime — an input-loss slate, an
avail-blanking image, a blackout-slate image, or a burn-in caption font. These all take a
`FileLocation`, created from an S3 bucket (which auto-grants the channel role read access) or a URL
(with optional SSM-backed credentials):

```go
import "github.com/aws/aws-cdk-go/awscdk"

var bucket IBucket
var passwordParam StringParameter


// From an S3 bucket — the channel role is granted read access automatically
fromS3 := medialive.FileLocation_FromBucket(bucket, jsii.String("assets/slate.png"))

// From a URL with optional credentials (SSM parameter read access auto-granted)
fromUrl := medialive.FileLocation_Url(jsii.String("https://origin.example.com/font.ttf"), &FileLocationOptions{
	Username: jsii.String("ingest-user"),
	Password: passwordParam,
})
```

## Color correction

A channel can apply one or more color-space conversions to its video, optionally using a 3D LUT
to remap colors. Each `ColorCorrection` declares the `inputColorSpace` to match and the
`outputColorSpace` to convert to. MediaLive reads the LUT from S3 at runtime, so it must be an S3
location — provide it via `Lut.fromBucket()` (which uses the secure `s3ssl://` form and auto-grants
the channel role read access) or `Lut.url()` with an `s3://`/`s3ssl://` URL:

```go
var stack Stack
var bucket IBucket
var input IInput
var video EncodeConfiguration
var destination OutputDestination


medialive.NewChannel(stack, jsii.String("Channel"), &ChannelProps{
	Inputs: []InputAttachment{
		&InputAttachment{
			Input: *Input,
		},
	},
	ColorCorrections: []ColorCorrection{
		&ColorCorrection{
			InputColorSpace: medialive.ColorSpace_REC_601(),
			OutputColorSpace: medialive.ColorSpace_REC_709(),
			Lut: medialive.Lut_FromBucket(bucket, jsii.String("luts/rec601-to-rec709.cube")),
		},
	},
	OutputGroups: []OutputGroupConfiguration{
		medialive.OutputGroupConfiguration_Hls(&HlsOutputGroupProps{
			Name: jsii.String("hls"),
			Destinations: []OutputDestination{
				destination,
			},
			Outputs: []HlsOutputDefinition{
				&HlsOutputDefinition{
					Encodes: []EncodeConfiguration{
						video,
					},
					OutputName: jsii.String("video"),
				},
			},
		}),
	},
})
```

## Encode Configuration

Use `EncodeConfiguration.video()`, `EncodeConfiguration.audio()`, and `EncodeConfiguration.caption()` to define encodes.

### Video

```go
// H.264
h264 := medialive.EncodeConfiguration_Video(&VideoEncodeProps{
	Name: jsii.String("h264_720p"),
	Codec: medialive.VideoCodecSettings_H264(&H264SettingsProps{
		RateControl: medialive.H264RateControl_Cbr(&CbrRateControlProps{
			Bitrate: awscdk.Bitrate_Mbps(jsii.Number(3)),
		}),
		Framerate: medialive.Framerate_FPS_30(),
		Profile: medialive.H264Profile_HIGH(),
	}),
	Width: jsii.Number(1280),
	Height: jsii.Number(720),
})

// H.265
h265 := medialive.EncodeConfiguration_Video(&VideoEncodeProps{
	Name: jsii.String("h265_1080p"),
	Codec: medialive.VideoCodecSettings_H265(&H265SettingsProps{
		RateControl: medialive.H265RateControl_Qvbr(&QvbrRateControlProps{
			MaxBitrate: awscdk.Bitrate_*Mbps(jsii.Number(5)),
			QvbrQualityLevel: jsii.Number(7),
		}),
		Framerate: medialive.Framerate_FPS_30(),
		Profile: medialive.H265Profile_MAIN(),
		Tier: medialive.H265Tier_HIGH(),
	}),
	Width: jsii.Number(1920),
	Height: jsii.Number(1080),
})
```

Video codecs accept optional overrides for adaptive quantization, scene-change detection, color space, and more. See the props interfaces for the full list:

```go
hdr := medialive.EncodeConfiguration_Video(&VideoEncodeProps{
	Name: jsii.String("h265_hdr"),
	Codec: medialive.VideoCodecSettings_H265(&H265SettingsProps{
		RateControl: medialive.H265RateControl_Qvbr(&QvbrRateControlProps{
			MaxBitrate: awscdk.Bitrate_Mbps(jsii.Number(8)),
			QvbrQualityLevel: jsii.Number(8),
		}),
		Framerate: medialive.Framerate_FPS_30(),
		SceneChangeDetect: medialive.H265SceneChangeDetect_ENABLED(),
		ColorSpaceSettings: medialive.H265ColorSpaceSettings_Hlg2020(),
	}),
	Width: jsii.Number(1920),
	Height: jsii.Number(1080),
})
```

### Audio

```go
// AAC stereo
aac := medialive.EncodeConfiguration_Audio(&AudioEncodeProps{
	Name: jsii.String("aac_stereo"),
	Codec: medialive.AudioCodecSettings_Aac(&AacSettingsProps{
		Bitrate: awscdk.Bitrate_Kbps(jsii.Number(192)),
		CodingMode: medialive.AacCodingMode_CODING_MODE_2_0(),
	}),
})

// AC3 5.1
ac3 := medialive.EncodeConfiguration_Audio(&AudioEncodeProps{
	Name: jsii.String("ac3_surround"),
	Codec: medialive.AudioCodecSettings_Ac3(&Ac3SettingsProps{
		Bitrate: awscdk.Bitrate_*Kbps(jsii.Number(384)),
		CodingMode: medialive.Ac3CodingMode_CODING_MODE_3_2_LFE(),
	}),
})
```

### Caption

A caption encode converts a source caption track (referenced by `captionSelectorName`) to an
output format via the `CaptionDestination` factory. One selector can feed multiple encodes:

```go
// Define a caption selector on the input attachment (see Input Attachment Settings below)
captionSelector := medialive.CaptionSelector_Embedded(jsii.String("captions"))

// WebVTT captions — packaged alongside the video encode in the same output
webvtt := medialive.EncodeConfiguration_Caption(&CaptionEncodeProps{
	Name: jsii.String("eng_webvtt"),
	CaptionSelectorName: captionSelector.Name,
	LanguageCode: jsii.String("eng"),
	LanguageDescription: jsii.String("English"),
	Destination: medialive.CaptionDestination_Webvtt(),
})

// Burned-in captions — rendered into the video, styled via the burn-in options
burnIn := medialive.EncodeConfiguration_Caption(&CaptionEncodeProps{
	Name: jsii.String("eng_burnin"),
	CaptionSelectorName: captionSelector.*Name,
	Destination: medialive.CaptionDestination_BurnIn(&BurnInDestinationProps{
		Alignment: medialive.CaptionAlignment_CENTERED(),
		FontColor: medialive.CaptionFontColor_WHITE(),
		OutlineColor: medialive.CaptionOutlineColor_BLACK(),
		FontSize: medialive.CaptionFontSize_AUTO(),
	}),
})
```

## Cross-service integrations

| Destination | MediaLive side | Other side | Package |
|---|---|---|---|
| MediaPackage V2 | `medialive.OutputGroupConfiguration.mediaPackageV2()` | `mediapackagev2.Channel` | `@aws-cdk/aws-mediapackagev2-alpha` |
| MediaConnect Router (output) | `medialive.OutputGroupConfiguration.mediaConnectRouter()` | `mediaconnect.RouterInputConfiguration.mediaLiveChannel()` | `@aws-cdk/aws-mediaconnect-alpha` |
| MediaConnect Router (input) | `medialive.InputConfiguration.mediaConnectRouter()` | `mediaconnect.RouterOutputConfiguration.mediaLiveInput()` | `@aws-cdk/aws-mediaconnect-alpha` |

### AWS Elemental MediaPackage V2

Use `mediaPackageV2()` and pass a single `channel` — MediaLive maps each pipeline to a MediaPackage ingest endpoint automatically (one for `SINGLE_PIPELINE`, both for `STANDARD`). Each output contains a single encode (one track per output).

In-band captions (burn-in, embedded) ride alongside a video encode via the `captions` prop:

```go
var mpChannel IChannel
var hdVideo EncodeConfiguration
var sdVideo EncodeConfiguration
var audio EncodeConfiguration
var burnIn EncodeConfiguration


medialive.OutputGroupConfiguration_MediaPackageV2(&MediaPackageV2OutputGroupProps{
	Name: jsii.String("emp"),
	Channel: mpChannel,
	Outputs: []MediaPackageV2OutputDefinition{
		&MediaPackageV2OutputDefinition{
			Encode: hdVideo,
			Captions: []EncodeConfiguration{
				burnIn,
			},
			OutputName: jsii.String("hd"),
		},
		&MediaPackageV2OutputDefinition{
			Encode: sdVideo,
			OutputName: jsii.String("sd"),
		},
		&MediaPackageV2OutputDefinition{
			Encode: audio,
			OutputName: jsii.String("audio"),
		},
	},
})
```

For per-pipeline control — for example pinning pipeline 0 to a specific endpoint, or delivering each pipeline to a different (cross-region) channel — use `mediaPackageV2PerPipeline()` with explicit destinations:

```go
var primary IChannel
var hdVideo EncodeConfiguration


medialive.OutputGroupConfiguration_MediaPackageV2PerPipeline(&MediaPackageV2PerPipelineOutputGroupProps{
	Name: jsii.String("emp"),
	Destinations: []MediaPackageV2Destination{
		medialive.MediaPackageV2Destination_Channel(primary, medialive.MediaPackageV2EndpointId_ENDPOINT_2()),
		medialive.MediaPackageV2Destination_*Channel(primary, medialive.MediaPackageV2EndpointId_ENDPOINT_1()),
	},
	Outputs: []MediaPackageV2OutputDefinition{
		&MediaPackageV2OutputDefinition{
			Encode: hdVideo,
			OutputName: jsii.String("hd"),
		},
	},
})
```

### HLS

Use `OutputDestination.url()` for HTTP origins or `OutputDestination.toBucket()` for S3:

```go
var bucket IBucket
var video EncodeConfiguration
var audio EncodeConfiguration


// HLS to S3
medialive.OutputGroupConfiguration_Hls(&HlsOutputGroupProps{
	Name: jsii.String("hls_s3"),
	Destinations: []OutputDestination{
		medialive.OutputDestination_ToBucket(bucket, jsii.String("live/stream")),
	},
	Outputs: []HlsOutputDefinition{
		&HlsOutputDefinition{
			Encodes: []EncodeConfiguration{
				video,
				audio,
			},
			OutputName: jsii.String("hls_out"),
		},
	},
})

// HLS to an HTTPS CDN origin.
medialive.OutputGroupConfiguration_Hls(&HlsOutputGroupProps{
	Name: jsii.String("hls-http"),
	Destinations: []OutputDestination{
		medialive.OutputDestination_Url(jsii.String("https://203.0.113.10/ingest/stream")),
	},
	HlsCdnSettings: medialive.HlsCdnSettings_BasicPut(),
	Outputs: []HlsOutputDefinition{
		&HlsOutputDefinition{
			Encodes: []EncodeConfiguration{
				video,
				audio,
			},
			OutputName: jsii.String("hls_out"),
		},
	},
})
```

### Archive

Archive outputs write long-form recordings to S3:

```go
var bucket IBucket
var video EncodeConfiguration
var audio EncodeConfiguration


medialive.OutputGroupConfiguration_Archive(&ArchiveOutputGroupProps{
	Name: jsii.String("archive"),
	Destinations: []S3OutputDestination{
		medialive.S3OutputDestination_ToBucket(bucket, jsii.String("archive/recording")),
	},
	RolloverInterval: awscdk.Duration_Seconds(jsii.Number(600)),
	Outputs: []ArchiveOutputDefinition{
		&ArchiveOutputDefinition{
			Encodes: []EncodeConfiguration{
				video,
				audio,
			},
			OutputName: jsii.String("archive_out"),
		},
	},
})
```

### RTMP

RTMP outputs support H.264 + AAC only. Each output takes one destination per channel pipeline (the console's "Destination A" / "Destination B") via `RtmpDestination.url()` — one for `SINGLE_PIPELINE`, two for `STANDARD`:

```go
var video EncodeConfiguration
var audio EncodeConfiguration


medialive.OutputGroupConfiguration_Rtmp(&RtmpOutputGroupProps{
	Name: jsii.String("social"),
	Outputs: []RtmpOutputDefinition{
		&RtmpOutputDefinition{
			Encodes: []EncodeConfiguration{
				video,
				audio,
			},
			OutputName: jsii.String("live"),
			Destinations: []RtmpDestination{
				medialive.RtmpDestination_Url(jsii.String("rtmp://rtmp.example.com/live"), jsii.String("your-stream-key")),
			},
		},
	},
})
```

### SRT

SRT outputs use `SrtDestination.caller()` for caller mode or `SrtDestination.listener()` for listener mode. When you already have a full SRT URL rather than a separate host and port, use `SrtDestination.callerUrl()`. SRT output is always encrypted, so every destination takes an `encryptionPassphraseSecret` (a Secrets Manager secret). Each output takes one destination per channel pipeline ("Destination A"/"Destination B") — one for `SINGLE_PIPELINE`, two for `STANDARD`:

```go
var video EncodeConfiguration
var audio EncodeConfiguration
var passphrase ISecret


// SRT caller to a remote listener
medialive.OutputGroupConfiguration_Srt(&SrtOutputGroupProps{
	Name: jsii.String("srt_out"),
	Outputs: []SrtOutputDefinition{
		&SrtOutputDefinition{
			Encodes: []EncodeConfiguration{
				video,
				audio,
			},
			OutputName: jsii.String("srt_caller"),
			Destinations: []SrtDestination{
				medialive.SrtDestination_Caller(&SrtCallerDestinationProps{
					Address: jsii.String("203.0.113.20"),
					Port: jsii.Number(5000),
					EncryptionPassphraseSecret: passphrase,
				}),
			},
		},
	},
})

// SRT listener — MediaLive waits for the downstream system to connect
medialive.OutputGroupConfiguration_Srt(&SrtOutputGroupProps{
	Name: jsii.String("srt_listen"),
	Outputs: []SrtOutputDefinition{
		&SrtOutputDefinition{
			Encodes: []EncodeConfiguration{
				video,
				audio,
			},
			OutputName: jsii.String("srt_listener"),
			Destinations: []SrtDestination{
				medialive.SrtDestination_Listener(&SrtListenerDestinationProps{
					ListenerPort: jsii.Number(5000),
					EncryptionPassphraseSecret: passphrase,
				}),
			},
		},
	},
})
```

### AWS Elemental MediaConnect Router

`mediaConnectRouter()` delivers each channel pipeline to an AWS Elemental MediaConnect Router. Transit encryption defaults to AUTOMATIC; CDK derives one destination per pipeline from the channel class, so the common case needs no per-pipeline configuration. You must specify `availabilityZones` — exactly one for a `SINGLE_PIPELINE` channel, or two (one per pipeline) for `STANDARD`. The downstream wiring — which router input each pipeline feeds — is configured on the MediaConnect side, referencing this group's output by name and pipeline id.

```go
var video EncodeConfiguration
var audio EncodeConfiguration
var passphrase ISecret
var passphrase1 ISecret


// AUTOMATIC encryption on every pipeline (MPEG-TS container, like UDP)
medialive.OutputGroupConfiguration_MediaConnectRouter(&MediaConnectRouterOutputGroupProps{
	Name: jsii.String("router_out"),
	AvailabilityZones: []*string{
		jsii.String("us-east-1a"),
	},
	Outputs: []MediaConnectRouterOutputDefinition{
		&MediaConnectRouterOutputDefinition{
			Encodes: []EncodeConfiguration{
				video,
				audio,
			},
			OutputName: jsii.String("router_ts"),
		},
	},
})

// One shared Secrets Manager passphrase across all pipelines (SECRETS_MANAGER encryption)
medialive.OutputGroupConfiguration_MediaConnectRouter(&MediaConnectRouterOutputGroupProps{
	Name: jsii.String("router_out"),
	AvailabilityZones: []*string{
		jsii.String("us-east-1a"),
	},
	RouterSettings: medialive.MediaConnectRouterSettings_Shared(&MediaConnectRouterPipelineConfig{
		EncryptionSecret: passphrase,
	}),
	Outputs: []MediaConnectRouterOutputDefinition{
		&MediaConnectRouterOutputDefinition{
			Encodes: []EncodeConfiguration{
				video,
				audio,
			},
			OutputName: jsii.String("router_ts"),
		},
	},
})

// Distinct encryption per pipeline — an omitted pipeline stays AUTOMATIC (STANDARD channels)
medialive.OutputGroupConfiguration_MediaConnectRouter(&MediaConnectRouterOutputGroupProps{
	Name: jsii.String("router_out"),
	AvailabilityZones: []*string{
		jsii.String("us-east-1a"),
		jsii.String("us-east-1b"),
	},
	RouterSettings: medialive.MediaConnectRouterSettings_PerPipeline(&MediaConnectRouterPerPipelineSettings{
		Pipeline1: &MediaConnectRouterPipelineConfig{
			EncryptionSecret: passphrase1,
		},
	}),
	Outputs: []MediaConnectRouterOutputDefinition{
		&MediaConnectRouterOutputDefinition{
			Encodes: []EncodeConfiguration{
				video,
				audio,
			},
			OutputName: jsii.String("router_ts"),
		},
	},
})
```

When a passphrase secret is supplied, the channel's IAM role is automatically granted read access to it.

### UDP

UDP outputs deliver MPEG-TS over UDP or RTP. Use `UdpOutputDestination.udp()` for plain UDP or `.rtp()` for RTP (required if using FEC):

```go
var video EncodeConfiguration
var audio EncodeConfiguration


medialive.OutputGroupConfiguration_Udp(&UdpOutputGroupProps{
	Name: jsii.String("udp_out"),
	Destinations: []UdpOutputDestination{
		medialive.UdpOutputDestination_Udp(&TransportOutputDestinationProps{
			Address: jsii.String("203.0.113.5"),
			Port: jsii.Number(5000),
		}),
	},
	Outputs: []UdpOutputDefinition{
		&UdpOutputDefinition{
			Encodes: []EncodeConfiguration{
				video,
				audio,
			},
			OutputName: jsii.String("ts_out"),
		},
	},
})
```

### Frame Capture

Frame capture outputs write periodic JPEG snapshots to S3:

```go
var bucket IBucket
var video EncodeConfiguration


medialive.OutputGroupConfiguration_FrameCapture(&FrameCaptureOutputGroupProps{
	Name: jsii.String("thumbnails"),
	Destinations: []S3OutputDestination{
		medialive.S3OutputDestination_ToBucket(bucket, jsii.String("thumbnails/live")),
	},
	Outputs: []FrameCaptureOutputDefinition{
		&FrameCaptureOutputDefinition{
			Encodes: []EncodeConfiguration{
				video,
			},
			OutputName: jsii.String("thumb"),
		},
	},
})
```

### Microsoft Smooth Streaming

MS Smooth outputs push fragmented MP4 to an IIS Smooth Streaming endpoint:

```go
var video EncodeConfiguration
var audio EncodeConfiguration


medialive.OutputGroupConfiguration_MsSmooth(&MsSmoothOutputGroupProps{
	Name: jsii.String("smooth"),
	Destinations: []OutputDestination{
		medialive.OutputDestination_Url(jsii.String("https://smooth.example.com/live")),
	},
	Outputs: []MsSmoothOutputDefinition{
		&MsSmoothOutputDefinition{
			Encodes: []EncodeConfiguration{
				video,
				audio,
			},
			OutputName: jsii.String("smooth_out"),
		},
	},
})
```

### Per-output HLS settings

HLS outputs accept per-output `hlsSettings` via the `HlsSettings` factory — `standard()` for a video
rendition (with optional `M3u8Settings` for the transport stream), `audioOnly()` for an audio
rendition (with optional cover art as a [`FileLocation`](#file-locations)), `fmp4()`, or
`frameCapture()`.

```go
var bucket IBucket
var video EncodeConfiguration
var audio EncodeConfiguration


medialive.OutputGroupConfiguration_Hls(&HlsOutputGroupProps{
	Name: jsii.String("hls"),
	Destinations: []OutputDestination{
		medialive.OutputDestination_ToBucket(bucket, jsii.String("live/stream")),
	},
	Outputs: []HlsOutputDefinition{
		&HlsOutputDefinition{
			Encodes: []EncodeConfiguration{
				video,
			},
			OutputName: jsii.String("video"),
			HlsSettings: medialive.HlsSettings_Standard(&StandardHlsSettingsProps{
				M3u8Settings: medialive.M3u8Settings_Of(&M3u8SettingsProps{
					Scte35Behavior: medialive.M3u8Scte35Behavior_PASSTHROUGH(),
					ProgramNum: jsii.Number(1),
				}),
			}),
		},
		&HlsOutputDefinition{
			Encodes: []EncodeConfiguration{
				audio,
			},
			OutputName: jsii.String("audio"),
			HlsSettings: medialive.HlsSettings_AudioOnly(&AudioOnlyHlsSettingsProps{
				AudioGroupId: jsii.String("program"),
				AudioOnlyImage: medialive.FileLocation_FromBucket(bucket, jsii.String("art/cover.png")),
			}),
		},
	},
})
```

### Forward Error Correction (UDP)

UDP outputs accept optional `fec` settings (SMPTE 2022-1) — column-only or column-and-row FEC.
FEC requires an `rtp://` destination:

```go
var video EncodeConfiguration


medialive.OutputGroupConfiguration_Udp(&UdpOutputGroupProps{
	Name: jsii.String("udp"),
	Destinations: []UdpOutputDestination{
		medialive.UdpOutputDestination_Rtp(&TransportOutputDestinationProps{
			Address: jsii.String("203.0.113.5"),
			Port: jsii.Number(5000),
		}),
	},
	Outputs: []UdpOutputDefinition{
		&UdpOutputDefinition{
			Encodes: []EncodeConfiguration{
				video,
			},
			OutputName: jsii.String("ts"),
			Fec: &FecOutputSettings{
				Mode: medialive.FecMode_COLUMN_AND_ROW(),
				ColumnDepth: jsii.Number(10),
				RowLength: jsii.Number(10),
			},
		},
	},
})
```

### MPEG-TS Container Settings

The MPEG-TS output groups — `udp()`, `archive()`, `srt()`, and `mediaConnectRouter()` — accept optional per-output `m2tsSettings` via `M2tsSettings.of()`. Omit it to use MediaLive's service defaults. Bitrates use `Bitrate`, intervals use `Duration`, and closed-value fields use enums (e.g. `M2tsRateMode`, `M2tsScte35Control`); PID fields are strings that accept decimal, hexadecimal, ranges, or comma-separated lists.

```go
var video EncodeConfiguration
var audio EncodeConfiguration


medialive.OutputGroupConfiguration_Udp(&UdpOutputGroupProps{
	Name: jsii.String("udp_out"),
	Destinations: []UdpOutputDestination{
		medialive.UdpOutputDestination_Udp(&TransportOutputDestinationProps{
			Address: jsii.String("203.0.113.5"),
			Port: jsii.Number(5000),
		}),
	},
	Outputs: []UdpOutputDefinition{
		&UdpOutputDefinition{
			Encodes: []EncodeConfiguration{
				video,
				audio,
			},
			OutputName: jsii.String("ts"),
			M2tsSettings: medialive.M2tsSettings_Of(&M2tsSettingsProps{
				Bitrate: awscdk.Bitrate_Mbps(jsii.Number(8)),
				RateMode: medialive.M2tsRateMode_VBR(),
				ProgramNum: jsii.Number(1),
				PatInterval: awscdk.Duration_Millis(jsii.Number(100)),
				PmtInterval: awscdk.Duration_*Millis(jsii.Number(100)),
				Scte35Control: medialive.M2tsScte35Control_PASSTHROUGH(),
				DvbSdtSettings: &DvbSdtSettings{
					OutputSdt: medialive.DvbSdtOutputMode_SDT_MANUAL(),
					ServiceName: jsii.String("My Service"),
					RepInterval: awscdk.Duration_*Millis(jsii.Number(2000)),
				},
			}),
		},
	},
})
```

## Destinations

Each output group type uses a specific destination class. Destinations are created via static factory methods:

| Destination class | Factory methods | Used by |
|---|---|---|
| `OutputDestination` | `url()`, `toBucket()` | HLS, MS Smooth, CMAF Ingest |
| `S3OutputDestination` | `url()`, `toBucket()` | Archive, Frame Capture |
| `UdpOutputDestination` | `udp()`, `rtp()`, `url()` | UDP |
| `MediaPackageV2Destination` | `channel()` | MediaPackage V2 |
| `RtmpDestination` | `url()` | RTMP |
| `SrtDestination` | `caller()`, `callerUrl()`, `listener()` | SRT |

`OutputDestination.toBucket()` (and `S3OutputDestination.toBucket()`) build canonical `s3ssl://` URLs and automatically grant the channel's IAM role the required S3 permissions; `InputSource.fromBucket()` does the same for input reads. `MediaPackageV2Destination.channel()` automatically grants ingest permissions on the MediaPackage V2 channel.

The MediaConnect Router output group has no destination class — its delivery is configured on the MediaConnect side. Per-pipeline transit encryption is set via the group's `routerSettings` prop using `MediaConnectRouterSettings.shared()` / `.perPipeline()` (see [MediaConnect Router](#mediaconnect-router) above).

## Additional Destinations

MediaPackage V2 and CMAF Ingest output groups support `additionalDestinations` for cross-region delivery or backup packaging. These are separate from pipeline redundancy — they fan out the same content to extra endpoints.

The region for each destination is resolved automatically from the channel's stack. For cross-region imports, pass the region explicitly:

```go
var primaryChannel IChannel
var video EncodeConfiguration
var audio EncodeConfiguration


// Import a channel from another region — the region travels with the channel
backupChannel := mediapackagev2.Channel_FromChannelAttributes(this, jsii.String("BackupChannel"), &ChannelAttributes{
	ChannelName: jsii.String("backup-channel"),
	ChannelGroupName: jsii.String("backup-group"),
	Region: jsii.String("us-west-2"),
})

medialive.OutputGroupConfiguration_MediaPackageV2(&MediaPackageV2OutputGroupProps{
	Name: jsii.String("emp"),
	Channel: primaryChannel,
	AdditionalDestinations: []MediaPackageV2Destination{
		medialive.MediaPackageV2Destination_Channel(backupChannel, medialive.MediaPackageV2EndpointId_ENDPOINT_1()),
	},
	Outputs: []MediaPackageV2OutputDefinition{
		&MediaPackageV2OutputDefinition{
			Encode: video,
			OutputName: jsii.String("video"),
		},
		&MediaPackageV2OutputDefinition{
			Encode: audio,
			OutputName: jsii.String("audio"),
		},
	},
})
```

## Pipeline Redundancy

Channels default to `SINGLE_PIPELINE`. Set `channelClass: ChannelClass.STANDARD` for two-pipeline redundancy.

When using STANDARD:

* Each output group's `destinations` array must have two entries — `destinations[0]` maps to Pipeline 0, `destinations[1]` maps to Pipeline 1.
* For MediaPackage V2, use `ENDPOINT_1` for Pipeline 0 and `ENDPOINT_2` for Pipeline 1.
* `additionalDestinations` are separate from pipeline redundancy — they fan out to extra endpoints.

```go
var stack Stack
var input IInput
var bucket IBucket
var video EncodeConfiguration
var audio EncodeConfiguration


medialive.NewChannel(stack, jsii.String("StandardChannel"), &ChannelProps{
	ChannelClass: medialive.ChannelClass_STANDARD(),
	Inputs: []InputAttachment{
		&InputAttachment{
			Input: *Input,
		},
	},
	OutputGroups: []OutputGroupConfiguration{
		medialive.OutputGroupConfiguration_Hls(&HlsOutputGroupProps{
			Name: jsii.String("hls"),
			Destinations: []OutputDestination{
				medialive.OutputDestination_ToBucket(bucket, jsii.String("live/pipeline0")),
				medialive.OutputDestination_*ToBucket(bucket, jsii.String("live/pipeline1")),
			},
			Outputs: []HlsOutputDefinition{
				&HlsOutputDefinition{
					Encodes: []EncodeConfiguration{
						video,
						audio,
					},
					OutputName: jsii.String("hls_out"),
				},
			},
		}),
	},
})
```

## Input Attachment Settings

Each entry in `inputs` is an input attachment, which can carry per-input extraction and connection
settings beyond the input itself.

**Selectors** pick specific tracks out of the input. Use `AudioSelector` (`byLanguage()`, `byPid()`,
`byTrack()`, `hlsRendition()`, `default()`), `CaptionSelector` (`byLanguage()`, `embedded()`,
`ancillary()`, `dvbSub()`, `scte27()`, `teletext()`, `arib()`), and `videoSelector` (color space,
HDR10 metadata, and program/PID selection via `VideoSelection`). A caption encode then references a
caption selector by name.

```go
var stack Stack
var input IInput
var bucket IBucket
var video EncodeConfiguration


medialive.NewChannel(stack, jsii.String("Channel"), &ChannelProps{
	Inputs: []InputAttachment{
		&InputAttachment{
			Input: *Input,
			AudioSelectors: []AudioSelector{
				medialive.AudioSelector_ByLanguage(jsii.String("eng"), jsii.String("eng"), medialive.AudioLanguageSelectionPolicy_STRICT()),
			},
			CaptionSelectors: []CaptionSelector{
				medialive.CaptionSelector_Embedded(jsii.String("embedded")),
			},
			VideoSelector: &VideoSelectorSettings{
				ColorSpace: medialive.VideoColorSpace_HDR10(),
				ColorSpaceUsage: medialive.VideoColorSpaceUsage_FORCE(),
				SelectBy: medialive.VideoSelection_ByProgramId(jsii.Number(1)),
			},
		},
	},
	OutputGroups: []OutputGroupConfiguration{
		medialive.OutputGroupConfiguration_Hls(&HlsOutputGroupProps{
			Name: jsii.String("hls"),
			Destinations: []OutputDestination{
				medialive.OutputDestination_ToBucket(bucket, jsii.String("live/stream")),
			},
			Outputs: []HlsOutputDefinition{
				&HlsOutputDefinition{
					Encodes: []EncodeConfiguration{
						video,
					},
					OutputName: jsii.String("hls_out"),
				},
			},
		}),
	},
})
```

**Network input settings** apply to URL-pull and multicast inputs — HLS bandwidth/buffer/retry
behaviour, the SCTE-35 source (`HlsScte35Source.SEGMENTS` or `MANIFEST`), HTTPS server validation,
and a multicast source IP for source-specific multicast. `logicalInterfaceNames` maps the input to
network interfaces on MediaLive Anywhere nodes.

```go
var stack Stack
var input IInput
var bucket IBucket
var video EncodeConfiguration


medialive.NewChannel(stack, jsii.String("Channel"), &ChannelProps{
	Inputs: []InputAttachment{
		&InputAttachment{
			Input: *Input,
			NetworkInputSettings: &NetworkInputSettings{
				ServerValidation: medialive.ServerValidation_CHECK_CRYPTOGRAPHY_AND_VALIDATE_NAME(),
				HlsInputSettings: &HlsInputSettings{
					Bandwidth: awscdk.Bitrate_Mbps(jsii.Number(5)),
					Scte35Source: medialive.HlsScte35Source_MANIFEST(),
				},
			},
			LogicalInterfaceNames: []*string{
				jsii.String("eth0"),
				jsii.String("eth1"),
			},
		},
	},
	OutputGroups: []OutputGroupConfiguration{
		medialive.OutputGroupConfiguration_Hls(&HlsOutputGroupProps{
			Name: jsii.String("hls"),
			Destinations: []OutputDestination{
				medialive.OutputDestination_ToBucket(bucket, jsii.String("live/stream")),
			},
			Outputs: []HlsOutputDefinition{
				&HlsOutputDefinition{
					Encodes: []EncodeConfiguration{
						video,
					},
					OutputName: jsii.String("hls_out"),
				},
			},
		}),
	},
})
```

## Automatic Input Failover

Automatic input failover gives you input-*source* redundancy: attach a secondary input, and
MediaLive switches to it without restarting the channel when the active input meets a failover
condition. This is separate from the pipeline redundancy of `ChannelClass.STANDARD` (which
duplicates a single source across two pipelines).

Provide `automaticInputFailover` on the input attachment. If you don't specify conditions, a
single input-loss condition is used:

```go
var stack Stack
var primaryInput IInput
var secondaryInput IInput
var audioSelector AudioSelector
var video EncodeConfiguration
var audio EncodeConfiguration
var bucket IBucket


medialive.NewChannel(stack, jsii.String("Channel"), &ChannelProps{
	Inputs: []InputAttachment{
		&InputAttachment{
			Input: primaryInput,
			AutomaticInputFailover: &AutomaticInputFailover{
				SecondaryInput: *SecondaryInput,
				InputPreference: medialive.InputPreference_PRIMARY_INPUT_PREFERRED(),
				ErrorClearTime: awscdk.Duration_Seconds(jsii.Number(3)),
				FailoverConditions: []FailoverCondition{
					medialive.FailoverCondition_InputLoss(&InputLossFailoverProps{
						Threshold: awscdk.Duration_Millis(jsii.Number(1500)),
					}),
					medialive.FailoverCondition_AudioSilence(&AudioSilenceFailoverProps{
						AudioSelector: *AudioSelector,
						Threshold: awscdk.Duration_*Seconds(jsii.Number(2)),
					}),
					medialive.FailoverCondition_VideoBlack(&VideoBlackFailoverProps{
						BlackDetectThreshold: jsii.Number(0.1),
						Threshold: awscdk.Duration_*Seconds(jsii.Number(1)),
					}),
				},
			},
		},
		&InputAttachment{
			// The secondary input must also be attached to the channel as its own input.
			Input: secondaryInput,
		},
	},
	OutputGroups: []OutputGroupConfiguration{
		medialive.OutputGroupConfiguration_Hls(&HlsOutputGroupProps{
			Name: jsii.String("hls"),
			Destinations: []OutputDestination{
				medialive.OutputDestination_ToBucket(bucket, jsii.String("live/stream")),
			},
			Outputs: []HlsOutputDefinition{
				&HlsOutputDefinition{
					Encodes: []EncodeConfiguration{
						video,
						audio,
					},
					OutputName: jsii.String("hls_out"),
				},
			},
		}),
	},
})
```

The primary and secondary inputs must have the same input class. The channel's IAM role is
granted read access to the secondary input's sources automatically, just like the primary.

## Ad Avail Handling

MediaLive can blank content during ad avails, insert blackout slates, and signal SCTE-35 ad avails
to downstream systems. These are all channel-level props.

`availBlanking` replaces video/audio/captions with black (or an image) during an ad avail, and
`blackoutSlate` shows a slate when a SCTE-35 blackout is signalled. Both image fields take a
[`FileLocation`](#file-locations).

```go
var stack Stack
var input IInput
var bucket IBucket
var video EncodeConfiguration


medialive.NewChannel(stack, jsii.String("Channel"), &ChannelProps{
	Inputs: []InputAttachment{
		&InputAttachment{
			Input: *Input,
		},
	},
	AvailBlanking: &AvailBlanking{
		State: medialive.AvailBlankingState_ENABLED(),
		Image: medialive.FileLocation_FromBucket(bucket, jsii.String("slates/avail.png")),
	},
	BlackoutSlate: &BlackoutSlate{
		State: medialive.BlackoutSlateState_ENABLED(),
		Image: medialive.FileLocation_*FromBucket(bucket, jsii.String("slates/blackout.png")),
	},
	OutputGroups: []OutputGroupConfiguration{
		medialive.OutputGroupConfiguration_Hls(&HlsOutputGroupProps{
			Name: jsii.String("hls"),
			Destinations: []OutputDestination{
				medialive.OutputDestination_ToBucket(bucket, jsii.String("live/stream")),
			},
			Outputs: []HlsOutputDefinition{
				&HlsOutputDefinition{
					Encodes: []EncodeConfiguration{
						video,
					},
					OutputName: jsii.String("hls_out"),
				},
			},
		}),
	},
})
```

`availSettings` selects how SCTE-35 ad avails are handled — `AvailSettings.spliceInsert()`,
`AvailSettings.timeSignalApos()`, or `AvailSettings.esam()` for Event Signaling and Management
against an external POIS endpoint. `scte35SegmentationScope` controls which output groups receive
the segmentation cues. The ESAM POIS password is supplied as an SSM parameter, and the channel role
is granted read access to it automatically.

```go
import "github.com/aws/aws-cdk-go/awscdk"

var stack Stack
var input IInput
var bucket IBucket
var video EncodeConfiguration
var poisPassword StringParameter


medialive.NewChannel(stack, jsii.String("Channel"), &ChannelProps{
	Inputs: []InputAttachment{
		&InputAttachment{
			Input: *Input,
		},
	},
	AvailSettings: medialive.AvailSettings_Esam(&EsamSettings{
		Pois: &PoisEndpoint{
			Url: jsii.String("https://pois.example.com/esam"),
			Username: jsii.String("pois-user"),
			Password: poisPassword,
		},
		AcquisitionPointId: jsii.String("acquisition-point-1"),
		AdAvailOffset: awscdk.Duration_Millis(jsii.Number(200)),
	}),
	Scte35SegmentationScope: medialive.Scte35SegmentationScope_SCTE35_ENABLED_OUTPUT_GROUPS(),
	OutputGroups: []OutputGroupConfiguration{
		medialive.OutputGroupConfiguration_Hls(&HlsOutputGroupProps{
			Name: jsii.String("hls"),
			Destinations: []OutputDestination{
				medialive.OutputDestination_ToBucket(bucket, jsii.String("live/stream")),
			},
			Outputs: []HlsOutputDefinition{
				&HlsOutputDefinition{
					Encodes: []EncodeConfiguration{
						video,
					},
					OutputName: jsii.String("hls_out"),
				},
			},
		}),
	},
})
```

## Auto-Created Role and Grants

When no `role` is provided, the channel auto-creates an IAM role with the `medialive.amazonaws.com` service principal and grants it only the permissions your configuration actually needs. These automatic grants apply **only** to the channel-managed role; if you bring your own `role`, none are added.

**Channel role grants** — wired based on what you configure (channel-managed role only):

| Configuration | Grant | Scope |
|---|---|---|
| `OutputDestination.toBucket()` | S3 read/write | The destination bucket/prefix |
| `InputSource.fromBucket()` | S3 read | The input source bucket/prefix |
| `MediaPackageV2Destination.channel()` | `mediapackagev2:PutObject` | The MediaPackage V2 channel |
| `SrtDestination` with an encryption secret | Secrets Manager read | The secret |
| URL pull input with a password parameter | SSM parameter read | The parameter |
| Thumbnails (on by default) | `s3:PutObject` | `*` — uploads to an AWS service-owned bucket |
| Channel logging (`logLevel` set) | CloudWatch Logs write | The `ElementalMediaLive` log group in your account/region |
| VPC output (`vpc` set) | EC2 ENI create/delete + describe | Scoped to your subnets/SGs; `Describe*` requires `*` |

**Input role grants** — separate from the channel role, used at input create/delete time. Like the channel role, these are added only when the input auto-creates its role; pass a `role` to `mediaConnect()` or `cdi()` and no grants are added:

| Input type | Grant | Scope |
|---|---|---|
| `InputConfiguration.mediaConnect()` | `mediaconnect:ManagedDescribeFlow`, `ManagedAddOutput`, `ManagedRemoveOutput` | `*` — service rejects flow-scoped grants |
| `InputConfiguration.cdi()` | EC2 ENI create/delete + describe | Scoped to your subnets/SGs; `Describe*` requires `*` |

Both channel and input auto-created roles include confused-deputy prevention (`aws:SourceAccount` + `aws:SourceArn` conditions). For the full list of trusted-entity requirements, see [the documentation](https://docs.aws.amazon.com/medialive/latest/ug/trusted-entity-requirements.html).

The auto-created role is available on `channel.role` if you need to add further permissions.

### Bringing your pre-defined role

When you pass a `role`, the channel makes **no** automatic grants — you will need to add the permissions that role needs. That covers both the principal policy and any referenced resource policies: S3 output destinations and input sources, Secrets Manager and SSM reads, MediaPackage V2 ingest, CloudWatch Logs, and VPC output ENI management. See the [trusted-entity requirements](https://docs.aws.amazon.com/medialive/latest/ug/trusted-entity-requirements.html), or pass the account's `MediaLiveAccessRole` — an IAM role that MediaLive can assume.

## CloudWatch Metrics

Channels expose CloudWatch metric helpers in the `AWS/MediaLive` namespace, dimensioned by `ChannelId` and `Pipeline`. Use the named helpers below for the most common metrics, or `metric(metricName, pipeline)` to access any metric documented by the [MediaLive metrics reference](https://docs.aws.amazon.com/medialive/latest/ug/monitoring-eml-metrics.html).

MediaLive publishes metrics per pipeline. Every helper takes a `Pipeline` argument so you make an explicit decision about which pipeline you're monitoring. `STANDARD` channels run two redundant pipelines (`PIPELINE_0`, `PIPELINE_1`) — alarm on both to cover the full channel. `SINGLE_PIPELINE` channels only publish on `PIPELINE_0`; passing `PIPELINE_1` throws at synth time.

```go
var channel Channel
var stack Stack


channel.metricDroppedFrames(medialive.Pipeline_PIPELINE_0()).CreateAlarm(stack, jsii.String("DroppedFrames"), &CreateAlarmOptions{
	Threshold: jsii.Number(1),
	EvaluationPeriods: jsii.Number(2),
})

channel.metricSvqTime(medialive.Pipeline_PIPELINE_0()).CreateAlarm(stack, jsii.String("SvqTime"), &CreateAlarmOptions{
	Threshold: jsii.Number(0),
	EvaluationPeriods: jsii.Number(1),
})

// Custom metric by name with sum statistic
channel.metric(jsii.String("Output4xxErrors"), medialive.Pipeline_PIPELINE_0(), &MetricOptions{
	Statistic: jsii.String("sum"),
})
```

For STANDARD channels, alarm on both pipelines:

```go
var standardChannel Channel
var stack Stack


standardChannel.metricDroppedFrames(medialive.Pipeline_PIPELINE_0()).CreateAlarm(stack, jsii.String("Drops0"), &CreateAlarmOptions{
	Threshold: jsii.Number(1),
	EvaluationPeriods: jsii.Number(2),
})
standardChannel.metricDroppedFrames(medialive.Pipeline_PIPELINE_1()).CreateAlarm(stack, jsii.String("Drops1"), &CreateAlarmOptions{
	Threshold: jsii.Number(1),
	EvaluationPeriods: jsii.Number(2),
})
```

### Channel metrics

| Helper | Metric name | Default statistic | Notes |
|---|---|---|---|
| `metricActiveAlerts(pipeline)` | `ActiveAlerts` | Max | Total active alerts on the channel |
| `metricNetworkIn(pipeline)` | `NetworkIn` | Avg | Inbound traffic in Mbps |
| `metricNetworkOut(pipeline)` | `NetworkOut` | Avg | Outbound traffic in Mbps |
| `metricInputVideoFrameRate(pipeline)` | `InputVideoFrameRate` | Max | Source video frame rate |
| `metricFillMsec(pipeline)` | `FillMsec` | Max | Time filled with fill frames — non-zero indicates input loss |
| `metricInputLossSeconds(pipeline)` | `InputLossSeconds` | Sum | Seconds without packets (RTP / MediaConnect inputs) |
| `metricDroppedFrames(pipeline)` | `DroppedFrames` | Sum | Frames dropped because the encoder fell behind |
| `metricSvqTime(pipeline)` | `SvqTime` | Max | Percent of time MediaLive reduced quality to keep up |
| `metric(name, pipeline, props?)` | (custom) | (caller-provided) | Build any metric in `AWS/MediaLive` |

The defaults match the AWS-recommended statistic for each metric. Pass `props` to override statistic, period, dimensions, or any other `MetricOptions` field.

## MediaLive Anywhere

MediaLive Anywhere lets you run MediaLive channels on your own on-premises hardware.

Certain input types are only available with Anywhere channels (channels configured with `anywhereSettings`):
SDI, SMPTE 2110 Receiver Group, and Multicast. Attempting to use these input types on a cloud channel will throw a validation error at synth time.

### Network

A network defines IP address pools and routes for Anywhere resources:

```go
var stack Stack

network := medialive.NewNetwork(stack, jsii.String("Network"), &NetworkProps{
	NetworkName: jsii.String("on-prem-network"),
	IpPools: []*string{
		jsii.String("10.0.0.0/24"),
	},
	Routes: []NetworkRoute{
		&NetworkRoute{
			Cidr: jsii.String("0.0.0.0/0"),
			Gateway: jsii.String("10.0.0.1"),
		},
	},
})
```

### Cluster

A cluster represents a group of on-premises hardware nodes:

```go
var stack Stack
var instanceRole IRole


cluster := medialive.NewCluster(stack, jsii.String("Cluster"), &ClusterProps{
	ClusterName: jsii.String("on-prem-cluster"),
	ClusterType: medialive.ClusterType_ON_PREMISES(),
	InstanceRole: InstanceRole,
})
```

### Channel Placement Group

A channel placement group assigns channels to specific nodes within a cluster. Associate it with a channel via `anywhereSettings`:

```go
var stack Stack
var cluster ICluster
var input IInput
var video EncodeConfiguration
var bucket IBucket


cpg := medialive.NewChannelPlacementGroup(stack, jsii.String("CPG"), &ChannelPlacementGroupProps{
	ChannelPlacementGroupName: jsii.String("my-cpg"),
	Cluster: Cluster,
})

medialive.NewChannel(stack, jsii.String("AnywhereChannel"), &ChannelProps{
	Inputs: []InputAttachment{
		&InputAttachment{
			Input: *Input,
		},
	},
	AnywhereSettings: &AnywhereSettings{
		Cluster: *Cluster,
		ChannelPlacementGroup: cpg,
	},
	OutputGroups: []OutputGroupConfiguration{
		medialive.OutputGroupConfiguration_Hls(&HlsOutputGroupProps{
			Name: jsii.String("hls"),
			Destinations: []OutputDestination{
				medialive.OutputDestination_ToBucket(bucket, jsii.String("live/stream")),
			},
			Outputs: []HlsOutputDefinition{
				&HlsOutputDefinition{
					Encodes: []EncodeConfiguration{
						video,
					},
					OutputName: jsii.String("hls_out"),
				},
			},
		}),
	},
})
```

### SDI Source

An SDI source represents a physical SDI input on Anywhere hardware:

```go
var stack Stack

sdi := medialive.NewSdiSource(stack, jsii.String("Sdi"), &SdiSourceProps{
	SdiSourceName: jsii.String("camera-1"),
	Type: medialive.SdiType_SINGLE(),
})
```

### On-premises input networking

For inputs that live in an on-premises network, set `inputNetworkLocation` to
`InputNetworkLocation.ON_PREMISES`. On-premises inputs do not use input security groups. Push
inputs (RTMP/RTP/UDP) can pin their destination to a `Network`, declare the `networkRoutes` to
reach it on the local network, and request a `staticIpAddress`:

```go
var stack Stack


network := medialive.NewNetwork(stack, jsii.String("Network"), &NetworkProps{
	NetworkName: jsii.String("on-prem-network"),
	IpPools: []*string{
		jsii.String("192.168.1.0/24"),
	},
})

medialive.NewInput(stack, jsii.String("OnPremInput"), &InputProps{
	InputName: jsii.String("on-prem-rtp"),
	InputNetworkLocation: medialive.InputNetworkLocation_ON_PREMISES(),
	Input: medialive.InputConfiguration_RtpPush(&PushInputProps{
		Destinations: []PushInputDestination{
			&PushInputDestination{
				Network: *Network,
				NetworkRoutes: []NetworkRoute{
					&NetworkRoute{
						Cidr: jsii.String("10.0.0.0/24"),
						Gateway: jsii.String("10.0.0.1"),
					},
				},
				StaticIpAddress: jsii.String("192.168.1.50"),
			},
		},
	}),
})
```

SRT listener inputs accept a `streamId` that the upstream system uses when connecting:

```go
var stack Stack
var sg IInputSecurityGroup


medialive.NewInput(stack, jsii.String("SrtListener"), &InputProps{
	InputName: jsii.String("srt-listener"),
	Input: medialive.InputConfiguration_SrtListener(&SrtListenerInputProps{
		InputSecurityGroups: []IInputSecurityGroupRef{
			sg,
		},
		StreamId: jsii.String("my-stream-id"),
	}),
})
```
