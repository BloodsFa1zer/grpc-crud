package handler

import (
	"app4/database"
	pb "app4/proto"
	"app4/service"
	"context"
	"google.golang.org/grpc"
)

type ServerService struct {
	UserService service.ServiceInterface
	GrpcServer  *grpc.Server
	pb.UnimplementedUserServiceServer
}

func NewServerService(grpcServer *grpc.Server, UserService service.ServiceInterface) {
	userGrpc := &ServerService{UserService: UserService}
	pb.RegisterUserServiceServer(grpcServer, userGrpc)
}

func (ss *ServerService) Create(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserProfileResponse, error) {
	data := ss.transformUserRPC(req)
	userID, err := ss.UserService.Create(*data)
	if err != nil {
		return nil, err
	}
	// no need to check error as transformUserModel() used only with previously created users
	user, _ := ss.UserService.Get(userID)
	return ss.transformUserModel(*user), nil
}

func (ss *ServerService) Get(ctx context.Context, req *pb.SingleUserRequest) (*pb.UserProfileResponse, error) {
	ID := req.GetID()

	user, err := ss.UserService.Get(ID)
	if err != nil {
		return nil, err
	}

	return ss.transformUserModel(*user), nil
}

func (ss *ServerService) transformUserRPC(req *pb.CreateUserRequest) *database.User {
	return &database.User{
		ID:        req.GetID(),
		Nickname:  req.GetNickname(),
		FirstName: req.GetFirstName(),
		LastName:  req.GetLastName(),
		Password:  req.GetPassword(),
		Role:      req.GetRole(),
		CreatedAt: req.GetCreatedAt(),
		UpdatedAt: req.GetUpdatedAt(),
		DeletedAt: req.GetDeletedAt(),
	}
}

func (ss *ServerService) transformUserModel(user database.User) *pb.UserProfileResponse {
	return &pb.UserProfileResponse{
		ID:        user.ID,
		Nickname:  user.Nickname,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Password:  user.Password,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}
}
