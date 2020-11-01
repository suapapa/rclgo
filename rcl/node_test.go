package rcl

import (
	"fmt"
	"testing"
	"time"
)

func TestNodeCreation(t *testing.T) {

	// * TODO for this test: check the retValues from C functions

	// Initialization
	rclCtx := GetZeroInitializatiedContext()
	Init(rclCtx)
	myNode := GetZeroInitializedNode()
	myNodeOpts := GetNodeDefaultOptions()

	fmt.Printf("Creating the node! \n")
	NodeInit(&myNode, "fakeNameForNode", "", rclCtx, myNodeOpts)
	time.Sleep(10 * time.Second) // or runtime.Gosched() or similar per @misterbee
	NodeFini(myNode)
	Shutdown(rclCtx)
}
