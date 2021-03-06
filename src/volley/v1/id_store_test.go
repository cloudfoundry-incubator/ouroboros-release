package v1_test

import (
	"volley/v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("AppIDStore", func() {
	It("returns an app ID weighted on the number of times it has been added", func() {
		store := v1.NewIDStore(3)
		store.Add("some-id")
		store.Add("some-id")
		store.Add("some-more-id")

		firstCount, secondCount := 0, 0
		for i := 0; i < 1000; i++ {
			switch store.Get() {
			case "some-id":
				firstCount++
			case "some-more-id":
				secondCount++
			}
		}
		Expect(firstCount).Should(BeNumerically("~", 666, 50))
		Expect(secondCount).Should(BeNumerically("~", 333, 50))
	})

	It("blocks until it is full", func() {
		store := v1.NewIDStore(3)
		store.Add("some-id")
		store.Add("some-id")

		done := make(chan struct{})

		go func() {
			store.Get()
			close(done)
		}()

		Consistently(done).ShouldNot(BeClosed())

		store.Add("some-more-id")
		Eventually(done).Should(BeClosed())
	})

	It("can get multiple unique values", func() {
		store := v1.NewIDStore(3)
		store.Add("some-id-1")
		store.Add("some-id-2")
		store.Add("some-id-3")

		Expect(store.GetN(3)).To(ConsistOf(
			"some-id-1",
			"some-id-2",
			"some-id-3",
		))
	})

	It("does not return empty keys requesting more then are available", func() {
		store := v1.NewIDStore(2)
		store.Add("some-id-1")

		Expect(store.GetN(3)).To(ConsistOf(
			"some-id-1",
		))
	})

	It("returns an empty list with an empty store", func() {
		store := v1.NewIDStore(2)
		Expect(store.GetN(3)).To(BeEmpty())
	})

	It("handles more adds then capacity allows", func() {
		store := v1.NewIDStore(2)
		store.Add("some-id-1")
		store.Add("some-id-2")
		store.Add("some-id-3")

		Expect(store.GetN(3)).To(ConsistOf(
			"some-id-3",
			"some-id-2",
		))
	})
})
