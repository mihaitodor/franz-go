module aws_msk_iam

go 1.16

require (
	github.com/aws/aws-sdk-go v1.40.42
	github.com/twmb/franz-go v0.8.3
	github.com/twmb/franz-go/pkg/kmsg v0.0.0-20220310045829-806cf534caf0
)

replace github.com/twmb/franz-go => ../../..
