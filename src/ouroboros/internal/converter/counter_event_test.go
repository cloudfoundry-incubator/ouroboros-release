package converter_test

import (
	v2 "loggregator/v2"
	"ouroboros/internal/converter"

	"github.com/cloudfoundry/sonde-go/events"
	"github.com/gogo/protobuf/proto"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
)

var _ = Describe("CounterEvent", func() {
	Context("given a v2 envelope", func() {
		It("converts to a v1 envelope", func() {
			envelope := &v2.Envelope{
				Message: &v2.Envelope_Counter{
					Counter: &v2.Counter{
						Name: "name",
						Value: &v2.Counter_Delta{
							Delta: 99,
						},
					},
				},
			}

			Expect(*converter.ToV1(envelope)).To(MatchFields(IgnoreExtras, Fields{
				"EventType": Equal(events.Envelope_CounterEvent.Enum()),
				"CounterEvent": Equal(&events.CounterEvent{
					Name:  proto.String("name"),
					Delta: proto.Uint64(99),
					Total: proto.Uint64(0),
				}),
			}))
		})
	})

	Context("given a v1 envelope", func() {
		It("converts to a v2 envelope", func() {
			v1Envelope := &events.Envelope{
				EventType: events.Envelope_CounterEvent.Enum(),
				CounterEvent: &events.CounterEvent{
					Name:  proto.String("name"),
					Delta: proto.Uint64(99),
				},
			}
			v2Envelope := &v2.Envelope{
				Message: &v2.Envelope_Counter{
					Counter: &v2.Counter{
						Name: "name",
						Value: &v2.Counter_Delta{
							Delta: 99,
						},
					},
				},
			}

			Expect(*converter.ToV2(v1Envelope)).To(MatchFields(IgnoreExtras, Fields{
				"Message": Equal(v2Envelope.Message),
			}))
		})
	})
})
