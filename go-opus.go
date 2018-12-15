package opus

/*
#cgo LDFLAGS: -lopus
#include <stdlib.h>
#include <opus/opus.h>
*/
import "C"

type OpusApplication int
type OpusError int

const (
	// Opus Application Types
	AppVoIP               = OpusApplication(C.OPUS_APPLICATION_VOIP)
	AppAudio              = OpusApplication(C.OPUS_APPLICATION_AUDIO)
	AppRestrictedLowDelay = OpusApplication(C.OPUS_APPLICATION_RESTRICTED_LOWDELAY)

	// Opus Error Types
	Ok            = OpusError(C.OPUS_OK)
	BadArg        = OpusError(C.OPUS_BAD_ARG)
	BufferToSmall = OpusError(C.OPUS_BUFFER_TOO_SMALL)
	InternalError = OpusError(C.OPUS_INTERNAL_ERROR)
	InvalidPacket = OpusError(C.OPUS_INVALID_PACKET)
	Unimplemented = OpusError(C.OPUS_UNIMPLEMENTED)
	InvalidState  = OpusError(C.OPUS_INVALID_STATE)
	AllocFail     = OpusError(C.OPUS_ALLOC_FAIL)
)

func CreateEncoder() OpusError {
	var err C.int
	enc := C.opus_encoder_create(16000, 2, C.int(AppAudio), &err)
	defer C.opus_encoder_destroy(enc)
	return OpusError(err)
}

func CreateDecoder() OpusError {
	var err C.int
	dec := C.opus_decoder_create(48000, 2, &err)
	defer C.opus_decoder_destroy(dec)
	return OpusError(err)
}
