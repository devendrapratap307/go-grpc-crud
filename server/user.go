package server

import (
	"context"
	"grpc-crud/models"
	pb "grpc-crud/pb"

	"gorm.io/gorm"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	DB *gorm.DB
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	user := models.User{Name: req.Name, Email: req.Email}
	s.DB.Create(&user)
	return &pb.User{Id: user.ID, Name: user.Name, Email: user.Email}, nil
}

func (s *UserService) GetUser(ctx context.Context, req *pb.UserId) (*pb.User, error) {
	var user models.User
	if err := s.DB.First(&user, req.Id).Error; err != nil {
		return nil, err
	}
	return &pb.User{Id: user.ID, Name: user.Name, Email: user.Email}, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	user := models.User{ID: req.Id}
	if err := s.DB.First(&user).Error; err != nil {
		return nil, err
	}
	user.Name = req.Name
	user.Email = req.Email
	s.DB.Save(&user)
	return req, nil
}

func (s *UserService) DeleteUser(ctx context.Context, req *pb.UserId) (*pb.Empty, error) {
	s.DB.Delete(&models.User{}, req.Id)
	return &pb.Empty{}, nil
}

func (s *UserService) ListUsers(req *pb.Empty, stream pb.UserService_ListUsersServer) error {
	var users []models.User
	s.DB.Find(&users)
	for _, u := range users {
		stream.Send(&pb.User{Id: u.ID, Name: u.Name, Email: u.Email})
	}
	return nil
}
