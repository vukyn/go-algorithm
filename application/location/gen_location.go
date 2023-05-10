package location

import (
	"fmt"
	"go-algorithms/application/utils"
	"log"
	"math/rand"
	"strings"
	"time"
)

func GenLocation() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	locations := make([]string, 0)
	for i := 0; i < 10; i++ {
		FLOOR := fmt.Sprintf("F%d", r1.Intn(9))
		AREA := fmt.Sprintf("A%d", r1.Intn(9))
		RACK := fmt.Sprintf("0%d", r1.Intn(9))
		ASILE := fmt.Sprintf("0%d", r1.Intn(9))
		SHELF := fmt.Sprintf("0%d", r1.Intn(9))
		BIN := fmt.Sprintf("0%d", r1.Intn(9))

		locations = append(locations, fmt.Sprintf("%s-%s-%s-%s-%s-%s", FLOOR, AREA, RACK, ASILE, SHELF, BIN))
	}

	if err := utils.WriteFile(strings.Join(locations, "\n"), "assets/location.txt"); err != nil {
		log.Fatal(err)
	}
}
