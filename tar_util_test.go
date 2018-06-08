package go_common

import (

	"testing"
)


func TestTar(t *testing.T) {

	t.Run("tar test", func(t *testing.T) {
		Tar("/Users/chaojie.xiao/shaw/code/yintech/nodejs/lottery-core","./a.tar")
	})
}