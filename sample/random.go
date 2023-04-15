package sample

import (
	"grpcProject/pb"
	"math/rand"
)

func randomBool() bool {
	return rand.Intn(2) == 0
}

func randInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func randomKeyboardLayout() pb.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return pb.Keyboard_QWERTY
	case 2:
		return pb.Keyboard_QWERTZ
	default:
		return pb.Keyboard_AZERTY
	}
}

func randomString(a ...string) string {
	n := len(a)
	return a[rand.Intn(n)]
}

func randFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)

}
