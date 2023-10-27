package handler

import (
	"app4/database"
	pb "app4/proto"
	"app4/service"
	"context"
	validator "github.com/mwitkow/go-proto-validators"
	"google.golang.org/grpc"
)

type ServerService struct {
	UserService service.UserServiceInterface
	GrpcServer  *grpc.Server
	pb.UnimplementedUserServiceServer
}

func NewServerService(UserService service.UserServiceInterface) *ServerService {
	userGrpc := &ServerService{UserService: UserService}
	return userGrpc
}

func (ss *ServerService) Create(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserProfileResponse, error) {

	validate := validator.Validator(req)
	err := validate.Validate()
	if err != nil {
		return nil, err
	}

	data := ss.transformUserRPCCreateUser(req)
	if err != nil {
		return nil, err
	}

	userID, err := ss.UserService.Create(*data)
	if err != nil {
		return nil, err
	}

	// no need to check error as transformUserModel() used only with previously created users
	user, _ := ss.UserService.Get(userID)
	return ss.transformUserModel(*user), nil
}

func (ss *ServerService) Read(ctx context.Context, req *pb.SingleUserRequest) (*pb.UserProfileResponse, error) {

	validate := validator.Validator(req)
	err := validate.Validate()
	if err != nil {
		return nil, err
	}

	ID := req.GetID()

	user, err := ss.UserService.Get(ID)
	if err != nil {
		return nil, err
	}

	return ss.transformUserModel(*user), nil
}

func (ss *ServerService) Update(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UserProfileResponse, error) {
	validate := validator.Validator(req)
	err := validate.Validate()
	if err != nil {
		return nil, err
	}

	data := ss.transformUserRPCUpdateUser(req)
	if err != nil {
		return nil, err
	}

	updatedID, err := ss.UserService.Update(data.ID, *data)
	if err != nil {
		return nil, err
	}

	user, _ := ss.UserService.Get(updatedID)

	return ss.transformUserModel(*user), nil
}

func (ss *ServerService) Delete(ctx context.Context, req *pb.SingleUserRequest) (*pb.SuccessResponse, error) {
	validate := validator.Validator(req)
	err := validate.Validate()
	if err != nil {
		return nil, err
	}

	ID := req.GetID()

	err = ss.UserService.Delete(ID)
	if err != nil {
		return nil, err
	}

	return &pb.SuccessResponse{Response: "User Successfully deleted"}, nil
}

func (ss *ServerService) transformUserRPCCreateUser(req *pb.CreateUserRequest) *database.User {

	return &database.User{
		Nickname:  req.GetNickname(),
		FirstName: req.GetFirstName(),
		LastName:  req.GetLastName(),
		Password:  req.GetPassword(),
	}
}

func (ss *ServerService) transformUserRPCUpdateUser(req *pb.UpdateUserRequest) *database.User {

	return &database.User{
		ID:        req.GetID(),
		Nickname:  req.GetNickname(),
		FirstName: req.GetFirstName(),
		LastName:  req.GetLastName(),
		Password:  req.GetPassword(),
	}
}

func (ss *ServerService) transformUserModel(user database.User) *pb.UserProfileResponse {
	return &pb.UserProfileResponse{
		ID:        user.ID,
		Nickname:  user.Nickname,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}
}
