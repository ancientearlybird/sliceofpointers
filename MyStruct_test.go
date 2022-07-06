package sliceofpointers_test

import (
	"fmt"
	"os"
	"strconv"
	"time"

	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega/gmeasure"
	s "gitlab.com/h19900401/playground/sliceofpointers"
)

var _ = Describe("MyStruct", func() {
	lenOfSlice := 20
	numSamples := 10000

	When("used as pointer", func() {
		It("on creation", func() {
			experiment := gmeasure.NewExperiment("slice of pointers")
			AddReportEntry(experiment.Name, experiment)

			experiment.Sample(func(_ int) {
				experiment.MeasureDuration("create slice of pointers", func() {
					slice := make([]*s.MyStruct, 0, lenOfSlice)
					for j := 0; j < lenOfSlice; j++ {
						slice = append(slice, &s.MyStruct{A: j, B: j + 1})
					}

					_ = slice
				}, gmeasure.Precision(time.Nanosecond))
			}, gmeasure.SamplingConfig{N: numSamples, Duration: time.Minute})

			m1 := experiment.Get("create slice of pointers")

			err := writeCSV(convertToNanoseconds(m1.Durations), "sliceOfPointers.csv")
			_ = err
		})
	})

	When("used as value", func() {
		It("on creation", func() {
			experiment := gmeasure.NewExperiment("slice of values")
			AddReportEntry(experiment.Name, experiment)

			experiment.Sample(func(_ int) {
				experiment.MeasureDuration("create slice of values", func() {
					slice := make([]s.MyStruct, 0, lenOfSlice)
					for j := 0; j < lenOfSlice; j++ {
						slice = append(slice, s.MyStruct{A: j, B: j + 1})
					}

					_ = slice
				}, gmeasure.Precision(time.Nanosecond))
			}, gmeasure.SamplingConfig{N: numSamples, Duration: time.Minute})

			m1 := experiment.Get("create slice of values")

			err := writeCSV(convertToNanoseconds(m1.Durations), "sliceOfValues.csv")
			_ = err
		})
	})
})

func convertToNanoseconds(in []time.Duration) []int64 {
	out := make([]int64, 0, len(in))

	for _, v := range in {
		out = append(out, v.Nanoseconds())
	}

	return out
}

func writeCSV(in []int64, filename string) error {
	_ = os.Remove(filename)

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("cannot create file: %w", err)
	}

	defer file.Close()

	for idx := range in {
		_, err := file.WriteString(strconv.FormatInt(in[idx], 10) + "\n")
		if err != nil {
			return fmt.Errorf("cannot write number: %w", err)
		}
	}

	return nil
}
