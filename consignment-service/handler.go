package main

import (
	"context"
	"log"

	pb "github.com/vfuntikov/shippy/consignment-service/proto/consignment"
	vesselProto "github.com/vfuntikov/shippy/vessel-service/proto/vessel"
	micro "go-micro.dev/v4"
)

type Repository interface {
	Create(context.Context, *pb.Consignment) error
	GetAll(context.Context) ([]*pb.Consignment, error)
}

// Service should implement all of the methods to satisfy the service
// we defined in our protobuf definition. You can check the interface
// in the generated code itself for the exact method signatures etc
// to give you a better idea.
type Service struct {
	repo         Repository
	vesselClient vesselProto.VesselService
	srv          micro.Service
}

func NewService(repo Repository) *Service {
	// Create a new service. Optionally include some options here.
	srv := micro.NewService(
		// This name must match the package name given in your protobuf definition
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)

	// Init will parse the command line flags.
	srv.Init()

	vesselClient := vesselProto.NewVesselService("go.micro.srv.vessel", srv.Client())
	return &Service{
		repo:         repo,
		vesselClient: vesselClient,
		srv:          srv,
	}
}

// CreateConsignment - we created just one method on our service,
// which is a create method, which takes a context and a request as an
// argument, these are handled by the gRPC server.
func (s *Service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	// Here we call a client instance of our vessel service with our consignment weight,
	// and the amount of containers as the capacity value
	vesselResponse, err := s.vesselClient.FindAvailable(ctx, &vesselProto.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})
	log.Printf("Found vessel: %s \n", vesselResponse.Vessel.Name)
	if err != nil {
		log.Printf("could not find vessel: %v\n", err)
		return err
	}

	// We set the VesselId as the vessel we got back from our
	// vessel service
	req.VesselId = vesselResponse.Vessel.Id

	err = s.repo.Create(ctx, req)
	if err != nil {
		log.Printf("could not create consignment: %v\n", err)
		return err
	}
	// Return matching the `Response` message we created in our
	// protobuf definition.
	res.Created = true
	res.Consignment = req
	return nil
}

func (s *Service) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {

	consignments, err := s.repo.GetAll(ctx)
	if err != nil {
		log.Printf("could not get consignments: %v\n", err)
		return err
	}
	res.Consignments = consignments
	return nil
}

func (s *Service) Serve() error {
	// Register handler
	pb.RegisterShippingServiceHandler(s.srv.Server(), s)

	// log.Println("Running on port:", port)
	// Run the server
	if err := s.srv.Run(); err != nil {
		log.Printf("could not serve service: %v\n", err)
		return err
	}
	return nil
}
