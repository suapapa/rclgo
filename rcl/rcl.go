package rcl

// #cgo CFLAGS: -I/opt/ros/foxy/include
// #cgo LDFLAGS: -L/opt/ros/foxy/lib -lrcl -lrcutils
// #include <rcl/rcl.h>
// typedef rcl_context_t* rcl_ContextP;
import "C"
import (
	"rclgo/types"
	"unsafe"
)

type Context struct {
	RCLContext *C.rcl_context_t
}

func GetZeroInitializatiedContext() Context {
	ctxP := (C.rcl_ContextP)(C.malloc(C.sizeof_rcl_context_t))
	cContext := C.rcl_get_zero_initialized_context()
	*ctxP = cContext
	return Context{ctxP}
}

func (ctx *Context) Shutdown() {
	C.free(unsafe.Pointer(ctx.RCLContext))
}

//Init represents the global initialization of rcl. This must be called before using any rcl functions.
func Init(rclCtx Context) types.RCLRetT {
	// argv := make([]*C.char, 1)
	// argv[0] = C.CString("")
	// defer C.free(unsafe.Pointer(argv[0]))

	opts := C.rcl_get_zero_initialized_init_options()
	ret := C.rcl_init_options_init(&opts, C.rcl_get_default_allocator())
	if ret != types.RCL_RET_OK {
		return types.RCLRetT(ret)
	}

	// return types.RCLRetT(C.rcl_init(0, &argv[0], &opts, rclCtx.RCLContext))
	return types.RCLRetT(C.rcl_init(0, (**C.char)(C.NULL), &opts, rclCtx.RCLContext))
}

//Shutdown represents Signal global shutdown of rcl.
func Shutdown(rclCtx Context) {
	C.rcl_shutdown(rclCtx.RCLContext)
	rclCtx.Shutdown()
}
