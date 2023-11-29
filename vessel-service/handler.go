package main

import (
	"context"
	"log"

	pb "github.com/vfuntikov/shippy/vessel-service/proto/vessel"
	micro "go-micro.dev/v4"
)

type Repository interface {
	// Create(context.Context, *pb.Consignment) error
	// GetAll(context.Context) ([]*pb.Consignment, error)
	FindAvailable(context.Context, *pb.Specification) (*pb.Vessel, error)
	Create(context.Context, *pb.Vessel) error
}

// Service should implement all of the methods to satisfy the service
// we defined in our protobuf definition. You can check the interface
// in the generated code itself for the exact method signatures etc
// to give you a better idea.
type Service struct {
	repo Repository
	srv  micro.Service
}

func NewService(repo Repository) *Service {
	// Create a new service. Optionally include some options here.
	srv := micro.NewService(
		// This name must match the package name given in your protobuf definition
		micro.Name("go.micro.srv.vessel"),
		micro.Version("latest"),
	)

	// Init will parse the command line flags.
	srv.Init()

	return &Service{
		repo: repo,
		srv:  srv,
	}
}

func (s *Service) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {
	// defer s.GetRepo().Close()
	// Find the next available vessel
	vessel, err := s.repo.FindAvailable(ctx, req)
	if err != nil {
		return err
	}

	// Set the vessel as part of the response message type
	res.Vessel = vessel
	return nil
}

func (s *Service) Create(ctx context.Context, req *pb.Vessel, res *pb.Response) error {
	// defer s.GetRepo().Close()
	if err := s.repo.Create(ctx, req); err != nil {
		return err
	}
	res.Vessel = req
	res.Created = true
	return nil
}

func (s *Service) Serve() error {
	// Register handler
	pb.RegisterVesselServiceHandler(s.srv.Server(), s)

	// log.Println("Running on port:", port)
	// Run the server
	if err := s.srv.Run(); err != nil {
		log.Printf("could not serve service: %v\n", err)
		return err
	}
	return nil
}
